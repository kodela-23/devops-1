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

package defaulttolerationseconds

import (
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/helper"
	defaulttolerationsecondsapi "k8s.io/kubernetes/plugin/pkg/admission/defaulttolerationseconds/apis/defaulttolerationseconds"
	"k8s.io/kubernetes/plugin/pkg/scheduler/algorithm"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("DefaultTolerationSeconds",
		func(config io.Reader) (admission.Interface, error) {
			// load the configuration provided (if any)
			configuration, err := LoadConfiguration(config)
			if err != nil {
				return nil, err
			}
			return NewDefaultTolerationSeconds(configuration), nil
		})
}

// Plugin contains the client used by the admission controller
// It will add default tolerations for every pod
// that tolerate taints `notReady:NoExecute` and `unreachable:NoExecute`,
// with tolerationSeconds from configuration, or with default
// tolerationSeconds of 300s if not configured or misconfigured.
// If the pod already specifies a toleration for taint `notReady:NoExecute`
// or `unreachable:NoExecute`, the plugin won't touch it.
type Plugin struct {
	*admission.Handler
	notReadyTolerationSeconds    *int64
	unreachableTolerationSeconds *int64
}

var _ admission.MutationInterface = &Plugin{}

// NewDefaultTolerationSeconds creates a new instance of the DefaultTolerationSeconds admission controller
func NewDefaultTolerationSeconds(config *defaulttolerationsecondsapi.Configuration) *Plugin {
	return &Plugin{
		Handler:                      admission.NewHandler(admission.Create, admission.Update),
		notReadyTolerationSeconds:    config.DefaultTolerationSecondsConfig.DefaultNotReadyTolerationSeconds,
		unreachableTolerationSeconds: config.DefaultTolerationSecondsConfig.DefaultUnreachableTolerationSeconds,
	}
}

// Admit makes an admission decision based on the request attributes
func (p *Plugin) Admit(attributes admission.Attributes) (err error) {
	if attributes.GetResource().GroupResource() != api.Resource("pods") {
		return nil
	}

	if len(attributes.GetSubresource()) > 0 {
		// only run the checks below on pods proper and not subresources
		return nil
	}

	pod, ok := attributes.GetObject().(*api.Pod)
	if !ok {
		return errors.NewBadRequest(fmt.Sprintf("expected *api.Pod but got %T", attributes.GetObject()))
	}

	tolerations := pod.Spec.Tolerations

	toleratesNodeNotReady := false
	toleratesNodeUnreachable := false
	for _, toleration := range tolerations {
		if (toleration.Key == algorithm.TaintNodeNotReady || len(toleration.Key) == 0) &&
			(toleration.Effect == api.TaintEffectNoExecute || len(toleration.Effect) == 0) {
			toleratesNodeNotReady = true
		}

		if (toleration.Key == algorithm.TaintNodeUnreachable || len(toleration.Key) == 0) &&
			(toleration.Effect == api.TaintEffectNoExecute || len(toleration.Effect) == 0) {
			toleratesNodeUnreachable = true
		}
	}

	// no change is required, return immediately
	if toleratesNodeNotReady && toleratesNodeUnreachable {
		return nil
	}

	if !toleratesNodeNotReady {
		helper.AddOrUpdateTolerationInPod(pod, &api.Toleration{
			Key:               algorithm.TaintNodeNotReady,
			Operator:          api.TolerationOpExists,
			Effect:            api.TaintEffectNoExecute,
			TolerationSeconds: p.notReadyTolerationSeconds,
		})
	}

	if !toleratesNodeUnreachable {
		helper.AddOrUpdateTolerationInPod(pod, &api.Toleration{
			Key:               algorithm.TaintNodeUnreachable,
			Operator:          api.TolerationOpExists,
			Effect:            api.TaintEffectNoExecute,
			TolerationSeconds: p.unreachableTolerationSeconds,
		})
	}
	return nil
}
