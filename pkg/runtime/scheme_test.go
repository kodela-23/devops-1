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

package runtime_test

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/conversion"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
)

type TypeMeta struct {
	Kind    string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

type InternalSimple struct {
	TypeMeta   `json:",inline" yaml:",inline"`
	TestString string `json:"testString" yaml:"testString"`
}

type ExternalSimple struct {
	TypeMeta   `json:",inline" yaml:",inline"`
	TestString string `json:"testString" yaml:"testString"`
}

func (*InternalSimple) IsAnAPIObject() {}
func (*ExternalSimple) IsAnAPIObject() {}

func TestScheme(t *testing.T) {
	scheme := runtime.NewScheme(conversion.SimpleMetaFactory{[]string{"TypeMeta"}})
	scheme.AddKnownTypeWithName("", "Simple", &InternalSimple{})
	scheme.AddKnownTypeWithName("externalVersion", "Simple", &ExternalSimple{})

	internalToExternalCalls := 0
	externalToInternalCalls := 0

	// Register functions to verify that scope.Meta() gets set correctly.
	err := scheme.AddConversionFuncs(
		func(in *InternalSimple, out *ExternalSimple, scope conversion.Scope) error {
			if e, a := "", scope.Meta().SrcVersion; e != a {
				t.Errorf("Expected '%v', got '%v'", e, a)
			}
			if e, a := "externalVersion", scope.Meta().DestVersion; e != a {
				t.Errorf("Expected '%v', got '%v'", e, a)
			}
			scope.Convert(&in.TypeMeta, &out.TypeMeta, 0)
			scope.Convert(&in.TestString, &out.TestString, 0)
			internalToExternalCalls++
			return nil
		},
		func(in *ExternalSimple, out *InternalSimple, scope conversion.Scope) error {
			if e, a := "externalVersion", scope.Meta().SrcVersion; e != a {
				t.Errorf("Expected '%v', got '%v'", e, a)
			}
			if e, a := "", scope.Meta().DestVersion; e != a {
				t.Errorf("Expected '%v', got '%v'", e, a)
			}
			scope.Convert(&in.TypeMeta, &out.TypeMeta, 0)
			scope.Convert(&in.TestString, &out.TestString, 0)
			externalToInternalCalls++
			return nil
		},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	simple := &InternalSimple{
		TestString: "foo",
	}

	// Test Encode, Decode, and DecodeInto
	obj := runtime.Object(simple)
	data, err := scheme.EncodeToVersion(obj, "externalVersion")
	obj2, err2 := scheme.Decode(data)
	obj3 := &InternalSimple{}
	err3 := scheme.DecodeInto(data, obj3)
	if err != nil || err2 != nil {
		t.Fatalf("Failure: '%v' '%v' '%v'", err, err2, err3)
	}
	if _, ok := obj2.(*InternalSimple); !ok {
		t.Fatalf("Got wrong type")
	}
	if e, a := simple, obj2; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected:\n %#v,\n Got:\n %#v", e, a)
	}
	if e, a := simple, obj3; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected:\n %#v,\n Got:\n %#v", e, a)
	}

	// Test Convert
	external := &ExternalSimple{}
	err = scheme.Convert(simple, external)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if e, a := simple.TestString, external.TestString; e != a {
		t.Errorf("Expected %v, got %v", e, a)
	}

	// Encode and Convert should each have caused an increment.
	if e, a := 2, internalToExternalCalls; e != a {
		t.Errorf("Expected %v, got %v", e, a)
	}
	// Decode and DecodeInto should each have caused an increment.
	if e, a := 2, externalToInternalCalls; e != a {
		t.Errorf("Expected %v, got %v", e, a)
	}
}

func TestBadJSONRejection(t *testing.T) {
	scheme := runtime.NewScheme(conversion.SimpleMetaFactory{})
	badJSONMissingKind := []byte(`{ }`)
	if _, err := scheme.Decode(badJSONMissingKind); err == nil {
		t.Errorf("Did not reject despite lack of kind field: %s", badJSONMissingKind)
	}
	badJSONUnknownType := []byte(`{"kind": "bar"}`)
	if _, err1 := scheme.Decode(badJSONUnknownType); err1 == nil {
		t.Errorf("Did not reject despite use of unknown type: %s", badJSONUnknownType)
	}
	/*badJSONKindMismatch := []byte(`{"kind": "Pod"}`)
	if err2 := DecodeInto(badJSONKindMismatch, &Minion{}); err2 == nil {
		t.Errorf("Kind is set but doesn't match the object type: %s", badJSONKindMismatch)
	}*/
}

type ExtensionA struct {
	runtime.PluginBase `json:",inline" yaml:",inline"`
	TestString         string `json:"testString" yaml:"testString"`
}

type ExtensionB struct {
	runtime.PluginBase `json:",inline" yaml:",inline"`
	TestString         string `json:"testString" yaml:"testString"`
}

type ExternalExtensionType struct {
	TypeMeta  `json:",inline" yaml:",inline"`
	Extension runtime.RawExtension `json:"extension" yaml:"extension"`
}

type InternalExtensionType struct {
	TypeMeta  `json:",inline" yaml:",inline"`
	Extension runtime.EmbeddedObject `json:"extension" yaml:"extension"`
}

func (*ExtensionA) IsAnAPIObject()            {}
func (*ExtensionB) IsAnAPIObject()            {}
func (*ExternalExtensionType) IsAnAPIObject() {}
func (*InternalExtensionType) IsAnAPIObject() {}

func TestExtensionMapping(t *testing.T) {
	scheme := runtime.NewScheme(conversion.SimpleMetaFactory{[]string{"TypeMeta", "PluginBase"}})
	scheme.AddKnownTypeWithName("", "ExtensionType", &InternalExtensionType{})
	scheme.AddKnownTypeWithName("", "A", &ExtensionA{})
	scheme.AddKnownTypeWithName("", "B", &ExtensionB{})
	scheme.AddKnownTypeWithName("testExternal", "ExtensionType", &ExternalExtensionType{})
	scheme.AddKnownTypeWithName("testExternal", "A", &ExtensionA{})
	scheme.AddKnownTypeWithName("testExternal", "B", &ExtensionB{})

	table := []struct {
		obj     runtime.Object
		encoded string
	}{
		{
			&InternalExtensionType{Extension: runtime.EmbeddedObject{&ExtensionA{TestString: "foo"}}},
			`{"kind":"ExtensionType","version":"testExternal","extension":{"kind":"A","testString":"foo"}}`,
		}, {
			&InternalExtensionType{Extension: runtime.EmbeddedObject{&ExtensionB{TestString: "bar"}}},
			`{"kind":"ExtensionType","version":"testExternal","extension":{"kind":"B","testString":"bar"}}`,
		}, {
			&InternalExtensionType{Extension: runtime.EmbeddedObject{nil}},
			`{"kind":"ExtensionType","version":"testExternal","extension":null}`,
		},
	}

	for _, item := range table {
		gotEncoded, err := scheme.EncodeToVersion(item.obj, "testExternal")
		if err != nil {
			t.Errorf("unexpected error '%v' (%#v)", err, item.obj)
		} else if e, a := item.encoded, string(gotEncoded); e != a {
			t.Errorf("expected %v, got %v", e, a)
		}

		gotDecoded, err := scheme.Decode([]byte(item.encoded))
		if err != nil {
			t.Errorf("unexpected error '%v' (%v)", err, item.encoded)
		} else if e, a := item.obj, gotDecoded; !reflect.DeepEqual(e, a) {
			var eEx, aEx runtime.Object
			if obj, ok := e.(*InternalExtensionType); ok {
				eEx = obj.Extension.Object
			}
			if obj, ok := a.(*InternalExtensionType); ok {
				aEx = obj.Extension.Object
			}
			t.Errorf("expected %#v, got %#v (%#v, %#v)", e, a, eEx, aEx)
		}
	}
}

func TestEncode(t *testing.T) {
	scheme := runtime.NewScheme(conversion.SimpleMetaFactory{[]string{"TypeMeta"}})
	scheme.AddKnownTypeWithName("", "Simple", &InternalSimple{})
	scheme.AddKnownTypeWithName("externalVersion", "Simple", &ExternalSimple{})
	codec := runtime.CodecFor(scheme, "externalVersion")
	test := &InternalSimple{
		TestString: "I'm the same",
	}
	obj := runtime.Object(test)
	data, err := codec.Encode(obj)
	obj2, err2 := codec.Decode(data)
	if err != nil || err2 != nil {
		t.Fatalf("Failure: '%v' '%v'", err, err2)
	}
	if _, ok := obj2.(*InternalSimple); !ok {
		t.Fatalf("Got wrong type")
	}
	if !reflect.DeepEqual(obj2, test) {
		t.Errorf("Expected:\n %#v,\n Got:\n %#v", &test, obj2)
	}
}
