/*
Copyright 2017 The Kubernetes Authors.

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

package topology

import (
	"fmt"

	cadvisorapi "github.com/google/cadvisor/info/v1"
	"k8s.io/kubernetes/pkg/kubelet/cpuset"
)

// CPUDetails is a map from CPU ID to Core ID and Socket ID.
type CPUDetails map[int]CPUInfo

// CPUTopology contains details of node cpu, where :
// CPU  - logical CPU, cadvisor - thread
// Core - physical CPU, cadvisor - Core
// Socket - socket, cadvisor - Node
type CPUTopology struct {
	NumCPUs          int
	NumCores         int
	HyperThreading   bool
	NumSockets       int
	CPUtopoDetails   CPUDetails
	NumReservedCores int
}

// CPUsPerCore returns the number of logical CPUs are associated with
// each core.
func (topo *CPUTopology) CPUsPerCore() int {
	if topo.NumCores == 0 {
		return 0
	}
	return topo.NumCPUs / topo.NumCores
}

// CPUsPerSocket returns the number of logical CPUs are associated with
// each socket.
func (topo *CPUTopology) CPUsPerSocket() int {
	if topo.NumSockets == 0 {
		return 0
	}
	return topo.NumCPUs / topo.NumSockets
}

// CPUInfo contains the socket and core IDs associated with a CPU.
type CPUInfo struct {
	SocketID int
	CoreID   int
}

// KeepOnly returns a new CPUDetails object with only the supplied cpus.
func (d CPUDetails) KeepOnly(cpus cpuset.CPUSet) CPUDetails {
	result := CPUDetails{}
	for cpu, info := range d {
		if cpus.Contains(cpu) {
			result[cpu] = info
		}
	}
	return result
}

// Sockets returns all of the socket IDs associated with the CPUs in this
// CPUDetails.
func (d CPUDetails) Sockets() cpuset.CPUSet {
	result := cpuset.NewCPUSet()
	for _, info := range d {
		result.Add(info.SocketID)
	}
	return result
}

// CPUsInSocket returns all of the logical CPU IDs associated with the
// given socket ID in this CPUDetails.
func (d CPUDetails) CPUsInSocket(id int) cpuset.CPUSet {
	result := cpuset.NewCPUSet()
	for cpu, info := range d {
		if info.SocketID == id {
			result.Add(cpu)
		}
	}
	return result
}

// Cores returns all of the core IDs associated with the CPUs in this
// CPUDetails.
func (d CPUDetails) Cores() cpuset.CPUSet {
	result := cpuset.NewCPUSet()
	for _, info := range d {
		result.Add(info.CoreID)
	}
	return result
}

// CoresInSocket returns all of the core IDs associated with the given
// socket ID in this CPUDetails.
func (d CPUDetails) CoresInSocket(id int) cpuset.CPUSet {
	result := cpuset.NewCPUSet()
	for _, info := range d {
		if info.SocketID == id {
			result.Add(info.CoreID)
		}
	}
	return result
}

// CPUs returns all of the logical CPU IDs in this CPUDetails.
func (d CPUDetails) CPUs() cpuset.CPUSet {
	result := cpuset.NewCPUSet()
	for cpuID := range d {
		result.Add(cpuID)
	}
	return result
}

// CPUsInCore returns all of the logical CPU IDs associated with the
// given core ID in this CPUDetails.
func (d CPUDetails) CPUsInCore(id int) cpuset.CPUSet {
	result := cpuset.NewCPUSet()
	for cpu, info := range d {
		if info.CoreID == id {
			result.Add(cpu)
		}
	}
	return result
}

// Discover returns CPUTopology based on cadvisor node info
func Discover(machineInfo *cadvisorapi.MachineInfo) (*CPUTopology, error) {

	if machineInfo.NumCores == 0 {
		return nil, fmt.Errorf("could not detect number of cpus")
	}

	CPUtopoDetails := CPUDetails{}

	numCPUs := machineInfo.NumCores
	htEnabled := false
	numPhysicalCores := 0
	for _, socket := range machineInfo.Topology {
		numPhysicalCores += len(socket.Cores)
		for _, core := range socket.Cores {
			for _, cpu := range core.Threads {
				CPUtopoDetails[cpu] = CPUInfo{
					CoreID:   core.Id,
					SocketID: socket.Id,
				}
				// a little bit naive
				if !htEnabled && len(core.Threads) != 1 {
					htEnabled = true
				}
			}
		}
	}

	return &CPUTopology{
		NumCPUs:        numCPUs,
		NumSockets:     len(machineInfo.Topology),
		NumCores:       numPhysicalCores,
		HyperThreading: htEnabled,
		CPUtopoDetails: CPUtopoDetails,
	}, nil
}
