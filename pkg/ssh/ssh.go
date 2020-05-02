/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// WARNING: DO NOT add new use-caes to this package as it is deprecated and slated for deletion.

package ssh

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	mathrand "math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"

	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/klog"
)

/*
 * By default, all the following metrics are defined as falling under
 * ALPHA stability level https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/20190404-kubernetes-control-plane-metrics-stability.md#stability-classes)
 *
 * Promoting the stability level of the metric is a responsibility of the component owner, since it
 * involves explicitly acknowledging support for the metric across multiple releases, in accordance with
 * the metric stability policy.
 */
var (
	tunnelOpenCounter = metrics.NewCounter(
		&metrics.CounterOpts{
			Name:           "ssh_tunnel_open_count",
			Help:           "Counter of ssh tunnel total open attempts",
			StabilityLevel: metrics.ALPHA,
		},
	)
	tunnelOpenFailCounter = metrics.NewCounter(
		&metrics.CounterOpts{
			Name:           "ssh_tunnel_open_fail_count",
			Help:           "Counter of ssh tunnel failed open attempts",
			StabilityLevel: metrics.ALPHA,
		},
	)
)

func init() {
	legacyregistry.MustRegister(tunnelOpenCounter)
	legacyregistry.MustRegister(tunnelOpenFailCounter)
}

// TODO: Unit tests for this code, we can spin up a test SSH server with instructions here:
// https://godoc.org/golang.org/x/crypto/ssh#ServerConn
type sshTunnel struct {
	Config  *ssh.ClientConfig
	Host    string
	SSHPort string
	client  *ssh.Client
}

func makeSSHTunnel(user string, signer ssh.Signer, host string) (*sshTunnel, error) {
	config := ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return &sshTunnel{
		Config:  &config,
		Host:    host,
		SSHPort: "22",
	}, nil
}

func (s *sshTunnel) Open() error {
	var err error
	s.client, err = realTimeoutDialer.Dial("tcp", net.JoinHostPort(s.Host, s.SSHPort), s.Config)
	tunnelOpenCounter.Inc()
	if err != nil {
		tunnelOpenFailCounter.Inc()
	}
	return err
}

func (s *sshTunnel) Dial(ctx context.Context, network, address string) (net.Conn, error) {
	if s.client == nil {
		return nil, errors.New("tunnel has not been opened")
	}
	// This Dial method does not allow to pass a context unfortunately
	return s.client.Dial(network, address)
}

func (s *sshTunnel) Close() error {
	if s.client == nil {
		return errors.New("cannot close tunnel. Tunnel was not open")
	}
	if err := s.client.Close(); err != nil {
		return err
	}
	return nil
}

// Interface to allow mocking of ssh.Dial, for testing SSH
type sshDialer interface {
	Dial(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error)
}

// Real implementation of sshDialer
type realSSHDialer struct{}

var _ sshDialer = &realSSHDialer{}

func (d *realSSHDialer) Dial(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error) {
	conn, err := net.DialTimeout(network, addr, config.Timeout)
	if err != nil {
		return nil, err
	}
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	c, chans, reqs, err := ssh.NewClientConn(conn, addr, config)
	if err != nil {
		return nil, err
	}
	conn.SetReadDeadline(time.Time{})
	return ssh.NewClient(c, chans, reqs), nil
}

// timeoutDialer wraps an sshDialer with a timeout around Dial(). The golang
// ssh library can hang indefinitely inside the Dial() call (see issue #23835).
// Wrapping all Dial() calls with a conservative timeout provides safety against
// getting stuck on that.
type timeoutDialer struct {
	dialer  sshDialer
	timeout time.Duration
}

// 150 seconds is longer than the underlying default TCP backoff delay (127
// seconds). This timeout is only intended to catch otherwise uncaught hangs.
const sshDialTimeout = 150 * time.Second

var realTimeoutDialer sshDialer = &timeoutDialer{&realSSHDialer{}, sshDialTimeout}

func (d *timeoutDialer) Dial(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error) {
	config.Timeout = d.timeout
	return d.dialer.Dial(network, addr, config)
}

// RunSSHCommand returns the stdout, stderr, and exit code from running cmd on
// host as specific user, along with any SSH-level error.
// If user=="", it will default (like SSH) to os.Getenv("USER")
func RunSSHCommand(cmd, user, host string, signer ssh.Signer) (string, string, int, error) {
	return runSSHCommand(realTimeoutDialer, cmd, user, host, signer, true)
}

// Internal implementation of runSSHCommand, for testing
func runSSHCommand(dialer sshDialer, cmd, user, host string, signer ssh.Signer, retry bool) (string, string, int, error) {
	if user == "" {
		user = os.Getenv("USER")
	}
	// Setup the config, dial the server, and open a session.
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := dialer.Dial("tcp", host, config)
	if err != nil && retry {
		err = wait.Poll(5*time.Second, 20*time.Second, func() (bool, error) {
			fmt.Printf("error dialing %s@%s: '%v', retrying\n", user, host, err)
			if client, err = dialer.Dial("tcp", host, config); err != nil {
				return false, err
			}
			return true, nil
		})
	}
	if err != nil {
		return "", "", 0, fmt.Errorf("error getting SSH client to %s@%s: '%v'", user, host, err)
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", "", 0, fmt.Errorf("error creating session to %s@%s: '%v'", user, host, err)
	}
	defer session.Close()

	// Run the command.
	code := 0
	var bout, berr bytes.Buffer
	session.Stdout, session.Stderr = &bout, &berr
	if err = session.Run(cmd); err != nil {
		// Check whether the command failed to run or didn't complete.
		if exiterr, ok := err.(*ssh.ExitError); ok {
			// If we got an ExitError and the exit code is nonzero, we'll
			// consider the SSH itself successful (just that the command run
			// errored on the host).
			if code = exiterr.ExitStatus(); code != 0 {
				err = nil
			}
		} else {
			// Some other kind of error happened (e.g. an IOError); consider the
			// SSH unsuccessful.
			err = fmt.Errorf("failed running `%s` on %s@%s: '%v'", cmd, user, host, err)
		}
	}
	return bout.String(), berr.String(), code, err
}

// MakePrivateKeySignerFromFile generates a private key signer from file.
func MakePrivateKeySignerFromFile(key string) (ssh.Signer, error) {
	// Create an actual signer.
	buffer, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("error reading SSH key %s: '%v'", key, err)
	}
	return MakePrivateKeySignerFromBytes(buffer)
}

// MakePrivateKeySignerFromBytes creates a private key signer from byte slice.
func MakePrivateKeySignerFromBytes(buffer []byte) (ssh.Signer, error) {
	signer, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, fmt.Errorf("error parsing SSH key: '%v'", err)
	}
	return signer, nil
}

// ParsePublicKeyFromFile parses a public key file from the provided file.
func ParsePublicKeyFromFile(keyFile string) (*rsa.PublicKey, error) {
	buffer, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("error reading SSH key %s: '%v'", keyFile, err)
	}
	keyBlock, _ := pem.Decode(buffer)
	if keyBlock == nil {
		return nil, fmt.Errorf("error parsing SSH key %s: 'invalid PEM format'", keyFile)
	}
	key, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing SSH key %s: '%v'", keyFile, err)
	}
	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("SSH key could not be parsed as rsa public key")
	}
	return rsaKey, nil
}

type tunnel interface {
	Open() error
	Close() error
	Dial(ctx context.Context, network, address string) (net.Conn, error)
}

type sshTunnelEntry struct {
	Address string
	Tunnel  tunnel
}

type sshTunnelCreator interface {
	newSSHTunnel(user, keyFile, host string) (tunnel, error)
}

type realTunnelCreator struct{}

func (*realTunnelCreator) newSSHTunnel(user, keyFile, host string) (tunnel, error) {
	signer, err := MakePrivateKeySignerFromFile(keyFile)
	if err != nil {
		return nil, err
	}
	return makeSSHTunnel(user, signer, host)
}

// SecureTunnelList describes a list of SSH tunnels and attributes.
type SecureTunnelList struct {
	entries       []sshTunnelEntry
	adding        map[string]bool
	tunnelCreator sshTunnelCreator
	tunnelsLock   sync.Mutex

	user           string
	keyfile        string
	healthCheckURL *url.URL
}

// NewSecureTunnelList generates a new list of SSH tunnels.
func NewSecureTunnelList(user, keyfile string, healthCheckURL *url.URL, stopChan chan struct{}) *SecureTunnelList {
	l := &SecureTunnelList{
		adding:         make(map[string]bool),
		tunnelCreator:  &realTunnelCreator{},
		user:           user,
		keyfile:        keyfile,
		healthCheckURL: healthCheckURL,
	}
	healthCheckPoll := 1 * time.Minute
	go wait.Until(func() {
		l.tunnelsLock.Lock()
		defer l.tunnelsLock.Unlock()
		// Healthcheck each tunnel every minute
		numTunnels := len(l.entries)
		for i, entry := range l.entries {
			// Stagger healthchecks evenly across duration of healthCheckPoll.
			delay := healthCheckPoll * time.Duration(i) / time.Duration(numTunnels)
			l.delayedHealthCheck(entry, delay)
		}
	}, healthCheckPoll, stopChan)
	return l
}

func (l *SecureTunnelList) delayedHealthCheck(e sshTunnelEntry, delay time.Duration) {
	go func() {
		defer runtime.HandleCrash()
		time.Sleep(delay)
		if err := l.healthCheck(e); err != nil {
			klog.Errorf("Healthcheck failed for tunnel to %q: %v", e.Address, err)
			klog.Infof("Attempting once to re-establish tunnel to %q", e.Address)
			l.removeAndReAdd(e)
		}
	}()
}

func (l *SecureTunnelList) healthCheck(e sshTunnelEntry) error {
	// GET the healthcheck path using the provided tunnel's dial function.
	transport := utilnet.SetTransportDefaults(&http.Transport{
		DialContext: e.Tunnel.Dial,
		// TODO(cjcullen): Plumb real TLS options through.
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// We don't reuse the clients, so disable the keep-alive to properly
		// close the connection.
		DisableKeepAlives: true,
	})
	client := &http.Client{Transport: transport}
	resp, err := client.Get(l.healthCheckURL.String())
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func (l *SecureTunnelList) removeAndReAdd(e sshTunnelEntry) {
	// Find the entry to replace.
	l.tunnelsLock.Lock()
	for i, entry := range l.entries {
		if entry.Tunnel == e.Tunnel {
			l.entries = append(l.entries[:i], l.entries[i+1:]...)
			l.adding[e.Address] = true
			break
		}
	}
	l.tunnelsLock.Unlock()
	if err := e.Tunnel.Close(); err != nil {
		klog.Infof("Failed to close removed tunnel: %v", err)
	}
	go l.createAndAddTunnel(e.Address)
}

// Dial will setup and return an established connection.
func (l *SecureTunnelList) Dial(ctx context.Context, net, addr string) (net.Conn, error) {
	start := time.Now()
	id := mathrand.Int63() // So you can match begins/ends in the log.
	klog.Infof("[%x: %v] Dialing...", id, addr)
	defer func() {
		klog.Infof("[%x: %v] Dialed in %v.", id, addr, time.Since(start))
	}()
	tunnel, err := l.pickTunnel(strings.Split(addr, ":")[0])
	if err != nil {
		return nil, err
	}
	return tunnel.Dial(ctx, net, addr)
}

func (l *SecureTunnelList) pickTunnel(addr string) (tunnel, error) {
	l.tunnelsLock.Lock()
	defer l.tunnelsLock.Unlock()
	if len(l.entries) == 0 {
		return nil, fmt.Errorf("No SSH tunnels currently open. Were the targets able to accept an ssh-key for user %q?", l.user)
	}
	// Prefer same tunnel as kubelet
	for _, entry := range l.entries {
		if entry.Address == addr {
			return entry.Tunnel, nil
		}
	}
	klog.Warningf("SSH tunnel not found for address %q, picking random node", addr)
	n := mathrand.Intn(len(l.entries))
	return l.entries[n].Tunnel, nil
}

// Update reconciles the list's entries with the specified addresses. Existing
// tunnels that are not in addresses are removed from entries and closed in a
// background goroutine. New tunnels specified in addresses are opened in a
// background goroutine and then added to entries.
func (l *SecureTunnelList) Update(addrs []string) {
	haveAddrsMap := make(map[string]bool)
	wantAddrsMap := make(map[string]bool)
	func() {
		l.tunnelsLock.Lock()
		defer l.tunnelsLock.Unlock()
		// Build a map of what we currently have.
		for i := range l.entries {
			haveAddrsMap[l.entries[i].Address] = true
		}
		// Determine any necessary additions.
		for i := range addrs {
			// Add tunnel if it is not in l.entries or l.adding
			if _, ok := haveAddrsMap[addrs[i]]; !ok {
				if _, ok := l.adding[addrs[i]]; !ok {
					l.adding[addrs[i]] = true
					addr := addrs[i]
					go func() {
						defer runtime.HandleCrash()
						// Actually adding tunnel to list will block until lock
						// is released after deletions.
						l.createAndAddTunnel(addr)
					}()
				}
			}
			wantAddrsMap[addrs[i]] = true
		}
		// Determine any necessary deletions.
		var newEntries []sshTunnelEntry
		for i := range l.entries {
			if _, ok := wantAddrsMap[l.entries[i].Address]; !ok {
				tunnelEntry := l.entries[i]
				klog.Infof("Removing tunnel to deleted node at %q", tunnelEntry.Address)
				go func() {
					defer runtime.HandleCrash()
					if err := tunnelEntry.Tunnel.Close(); err != nil {
						klog.Errorf("Failed to close tunnel to %q: %v", tunnelEntry.Address, err)
					}
				}()
			} else {
				newEntries = append(newEntries, l.entries[i])
			}
		}
		l.entries = newEntries
	}()
}

func (l *SecureTunnelList) createAndAddTunnel(addr string) {
	klog.Infof("Trying to add tunnel to %q", addr)
	tunnel, err := l.tunnelCreator.newSSHTunnel(l.user, l.keyfile, addr)
	if err != nil {
		klog.Errorf("Failed to create tunnel for %q: %v", addr, err)
		return
	}
	if err := tunnel.Open(); err != nil {
		klog.Errorf("Failed to open tunnel to %q: %v", addr, err)
		l.tunnelsLock.Lock()
		delete(l.adding, addr)
		l.tunnelsLock.Unlock()
		return
	}
	l.tunnelsLock.Lock()
	l.entries = append(l.entries, sshTunnelEntry{addr, tunnel})
	delete(l.adding, addr)
	l.tunnelsLock.Unlock()
	klog.Infof("Successfully added tunnel for %q", addr)
}

// EncodePrivateKey will PEM encode private key file.
func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}

// EncodePublicKey will PEM encode private key file.
func EncodePublicKey(public *rsa.PublicKey) ([]byte, error) {
	publicBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: publicBytes,
		Type:  "PUBLIC KEY",
	}), nil
}

// EncodeSSHKey will serializes key public key file.
func EncodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}

// GenerateKey will generate RSA keypair.
func GenerateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}
