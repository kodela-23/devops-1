// This is a generated file. Do not edit directly.

module k8s.io/controller-manager

go 1.15

require (
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/klog/v2 v2.4.0
)

replace (
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/apiserver => ../apiserver
	k8s.io/client-go => ../client-go
	k8s.io/component-base => ../component-base
	k8s.io/controller-manager => ../controller-manager
	sigs.k8s.io/structured-merge-diff/v4 => github.com/kwiesmueller/structured-merge-diff/v4 v4.0.0-20201112023308-27e979df5bac
)
