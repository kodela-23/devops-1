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

package common

import (
	"fmt"
	"os"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/kubernetes/test/e2e/framework"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("[sig-api-machinery] Secrets", func() {
	f := framework.NewDefaultFramework("secrets")

	It("should be consumable from pods in env vars [Conformance]", func() {
		name := "secret-test-" + string(uuid.NewUUID())
		secret := secretForTest(f.Namespace.Name, name)

		By(fmt.Sprintf("Creating secret with name %s", secret.Name))
		var err error
		if secret, err = f.ClientSet.Core().Secrets(f.Namespace.Name).Create(secret); err != nil {
			framework.Failf("unable to create test secret %s: %v", secret.Name, err)
		}

		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod-secrets-" + string(uuid.NewUUID()),
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:    "secret-env-test",
						Image:   "gcr.io/google_containers/busybox:1.24",
						Command: []string{"sh", "-c", "env"},
						Env: []v1.EnvVar{
							{
								Name: "SECRET_DATA",
								ValueFrom: &v1.EnvVarSource{
									SecretKeyRef: &v1.SecretKeySelector{
										LocalObjectReference: v1.LocalObjectReference{
											Name: name,
										},
										Key: "data-1",
									},
								},
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
			},
		}

		f.TestContainerOutput("consume secrets", pod, 0, []string{
			"SECRET_DATA=value-1",
		})
	})

	It("should be consumable via the environment [Conformance]", func() {
		name := "secret-test-" + string(uuid.NewUUID())
		secret := newEnvFromSecret(f.Namespace.Name, name)
		By(fmt.Sprintf("creating secret %v/%v", f.Namespace.Name, secret.Name))
		var err error
		if secret, err = f.ClientSet.Core().Secrets(f.Namespace.Name).Create(secret); err != nil {
			framework.Failf("unable to create test secret %s: %v", secret.Name, err)
		}

		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod-configmaps-" + string(uuid.NewUUID()),
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:    "env-test",
						Image:   "gcr.io/google_containers/busybox:1.24",
						Command: []string{"sh", "-c", "env"},
						EnvFrom: []v1.EnvFromSource{
							{
								SecretRef: &v1.SecretEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: name}},
							},
							{
								Prefix:    "p_",
								SecretRef: &v1.SecretEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: name}},
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
			},
		}

		f.TestContainerOutput("consume secrets", pod, 0, []string{
			"data_1=value-1", "data_2=value-2", "data_3=value-3",
			"p_data_1=value-1", "p_data_2=value-2", "p_data_3=value-3",
		})
	})
})

func newEnvFromSecret(namespace, name string) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Data: map[string][]byte{
			"data_1": []byte("value-1\n"),
			"data_2": []byte("value-2\n"),
			"data_3": []byte("value-3\n"),
		},
	}
}

func secretForTest(namespace, name string) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Data: map[string][]byte{
			"data-1": []byte("value-1\n"),
			"data-2": []byte("value-2\n"),
			"data-3": []byte("value-3\n"),
		},
	}
}

func doSecretE2EWithoutMapping(f *framework.Framework, defaultMode *int32, secretName string,
	fsGroup *int64, uid *int64) {
	var (
		volumeName      = "secret-volume"
		volumeMountPath = "/etc/secret-volume"
		secret          = secretForTest(f.Namespace.Name, secretName)
	)

	By(fmt.Sprintf("Creating secret with name %s", secret.Name))
	var err error
	if secret, err = f.ClientSet.Core().Secrets(f.Namespace.Name).Create(secret); err != nil {
		framework.Failf("unable to create test secret %s: %v", secret.Name, err)
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod-secrets-" + string(uuid.NewUUID()),
			Namespace: f.Namespace.Name,
		},
		Spec: v1.PodSpec{
			Volumes: []v1.Volume{
				{
					Name: volumeName,
					VolumeSource: v1.VolumeSource{
						Secret: &v1.SecretVolumeSource{
							SecretName: secretName,
						},
					},
				},
			},
			Containers: []v1.Container{
				{
					Name:  "secret-volume-test",
					Image: "gcr.io/google_containers/mounttest:0.8",
					Args: []string{
						"--file_content=/etc/secret-volume/data-1",
						"--file_mode=/etc/secret-volume/data-1"},
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      volumeName,
							MountPath: volumeMountPath,
						},
					},
				},
			},
			RestartPolicy: v1.RestartPolicyNever,
		},
	}

	if defaultMode != nil {
		pod.Spec.Volumes[0].VolumeSource.Secret.DefaultMode = defaultMode
	} else {
		mode := int32(0644)
		defaultMode = &mode
	}

	if fsGroup != nil || uid != nil {
		pod.Spec.SecurityContext = &v1.PodSecurityContext{
			FSGroup:   fsGroup,
			RunAsUser: uid,
		}
	}

	modeString := fmt.Sprintf("%v", os.FileMode(*defaultMode))
	expectedOutput := []string{
		"content of file \"/etc/secret-volume/data-1\": value-1",
		"mode of file \"/etc/secret-volume/data-1\": " + modeString,
	}

	f.TestContainerOutput("consume secrets", pod, 0, expectedOutput)
}

func doSecretE2EWithMapping(f *framework.Framework, mode *int32) {
	var (
		name            = "secret-test-map-" + string(uuid.NewUUID())
		volumeName      = "secret-volume"
		volumeMountPath = "/etc/secret-volume"
		secret          = secretForTest(f.Namespace.Name, name)
	)

	By(fmt.Sprintf("Creating secret with name %s", secret.Name))
	var err error
	if secret, err = f.ClientSet.Core().Secrets(f.Namespace.Name).Create(secret); err != nil {
		framework.Failf("unable to create test secret %s: %v", secret.Name, err)
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-secrets-" + string(uuid.NewUUID()),
		},
		Spec: v1.PodSpec{
			Volumes: []v1.Volume{
				{
					Name: volumeName,
					VolumeSource: v1.VolumeSource{
						Secret: &v1.SecretVolumeSource{
							SecretName: name,
							Items: []v1.KeyToPath{
								{
									Key:  "data-1",
									Path: "new-path-data-1",
								},
							},
						},
					},
				},
			},
			Containers: []v1.Container{
				{
					Name:  "secret-volume-test",
					Image: "gcr.io/google_containers/mounttest:0.8",
					Args: []string{
						"--file_content=/etc/secret-volume/new-path-data-1",
						"--file_mode=/etc/secret-volume/new-path-data-1"},
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      volumeName,
							MountPath: volumeMountPath,
						},
					},
				},
			},
			RestartPolicy: v1.RestartPolicyNever,
		},
	}

	if mode != nil {
		pod.Spec.Volumes[0].VolumeSource.Secret.Items[0].Mode = mode
	} else {
		defaultItemMode := int32(0644)
		mode = &defaultItemMode
	}

	modeString := fmt.Sprintf("%v", os.FileMode(*mode))
	expectedOutput := []string{
		"content of file \"/etc/secret-volume/new-path-data-1\": value-1",
		"mode of file \"/etc/secret-volume/new-path-data-1\": " + modeString,
	}

	f.TestContainerOutput("consume secrets", pod, 0, expectedOutput)
}
