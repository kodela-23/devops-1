/*
Copyright 2014 Google Inc. All rights reserved.

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

package conversion

import (
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	"github.com/google/gofuzz"
	"gopkg.in/v1/yaml"
)

var fuzzIters = flag.Int("fuzz_iters", 50, "How many fuzzing iterations to do.")

// Test a weird version/kind embedding format.
type MyWeirdCustomEmbeddedVersionKindField struct {
	ID         string `yaml:"ID,omitempty" json:"ID,omitempty"`
	APIVersion string `json:"myVersionKey,omitempty" yaml:"myVersionKey,omitempty"`
	ObjectKind string `json:"myKindKey,omitempty" yaml:"myKindKey,omitempty"`
	Z          string `yaml:"Z,omitempty" json:"Z,omitempty"`
	Y          uint64 `yaml:"Y,omitempty" json:"Y,omitempty"`
}

type TestType1 struct {
	MyWeirdCustomEmbeddedVersionKindField `json:",inline" yaml:",inline"`
	A                                     string               `yaml:"A,omitempty" json:"A,omitempty"`
	B                                     int                  `yaml:"B,omitempty" json:"B,omitempty"`
	C                                     int8                 `yaml:"C,omitempty" json:"C,omitempty"`
	D                                     int16                `yaml:"D,omitempty" json:"D,omitempty"`
	E                                     int32                `yaml:"E,omitempty" json:"E,omitempty"`
	F                                     int64                `yaml:"F,omitempty" json:"F,omitempty"`
	G                                     uint                 `yaml:"G,omitempty" json:"G,omitempty"`
	H                                     uint8                `yaml:"H,omitempty" json:"H,omitempty"`
	I                                     uint16               `yaml:"I,omitempty" json:"I,omitempty"`
	J                                     uint32               `yaml:"J,omitempty" json:"J,omitempty"`
	K                                     uint64               `yaml:"K,omitempty" json:"K,omitempty"`
	L                                     bool                 `yaml:"L,omitempty" json:"L,omitempty"`
	M                                     map[string]int       `yaml:"M,omitempty" json:"M,omitempty"`
	N                                     map[string]TestType2 `yaml:"N,omitempty" json:"N,omitempty"`
	O                                     *TestType2           `yaml:"O,omitempty" json:"O,omitempty"`
	P                                     []TestType2          `yaml:"Q,omitempty" json:"Q,omitempty"`
}

type TestType2 struct {
	A string `yaml:"A,omitempty" json:"A,omitempty"`
	B int    `yaml:"B,omitempty" json:"B,omitempty"`
}

type ExternalTestType2 struct {
	A string `yaml:"A,omitempty" json:"A,omitempty"`
	B int    `yaml:"B,omitempty" json:"B,omitempty"`
}
type ExternalTestType1 struct {
	MyWeirdCustomEmbeddedVersionKindField `json:",inline" yaml:",inline"`
	A                                     string                       `yaml:"A,omitempty" json:"A,omitempty"`
	B                                     int                          `yaml:"B,omitempty" json:"B,omitempty"`
	C                                     int8                         `yaml:"C,omitempty" json:"C,omitempty"`
	D                                     int16                        `yaml:"D,omitempty" json:"D,omitempty"`
	E                                     int32                        `yaml:"E,omitempty" json:"E,omitempty"`
	F                                     int64                        `yaml:"F,omitempty" json:"F,omitempty"`
	G                                     uint                         `yaml:"G,omitempty" json:"G,omitempty"`
	H                                     uint8                        `yaml:"H,omitempty" json:"H,omitempty"`
	I                                     uint16                       `yaml:"I,omitempty" json:"I,omitempty"`
	J                                     uint32                       `yaml:"J,omitempty" json:"J,omitempty"`
	K                                     uint64                       `yaml:"K,omitempty" json:"K,omitempty"`
	L                                     bool                         `yaml:"L,omitempty" json:"L,omitempty"`
	M                                     map[string]int               `yaml:"M,omitempty" json:"M,omitempty"`
	N                                     map[string]ExternalTestType2 `yaml:"N,omitempty" json:"N,omitempty"`
	O                                     *ExternalTestType2           `yaml:"O,omitempty" json:"O,omitempty"`
	P                                     []ExternalTestType2          `yaml:"Q,omitempty" json:"Q,omitempty"`
}

type ExternalInternalSame struct {
	MyWeirdCustomEmbeddedVersionKindField `json:",inline" yaml:",inline"`
	A                                     TestType2 `yaml:"A,omitempty" json:"A,omitempty"`
}

// TestObjectFuzzer can randomly populate all the above objects.
var TestObjectFuzzer = fuzz.New().NilChance(.5).NumElements(1, 100).Funcs(
	func(j *MyWeirdCustomEmbeddedVersionKindField, c fuzz.Continue) {
		// We have to customize the randomization of MyWeirdCustomEmbeddedVersionKindFields because their
		// APIVersion and Kind must remain blank in memory.
		j.APIVersion = ""
		j.ObjectKind = ""
		j.ID = c.RandString()
	},
	func(u *uint64, c fuzz.Continue) {
		// TODO: Fix JSON/YAML packages and/or write custom encoding
		// for uint64's. Somehow the LS *byte* of this is lost, but
		// only when all 8 bytes are set.
		*u = c.RandUint64() >> 8
	},
	func(u *uint, c fuzz.Continue) {
		// TODO: Fix JSON/YAML packages and/or write custom encoding
		// for uint64's. Somehow the LS *byte* of this is lost, but
		// only when all 8 bytes are set.
		*u = uint(c.RandUint64() >> 8)
	},
)

// Returns a new Scheme set up with the test objects.
func GetTestScheme() *Scheme {
	s := NewScheme()
	// Ordinarily, we wouldn't add TestType2, but because this is a test and
	// both types are from the same package, we need to get it into the system
	// so that converter will match it with ExternalType2.
	s.AddKnownTypes("", &TestType1{}, &TestType2{}, &ExternalInternalSame{})
	s.AddKnownTypes("v1", &ExternalInternalSame{})
	s.AddKnownTypeWithName("v1", "TestType1", &ExternalTestType1{})
	s.AddKnownTypeWithName("v1", "TestType2", &ExternalTestType2{})
	s.InternalVersion = ""
	s.MetaFactory = testMetaFactory{}
	return s
}

type testMetaFactory struct {
	MyWeirdCustomEmbeddedVersionKindField struct {
		APIVersion string `json:"myVersionKey,omitempty" yaml:"myVersionKey,omitempty"`
		ObjectKind string `json:"myKindKey,omitempty" yaml:"myKindKey,omitempty"`
	} `json:",inline" yaml:",inline"`
}

func (testMetaFactory) Update(version, kind string, obj interface{}) error {
	return UpdateVersionAndKind([]string{"MyWeirdCustomEmbeddedVersionKindField"}, "APIVersion", version, "ObjectKind", kind, obj)
}

func (testMetaFactory) Interpret(data []byte) (version, kind string, err error) {
	findKind := struct {
		Version string `json:"myVersionKey,omitempty" yaml:"myVersionKey,omitempty"`
		Kind    string `json:"myKindKey,omitempty" yaml:"myKindKey,omitempty"`
	}{}
	// yaml is a superset of json, so we use it to decode here. That way,
	// we understand both.
	err = yaml.Unmarshal(data, &findKind)
	if err != nil {
		return "", "", fmt.Errorf("couldn't get version/kind: %v", err)
	}
	return findKind.Version, findKind.Kind, nil
}

func objDiff(a, b interface{}) string {
	ab, err := json.Marshal(a)
	if err != nil {
		panic("a")
	}
	bb, err := json.Marshal(b)
	if err != nil {
		panic("b")
	}
	return util.StringDiff(string(ab), string(bb))

	// An alternate diff attempt, in case json isn't showing you
	// the difference. (reflect.DeepEqual makes a distinction between
	// nil and empty slices, for example.)
	return util.StringDiff(
		fmt.Sprintf("%#v", a),
		fmt.Sprintf("%#v", b),
	)
}

func runTest(t *testing.T, source interface{}) {
	name := reflect.TypeOf(source).Elem().Name()
	TestObjectFuzzer.Fuzz(source)

	s := GetTestScheme()
	data, err := s.EncodeToVersion(source, "v1")
	if err != nil {
		t.Errorf("%v: %v (%#v)", name, err, source)
		return
	}
	obj2, err := s.Decode(data)
	if err != nil {
		t.Errorf("%v: %v (%v)", name, err, string(data))
		return
	}
	if !reflect.DeepEqual(source, obj2) {
		t.Errorf("1: %v: diff: %v", name, objDiff(source, obj2))
		return
	}
	obj3 := reflect.New(reflect.TypeOf(source).Elem()).Interface()
	err = s.DecodeInto(data, obj3)
	if err != nil {
		t.Errorf("2: %v: %v", name, err)
		return
	}
	if !reflect.DeepEqual(source, obj3) {
		t.Errorf("3: %v: diff: %v", name, objDiff(source, obj3))
		return
	}
}

func TestCustomMetaFactoryUpdate(t *testing.T) {
	obj := &TestType1{}
	factory := testMetaFactory{}
	if err := factory.Update("1", "Type", obj); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if obj.APIVersion != "1" || obj.ObjectKind != "Type" {
		t.Errorf("object was not updated: %#v", obj)
	}
}

func TestCustomMetaFactoryInterpret(t *testing.T) {
	obj := &TestType1{}
	factory := testMetaFactory{}
	version, kind, err := factory.Interpret([]byte(`{"myVersionKey":"1","myKindKey":"Type"}`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if version != "1" || kind != "Type" {
		t.Errorf("object was not updated: %#v", obj)
	}
}

func TestTypes(t *testing.T) {
	table := []interface{}{
		&TestType1{},
		&ExternalInternalSame{},
	}
	for _, item := range table {
		// Try a few times, since runTest uses random values.
		for i := 0; i < *fuzzIters; i++ {
			runTest(t, item)
		}
	}
}

func TestEncode_NonPtr(t *testing.T) {
	s := GetTestScheme()
	tt := TestType1{A: "I'm not a pointer object"}
	obj := interface{}(tt)
	data, err := s.EncodeToVersion(obj, "v1")
	obj2, err2 := s.Decode(data)
	if err != nil || err2 != nil {
		t.Fatalf("Failure: '%v' '%v'", err, err2)
	}
	if _, ok := obj2.(*TestType1); !ok {
		t.Fatalf("Got wrong type")
	}
	if !reflect.DeepEqual(obj2, &tt) {
		t.Errorf("Expected:\n %#v,\n Got:\n %#v", &tt, obj2)
	}
}

func TestEncode_Ptr(t *testing.T) {
	s := GetTestScheme()
	tt := &TestType1{A: "I am a pointer object"}
	obj := interface{}(tt)
	data, err := s.EncodeToVersion(obj, "v1")
	obj2, err2 := s.Decode(data)
	if err != nil || err2 != nil {
		t.Fatalf("Failure: '%v' '%v'", err, err2)
	}
	if _, ok := obj2.(*TestType1); !ok {
		t.Fatalf("Got wrong type")
	}
	if !reflect.DeepEqual(obj2, tt) {
		t.Errorf("Expected:\n %#v,\n Got:\n %#v", &tt, obj2)
	}
}

func TestBadJSONRejection(t *testing.T) {
	s := GetTestScheme()
	badJSONs := [][]byte{
		[]byte(`{"myVersionKey":"v1"}`),                          // Missing kind
		[]byte(`{"myVersionKey":"v1","myKindKey":"bar"}`),        // Unknown kind
		[]byte(`{"myVersionKey":"bar","myKindKey":"TestType1"}`), // Unknown version
	}
	for _, b := range badJSONs {
		if _, err := s.Decode(b); err == nil {
			t.Errorf("Did not reject bad json: %s", string(b))
		}
	}
	badJSONKindMismatch := []byte(`{"myVersionKey":"v1","myKindKey":"ExternalInternalSame"}`)
	if err := s.DecodeInto(badJSONKindMismatch, &TestType1{}); err == nil {
		t.Errorf("Kind is set but doesn't match the object type: %s", badJSONKindMismatch)
	}
}

func TestBadJSONRejectionForSetInternalVersion(t *testing.T) {
	s := GetTestScheme()
	s.InternalVersion = "v1"
	badJSONs := [][]byte{
		[]byte(`{"myKindKey":"TestType1"}`), // Missing version
	}
	for _, b := range badJSONs {
		if _, err := s.Decode(b); err == nil {
			t.Errorf("Did not reject bad json: %s", string(b))
		}
	}
	badJSONKindMismatch := []byte(`{"myVersionKey":"v1","myKindKey":"ExternalInternalSame"}`)
	if err := s.DecodeInto(badJSONKindMismatch, &TestType1{}); err == nil {
		t.Errorf("Kind is set but doesn't match the object type: %s", badJSONKindMismatch)
	}
}
