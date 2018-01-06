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

package validation

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/validation/field"
	defaulttolerationsecondsapi "k8s.io/kubernetes/plugin/pkg/admission/defaulttolerationseconds/apis/defaulttolerationseconds"
)

// ValidateConfiguration validates the configuration.
func ValidateConfiguration(config *defaulttolerationsecondsapi.Configuration) error {
	allErrs := field.ErrorList{}

	defaultTolerationSecondsConfig := config.DefaultTolerationSecondsConfig
	if *defaultTolerationSecondsConfig.DefaultNotReadyTolerationSeconds < 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("defaultNotReadyTolerationSeconds"), *defaultTolerationSecondsConfig.DefaultNotReadyTolerationSeconds, "DefaultNotReadyTolerationSeconds must be not less than 0"))
	}
	if *defaultTolerationSecondsConfig.DefaultUnreachableTolerationSeconds < 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("defaultUnreachableTolerationSeconds"), *defaultTolerationSecondsConfig.DefaultUnreachableTolerationSeconds, "DefaultUnreachableTolerationSeconds must be not less than 0"))
	}
	if len(allErrs) > 0 {
		return fmt.Errorf("invalid config: %v", allErrs)
	}
	return nil
}
