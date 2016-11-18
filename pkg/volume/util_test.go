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

package volume

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/util/slice"
)

type testcase struct {
	// Input of the test
	name        string
	existingPod *v1.Pod
	createPod   *v1.Pod
	// eventSequence is list of events that are simulated during recycling. It
	// can be either event generated by a recycler pod or a state change of
	// the pod. (see newPodEvent and newEvent below).
	eventSequence []watch.Event

	// Expected output.
	// expectedEvents is list of events that were sent to the volume that was
	// recycled.
	expectedEvents []mockEvent
	expectedError  string
}

func newPodEvent(eventtype watch.EventType, name string, phase v1.PodPhase, message string) watch.Event {
	return watch.Event{
		Type:   eventtype,
		Object: newPod(name, phase, message),
	}
}

func newEvent(eventtype, message string) watch.Event {
	return watch.Event{
		Type: watch.Added,
		Object: &v1.Event{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: metav1.NamespaceDefault,
			},
			Reason:  "MockEvent",
			Message: message,
			Type:    eventtype,
		},
	}
}

func newPod(name string, phase v1.PodPhase, message string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: metav1.NamespaceDefault,
			Name:      name,
		},
		Status: v1.PodStatus{
			Phase:   phase,
			Message: message,
		},
	}
}

func TestRecyclerPod(t *testing.T) {
	tests := []testcase{
		{
			// Test recycler success with some events
			name:      "RecyclerSuccess",
			createPod: newPod("podRecyclerSuccess", v1.PodPending, ""),
			eventSequence: []watch.Event{
				// Pod gets Running and Succeeded
				newPodEvent(watch.Added, "podRecyclerSuccess", v1.PodPending, ""),
				newEvent(v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerSuccess to 127.0.0.1"),
				newEvent(v1.EventTypeNormal, "pulling image \"gcr.io/google_containers/busybox\""),
				newEvent(v1.EventTypeNormal, "Successfully pulled image \"gcr.io/google_containers/busybox\""),
				newEvent(v1.EventTypeNormal, "Created container with docker id 83d929aeac82"),
				newEvent(v1.EventTypeNormal, "Started container with docker id 83d929aeac82"),
				newPodEvent(watch.Modified, "podRecyclerSuccess", v1.PodRunning, ""),
				newPodEvent(watch.Modified, "podRecyclerSuccess", v1.PodSucceeded, ""),
			},
			expectedEvents: []mockEvent{
				{v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerSuccess to 127.0.0.1"},
				{v1.EventTypeNormal, "pulling image \"gcr.io/google_containers/busybox\""},
				{v1.EventTypeNormal, "Successfully pulled image \"gcr.io/google_containers/busybox\""},
				{v1.EventTypeNormal, "Created container with docker id 83d929aeac82"},
				{v1.EventTypeNormal, "Started container with docker id 83d929aeac82"},
			},
			expectedError: "",
		},
		{
			// Test recycler failure with some events
			name:      "RecyclerFailure",
			createPod: newPod("podRecyclerFailure", v1.PodPending, ""),
			eventSequence: []watch.Event{
				// Pod gets Running and Succeeded
				newPodEvent(watch.Added, "podRecyclerFailure", v1.PodPending, ""),
				newEvent(v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerFailure to 127.0.0.1"),
				newEvent(v1.EventTypeWarning, "Unable to mount volumes for pod \"recycler-for-podRecyclerFailure_default(3c9809e5-347c-11e6-a79b-3c970e965218)\": timeout expired waiting for volumes to attach/mount"),
				newEvent(v1.EventTypeWarning, "Error syncing pod, skipping: timeout expired waiting for volumes to attach/mount for pod \"default\"/\"recycler-for-podRecyclerFailure\". list of unattached/unmounted"),
				newPodEvent(watch.Modified, "podRecyclerFailure", v1.PodRunning, ""),
				newPodEvent(watch.Modified, "podRecyclerFailure", v1.PodFailed, "Pod was active on the node longer than specified deadline"),
			},
			expectedEvents: []mockEvent{
				{v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerFailure to 127.0.0.1"},
				{v1.EventTypeWarning, "Unable to mount volumes for pod \"recycler-for-podRecyclerFailure_default(3c9809e5-347c-11e6-a79b-3c970e965218)\": timeout expired waiting for volumes to attach/mount"},
				{v1.EventTypeWarning, "Error syncing pod, skipping: timeout expired waiting for volumes to attach/mount for pod \"default\"/\"recycler-for-podRecyclerFailure\". list of unattached/unmounted"},
			},
			expectedError: "Pod was active on the node longer than specified deadline",
		},
		{
			// Recycler pod gets deleted
			name:      "RecyclerDeleted",
			createPod: newPod("podRecyclerDeleted", v1.PodPending, ""),
			eventSequence: []watch.Event{
				// Pod gets Running and Succeeded
				newPodEvent(watch.Added, "podRecyclerDeleted", v1.PodPending, ""),
				newEvent(v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerDeleted to 127.0.0.1"),
				newPodEvent(watch.Deleted, "podRecyclerDeleted", v1.PodPending, ""),
			},
			expectedEvents: []mockEvent{
				{v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerDeleted to 127.0.0.1"},
			},
			expectedError: "recycler pod was deleted",
		},
		{
			// Another recycler pod is already running
			name:        "RecyclerRunning",
			existingPod: newPod("podOldRecycler", v1.PodRunning, ""),
			createPod:   newPod("podNewRecycler", v1.PodFailed, "mock message"),
			eventSequence: []watch.Event{
				// Old pod succeeds
				newPodEvent(watch.Modified, "podOldRecycler", v1.PodSucceeded, ""),
			},
			// No error = old pod succeeded. If the new pod was used, there
			// would be error with "mock message".
			expectedError: "",
		},
		{
			// Another recycler pod is already running and fails
			name:        "FailedRecyclerRunning",
			existingPod: newPod("podOldRecycler", v1.PodRunning, ""),
			createPod:   newPod("podNewRecycler", v1.PodFailed, "mock message"),
			eventSequence: []watch.Event{
				// Old pod failure
				newPodEvent(watch.Modified, "podOldRecycler", v1.PodFailed, "Pod was active on the node longer than specified deadline"),
			},
			// If the new pod was used, there would be error with "mock message".
			expectedError: "Pod was active on the node longer than specified deadline",
		},
	}

	for _, test := range tests {
		t.Logf("Test %q", test.name)
		client := &mockRecyclerClient{
			events: test.eventSequence,
			pod:    test.existingPod,
		}
		err := internalRecycleVolumeByWatchingPodUntilCompletion(test.createPod.Name, test.createPod, client)
		receivedError := ""
		if err != nil {
			receivedError = err.Error()
		}
		if receivedError != test.expectedError {
			t.Errorf("Test %q failed, expected error %q, got %q", test.name, test.expectedError, receivedError)
			continue
		}
		if !client.deletedCalled {
			t.Errorf("Test %q failed, expected deferred client.Delete to be called on recycler pod", test.name)
			continue
		}
		for i, expectedEvent := range test.expectedEvents {
			if len(client.receivedEvents) <= i {
				t.Errorf("Test %q failed, expected event %d: %q not received", test.name, i, expectedEvent.message)
				continue
			}
			receivedEvent := client.receivedEvents[i]
			if expectedEvent.eventtype != receivedEvent.eventtype {
				t.Errorf("Test %q failed, event %d does not match: expected eventtype %q, got %q", test.name, i, expectedEvent.eventtype, receivedEvent.eventtype)
			}
			if expectedEvent.message != receivedEvent.message {
				t.Errorf("Test %q failed, event %d does not match: expected message %q, got %q", test.name, i, expectedEvent.message, receivedEvent.message)
			}
		}
		for i := len(test.expectedEvents); i < len(client.receivedEvents); i++ {
			t.Errorf("Test %q failed, unexpected event received: %s, %q", test.name, client.receivedEvents[i].eventtype, client.receivedEvents[i].message)
		}
	}
}

type mockRecyclerClient struct {
	pod            *v1.Pod
	deletedCalled  bool
	receivedEvents []mockEvent
	events         []watch.Event
}

type mockEvent struct {
	eventtype, message string
}

func (c *mockRecyclerClient) CreatePod(pod *v1.Pod) (*v1.Pod, error) {
	if c.pod == nil {
		c.pod = pod
		return c.pod, nil
	}
	// Simulate "already exists" error
	return nil, errors.NewAlreadyExists(api.Resource("pods"), pod.Name)
}

func (c *mockRecyclerClient) GetPod(name, namespace string) (*v1.Pod, error) {
	if c.pod != nil {
		return c.pod, nil
	} else {
		return nil, fmt.Errorf("pod does not exist")
	}
}

func (c *mockRecyclerClient) DeletePod(name, namespace string) error {
	c.deletedCalled = true
	return nil
}

func (c *mockRecyclerClient) WatchPod(name, namespace string, stopChannel chan struct{}) (<-chan watch.Event, error) {
	eventCh := make(chan watch.Event, 0)
	go func() {
		for _, e := range c.events {
			eventCh <- e
		}
	}()
	return eventCh, nil
}

func (c *mockRecyclerClient) Event(eventtype, message string) {
	c.receivedEvents = append(c.receivedEvents, mockEvent{eventtype, message})
}

func TestCalculateTimeoutForVolume(t *testing.T) {
	pv := &v1.PersistentVolume{
		Spec: v1.PersistentVolumeSpec{
			Capacity: v1.ResourceList{
				v1.ResourceName(v1.ResourceStorage): resource.MustParse("500M"),
			},
		},
	}

	timeout := CalculateTimeoutForVolume(50, 30, pv)
	if timeout != 50 {
		t.Errorf("Expected 50 for timeout but got %v", timeout)
	}

	pv.Spec.Capacity[v1.ResourceStorage] = resource.MustParse("2Gi")
	timeout = CalculateTimeoutForVolume(50, 30, pv)
	if timeout != 60 {
		t.Errorf("Expected 60 for timeout but got %v", timeout)
	}

	pv.Spec.Capacity[v1.ResourceStorage] = resource.MustParse("150Gi")
	timeout = CalculateTimeoutForVolume(50, 30, pv)
	if timeout != 4500 {
		t.Errorf("Expected 4500 for timeout but got %v", timeout)
	}
}

func TestGenerateVolumeName(t *testing.T) {

	// Normal operation, no truncate
	v1 := GenerateVolumeName("kubernetes", "pv-cinder-abcde", 255)
	if v1 != "kubernetes-dynamic-pv-cinder-abcde" {
		t.Errorf("Expected kubernetes-dynamic-pv-cinder-abcde, got %s", v1)
	}

	// Truncate trailing "6789-dynamic"
	prefix := strings.Repeat("0123456789", 9) // 90 characters prefix + 8 chars. of "-dynamic"
	v2 := GenerateVolumeName(prefix, "pv-cinder-abcde", 100)
	expect := prefix[:84] + "-pv-cinder-abcde"
	if v2 != expect {
		t.Errorf("Expected %s, got %s", expect, v2)
	}

	// Truncate really long cluster name
	prefix = strings.Repeat("0123456789", 1000) // 10000 characters prefix
	v3 := GenerateVolumeName(prefix, "pv-cinder-abcde", 100)
	if v3 != expect {
		t.Errorf("Expected %s, got %s", expect, v3)
	}
}

func TestMountOptionFromSpec(t *testing.T) {
	scenarios := map[string]struct {
		volume            *Spec
		expectedMountList []string
		systemOptions     []string
	}{
		"volume-with-mount-options": {
			volume: createVolumeSpecWithMountOption("good-mount-opts", "ro,nfsvers=3", v1.PersistentVolumeSpec{
				PersistentVolumeSource: v1.PersistentVolumeSource{
					NFS: &v1.NFSVolumeSource{Server: "localhost", Path: "/srv", ReadOnly: false},
				},
			}),
			expectedMountList: []string{"ro", "nfsvers=3"},
			systemOptions:     nil,
		},
		"volume-with-bad-mount-options": {
			volume: createVolumeSpecWithMountOption("good-mount-opts", "", v1.PersistentVolumeSpec{
				PersistentVolumeSource: v1.PersistentVolumeSource{
					NFS: &v1.NFSVolumeSource{Server: "localhost", Path: "/srv", ReadOnly: false},
				},
			}),
			expectedMountList: []string{},
			systemOptions:     nil,
		},
		"vol-with-sys-opts": {
			volume: createVolumeSpecWithMountOption("good-mount-opts", "ro,nfsvers=3", v1.PersistentVolumeSpec{
				PersistentVolumeSource: v1.PersistentVolumeSource{
					NFS: &v1.NFSVolumeSource{Server: "localhost", Path: "/srv", ReadOnly: false},
				},
			}),
			expectedMountList: []string{"ro", "nfsvers=3", "fsid=100", "hard"},
			systemOptions:     []string{"fsid=100", "hard"},
		},
		"vol-with-sys-opts-with-dup": {
			volume: createVolumeSpecWithMountOption("good-mount-opts", "ro,nfsvers=3", v1.PersistentVolumeSpec{
				PersistentVolumeSource: v1.PersistentVolumeSource{
					NFS: &v1.NFSVolumeSource{Server: "localhost", Path: "/srv", ReadOnly: false},
				},
			}),
			expectedMountList: []string{"ro", "nfsvers=3", "fsid=100"},
			systemOptions:     []string{"fsid=100", "ro"},
		},
	}

	for name, scenario := range scenarios {
		mountOptions := MountOptionFromSpec(scenario.volume, scenario.systemOptions...)
		if !reflect.DeepEqual(slice.SortStrings(mountOptions), slice.SortStrings(scenario.expectedMountList)) {
			t.Errorf("for %s expected mount options : %v got %v", name, scenario.expectedMountList, mountOptions)
		}
	}
}

func createVolumeSpecWithMountOption(name string, mountOptions string, spec v1.PersistentVolumeSpec) *Spec {
	annotations := map[string]string{
		MountOptionAnnotation: mountOptions,
	}
	objMeta := metav1.ObjectMeta{
		Name:        name,
		Annotations: annotations,
	}

	pv := &v1.PersistentVolume{
		ObjectMeta: objMeta,
		Spec:       spec,
	}
	return &Spec{PersistentVolume: pv}
}

func checkFnv32(t *testing.T, s string, expected int) {
	h := fnv.New32()
	h.Write([]byte(s))
	h.Sum32()

	if int(h.Sum32()) != expected {
		t.Fatalf("hash of %q was %v, expected %v", s, h.Sum32(), expected)
	}
}

func TestChooseZoneForVolume(t *testing.T) {
	checkFnv32(t, "henley", 1180403676)
	// 1180403676 mod 3 == 0, so the offset from "henley" is 0, which makes it easier to verify this by inspection

	// A few others
	checkFnv32(t, "henley-", 2652299129)
	checkFnv32(t, "henley-a", 1459735322)
	checkFnv32(t, "", 2166136261)

	tests := []struct {
		Zones      []string
		VolumeName string
		Expected   string
	}{
		// Test for PVC names that don't have a dash
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "henley",
			Expected:   "a", // hash("henley") == 0
		},
		// Tests for PVC names that end in - number, but don't look like statefulset PVCs
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "henley-0",
			Expected:   "a", // hash("henley") == 0
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "henley-1",
			Expected:   "b", // hash("henley") + 1 == 1
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "henley-2",
			Expected:   "c", // hash("henley") + 2 == 2
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "henley-3",
			Expected:   "a", // hash("henley") + 3 == 3 === 0 mod 3
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "henley-4",
			Expected:   "b", // hash("henley") + 4 == 4 === 1 mod 3
		},
		// Tests for PVC names that are edge cases
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "henley-",
			Expected:   "c", // hash("henley-") = 2652299129 === 2 mod 3
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "henley-a",
			Expected:   "c", // hash("henley-a") = 1459735322 === 2 mod 3
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium--1",
			Expected:   "c", // hash("") + 1 == 2166136261 + 1 === 2 mod 3
		},
		// Tests for PVC names for simple StatefulSet cases
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-henley-1",
			Expected:   "b", // hash("henley") + 1 == 1
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "loud-henley-1",
			Expected:   "b", // hash("henley") + 1 == 1
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "quiet-henley-2",
			Expected:   "c", // hash("henley") + 2 == 2
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-henley-2",
			Expected:   "c", // hash("henley") + 2 == 2
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-henley-3",
			Expected:   "a", // hash("henley") + 3 == 3 === 0 mod 3
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-henley-4",
			Expected:   "b", // hash("henley") + 4 == 4 === 1 mod 3
		},
		// Tests for statefulsets (or claims) with dashes in the names
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-alpha-henley-2",
			Expected:   "c", // hash("henley") + 2 == 2
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-beta-henley-3",
			Expected:   "a", // hash("henley") + 3 == 3 === 0 mod 3
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-gamma-henley-4",
			Expected:   "b", // hash("henley") + 4 == 4 === 1 mod 3
		},
		// Tests for statefulsets name ending in -
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-henley--2",
			Expected:   "a", // hash("") + 2 == 0 mod 3
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-henley--3",
			Expected:   "b", // hash("") + 3 == 1 mod 3
		},
		{
			Zones:      []string{"a", "b", "c"},
			VolumeName: "medium-henley--4",
			Expected:   "c", // hash("") + 4 == 2 mod 3
		},
	}

	for _, test := range tests {
		zonesSet := sets.NewString(test.Zones...)

		actual := ChooseZoneForVolume(zonesSet, test.VolumeName)

		for actual != test.Expected {
			t.Errorf("Test %v failed, expected zone %q, actual %q", test, test.Expected, actual)
		}
	}
}

func TestZonesToSet(t *testing.T) {
	functionUnderTest := "ZonesToSet"
	// First part: want an error
	sliceOfZones := []string{"", ",", "us-east-1a, , us-east-1d", ", us-west-1b", "us-west-2b,"}
	for _, zones := range sliceOfZones {
		if got, err := ZonesToSet(zones); err == nil {
			t.Errorf("%v(%v) returned (%v), want (%v)", functionUnderTest, zones, got, "an error")
		}
	}

	// Second part: want no error
	tests := []struct {
		zones string
		want  sets.String
	}{
		{
			zones: "us-east-1a",
			want:  sets.String{"us-east-1a": sets.Empty{}},
		},
		{
			zones: "us-east-1a, us-west-2a",
			want: sets.String{
				"us-east-1a": sets.Empty{},
				"us-west-2a": sets.Empty{},
			},
		},
	}
	for _, tt := range tests {
		if got, err := ZonesToSet(tt.zones); err != nil || !got.Equal(tt.want) {
			t.Errorf("%v(%v) returned (%v), want (%v)", functionUnderTest, tt.zones, got, tt.want)
		}
	}
}

func TestValidateZone(t *testing.T) {
	functionUnderTest := "ValidateZone"

	// First part: want an error
	errCases := []string{"", " 	 	 "}
	for _, errCase := range errCases {
		if got := ValidateZone(errCase); got == nil {
			t.Errorf("%v(%v) returned (%v), want (%v)", functionUnderTest, errCase, got, "an error")
		}
	}

	// Second part: want no error
	succCases := []string{" us-east-1a	"}
	for _, succCase := range succCases {
		if got := ValidateZone(succCase); got != nil {
			t.Errorf("%v(%v) returned (%v), want (%v)", functionUnderTest, succCase, got, nil)
		}
	}
}

func TestValidatePVCSelector(t *testing.T) {
	functionUnderTest := "validatePVCSelector"
	// First part: want no error
	succTests := []struct {
		pvc       v1.PersistentVolumeClaim
		wantEmpty bool
	}{
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
			},
			wantEmpty: true,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: nil,
				},
			},
			wantEmpty: true,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{},
					},
				},
			},
			wantEmpty: true,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{},
					},
				},
			},
			wantEmpty: true,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      metav1.LabelZoneFailureDomain,
								Operator: metav1.LabelSelectorOpIn,
								Values:   []string{"us-east-1a", "us-east-1b"},
							},
						},
					},
				},
			},
			wantEmpty: false,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      metav1.LabelZoneRegion,
								Operator: metav1.LabelSelectorOpNotIn,
								Values:   []string{"us-east-1a", "us-east-1b"},
							},
						},
					},
				},
			},
			wantEmpty: false,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{metav1.LabelZoneFailureDomain: "us-east-1a"},
					},
				},
			},
			wantEmpty: false,
		},
	}
	for _, succTest := range succTests {
		if empty, err := validatePVCSelector(&succTest.pvc); err != nil {
			t.Errorf("%v(%v) returned (%v, %v), want (%v, %v)", functionUnderTest, succTest.pvc, empty, err.Error(), succTest.wantEmpty, nil)
		} else if empty != succTest.wantEmpty {
			t.Errorf("%v(%v) returned (%v, %v), want (%v, %v)", functionUnderTest, succTest.pvc, empty, err, succTest.wantEmpty, nil)
		}
	}

	// Second part: want an error
	errCases := []v1.PersistentVolumeClaim{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
			Spec: v1.PersistentVolumeClaimSpec{
				Selector: &metav1.LabelSelector{
					MatchExpressions: []metav1.LabelSelectorRequirement{
						{
							Key:      "key2",
							Operator: "In",
							Values:   []string{"value1", "value2"},
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
			Spec: v1.PersistentVolumeClaimSpec{
				Selector: &metav1.LabelSelector{
					MatchExpressions: []metav1.LabelSelectorRequirement{
						{
							Key:      metav1.LabelZoneFailureDomain,
							Operator: metav1.LabelSelectorOpExists,
							Values:   []string{"value1", "value2"},
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
			Spec: v1.PersistentVolumeClaimSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"foo": "bar"},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
			Spec: v1.PersistentVolumeClaimSpec{
				Selector: &metav1.LabelSelector{
					MatchExpressions: []metav1.LabelSelectorRequirement{
						{
							Key:      metav1.LabelZoneFailureDomain,
							Operator: metav1.LabelSelectorOpIn,
							Values:   []string{},
						},
					},
				},
			},
		},
	}
	for _, errCase := range errCases {
		if empty, err := validatePVCSelector(&errCase); err == nil {
			t.Errorf("%v(%v) returned (%v, %v), want (%v, %v)", functionUnderTest, errCase, empty, err, false, "an error")
		} else if empty != false {
			t.Errorf("%v(%v) returned (%v, %v), want (%v, %v)", functionUnderTest, errCase, empty, err.Error(), false, "an error")
		}
	}
}

func TestGetPVCMatchLabel(t *testing.T) {
	functionUnderTest := "getPVCMatchLabel"
	// First part: want no error
	succTests := []struct {
		pvc  v1.PersistentVolumeClaim
		key  string
		want string
	}{
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{metav1.LabelZoneFailureDomain: "us-east-1a"},
					},
				},
			},
			key:  metav1.LabelZoneFailureDomain,
			want: "us-east-1a",
		},
	}
	for _, succTest := range succTests {
		if zone, err := getPVCMatchLabel(&succTest.pvc, succTest.key); err != nil {
			t.Errorf("%v(%v, %v) returned (%v, %v), want (%v, %v)", functionUnderTest, succTest.pvc, succTest.key, zone, err.Error(), succTest.want, nil)
		}
	}

	// Second part: want an error
	errCases := []struct {
		pvc v1.PersistentVolumeClaim
		key string
	}{
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
			},
			key: metav1.LabelZoneFailureDomain,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{},
					},
				},
			},
			key: metav1.LabelZoneFailureDomain,
		},
	}
	for _, errCase := range errCases {
		if zone, err := getPVCMatchLabel(&errCase.pvc, errCase.key); err == nil {
			t.Errorf("%v(%v, %v) returned (%v, %v), want (%v, %v)", functionUnderTest, errCase.pvc, errCase.key, zone, err, "", "an error")
		}
	}
}

func isEqualSliceOfSets(s1, s2 []sets.String) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if !s1[i].Equal(s2[i]) {
			return false
		}
	}
	return true
}

func TestGetPVCMatchExpression(t *testing.T) {
	functionUnderTest := "getPVCMatchExpression"
	emptySet := make(sets.String)
	// First part: want no error
	succTests := []struct {
		pvc      v1.PersistentVolumeClaim
		key      string
		operator metav1.LabelSelectorOperator
		want     []sets.String
	}{
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      metav1.LabelZoneFailureDomain,
								Operator: metav1.LabelSelectorOpIn,
								Values:   []string{"us-east-1a", "us-east-1b"},
							},
						},
					},
				},
			},
			key:      metav1.LabelZoneFailureDomain,
			operator: metav1.LabelSelectorOpIn,
			want:     []sets.String{{"us-east-1a": sets.Empty{}, "us-east-1b": sets.Empty{}}},
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      metav1.LabelZoneFailureDomain,
								Operator: metav1.LabelSelectorOpIn,
								Values:   []string{"us-east-1a", "us-east-2a", "us-east-3a"},
							},
							{
								Key:      metav1.LabelZoneFailureDomain,
								Operator: metav1.LabelSelectorOpIn,
								Values:   []string{"us-east-3a", "us-east-4a"},
							},
						},
					},
				},
			},
			key:      metav1.LabelZoneFailureDomain,
			operator: metav1.LabelSelectorOpIn,
			want:     []sets.String{{"us-east-1a": sets.Empty{}, "us-east-2a": sets.Empty{}, "us-east-3a": sets.Empty{}}, {"us-east-3a": sets.Empty{}, "us-east-4a": sets.Empty{}}},
		},
	}
	for _, succTest := range succTests {
		if zones, err := getPVCMatchExpression(&succTest.pvc, succTest.key, succTest.operator); err != nil {
			t.Errorf("%v(%v, %v, %v) returned (%v, %v), want (%v, %v)", functionUnderTest, succTest.pvc, succTest.key, succTest.operator, zones, err.Error(), succTest.want, nil)
		} else if !isEqualSliceOfSets(zones, succTest.want) {
			t.Errorf("%v(%v, %v, %v) returned (%v, %v), want (%v, %v)", functionUnderTest, succTest.pvc, succTest.key, succTest.operator, zones, err, succTest.want, nil)
		}
	}

	// Second part: want an error
	errCases := []struct {
		pvc      v1.PersistentVolumeClaim
		key      string
		operator metav1.LabelSelectorOperator
	}{
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
			},
			key:      metav1.LabelZoneFailureDomain,
			operator: metav1.LabelSelectorOpIn,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{},
					},
				},
			},
			key:      metav1.LabelZoneRegion,
			operator: metav1.LabelSelectorOpNotIn,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      "foo",
								Operator: metav1.LabelSelectorOpIn,
								Values:   []string{"us-east-1a", "us-east-1b"},
							},
						},
					},
				},
			},
			key:      metav1.LabelZoneRegion,
			operator: metav1.LabelSelectorOpIn,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      metav1.LabelZoneFailureDomain,
								Operator: "foo",
								Values:   []string{"us-east-1a", "us-east-1b"},
							},
						},
					},
				},
			},
			key:      metav1.LabelZoneFailureDomain,
			operator: metav1.LabelSelectorOpIn,
		},
		{
			pvc: v1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{Name: "pvc", Namespace: "foo"},
				Spec: v1.PersistentVolumeClaimSpec{
					Selector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      metav1.LabelZoneFailureDomain,
								Operator: metav1.LabelSelectorOpIn,
								Values:   []string{},
							},
						},
					},
				},
			},
			key:      metav1.LabelZoneFailureDomain,
			operator: metav1.LabelSelectorOpIn,
		},
	}
	for _, errCase := range errCases {
		if zones, err := getPVCMatchExpression(&errCase.pvc, errCase.key, errCase.operator); err == nil {
			t.Errorf("%v(%v, %v, %v) returned (%v, %v), want (%v, %v)", functionUnderTest, errCase.pvc, errCase.key, errCase.operator, zones, err, emptySet, "an error")
		}
	}
}
