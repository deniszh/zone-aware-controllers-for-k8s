module github.com/aws/zone-aware-controllers-for-k8s

go 1.23.0

toolchain go1.24.1

require (
	github.com/aws/aws-sdk-go-v2 v1.16.3
	github.com/aws/aws-sdk-go-v2/config v1.15.5
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.18.2
	github.com/aws/aws-sdk-go-v2/service/eks v1.21.0
	github.com/go-logr/logr v1.2.3
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.24.1
	github.com/prometheus/client_golang v1.14.0
	github.com/stretchr/testify v1.8.0
	k8s.io/api v0.26.0
	k8s.io/apimachinery v0.26.0
	k8s.io/apiserver v0.26.0
	k8s.io/client-go v0.26.0
	k8s.io/klog/v2 v2.80.1
	k8s.io/utils v0.0.0-20221128185143-99ec85e7a448
	sigs.k8s.io/aws-iam-authenticator v0.6.3
	sigs.k8s.io/controller-runtime v0.14.1
	sigs.k8s.io/e2e-framework v0.1.0
)

require (
	github.com/aws/aws-sdk-go v1.44.145 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.12.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.4 // indirect
	github.com/aws/smithy-go v1.11.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/zapr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/gofrs/flock v0.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.36.0 // indirect
	golang.org/x/oauth2 v0.27.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/term v0.29.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.26.0 // indirect
	k8s.io/component-base v0.26.0 // indirect
	k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
