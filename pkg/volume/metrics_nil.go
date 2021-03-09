/*
Copyright 2014 The Kubernetes Authors.

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

package volume

var _ StatsProvider = &MetricsNil{}

// MetricsNil represents a StatsProvider that does not support returning
// Stats.  It serves as a placeholder for Volumes that do not yet support
// metrics.
type MetricsNil struct{}

// GetMetrics returns an empty Stats and an error.
// See StatsProvider.GetMetrics
func (*MetricsNil) GetStats() (*Stats, error) {
	return &Stats{}, NewNotSupportedError()
}
