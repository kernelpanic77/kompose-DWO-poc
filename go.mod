module github.com/kernelpanic77/poc

go 1.17

// Required for kompose
replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.9.0
	//github.com/docker/docker => github.com/docker/docker v20.10.0-beta1.0.20201030232932-c2cc352355d4+incompatible
	github.com/containerd/containerd => github.com/containerd/containerd v1.4.1-0.20201030150014-3662dc4c0b12
	github.com/docker/cli => github.com/docker/cli v20.10.0-beta1.0.20201029214301-1d20b15adc38+incompatible
	github.com/docker/docker => github.com/docker/docker v20.10.0-beta1.0.20201030232932-c2cc352355d4+incompatible
	github.com/docker/libcompose => github.com/docker/libcompose v0.4.1-0.20190808084053-143e0f3f1ab9
	github.com/xeipuuv/gojsonschema => github.com/xeipuuv/gojsonschema v1.2.1-0.20201027075954-b076d39a02e5
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8
)

require github.com/kubernetes/kompose v1.26.1

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.15-0.20200113171025-3fe6c5262873 // indirect
	github.com/Microsoft/hcsshim v0.8.7 // indirect
	github.com/containerd/cgroups v0.0.0-20190919134610-bf292b21730f // indirect
	github.com/containerd/containerd v1.3.0 // indirect
	github.com/containerd/continuity v0.0.0-20200228182428-0f16d7a0959c // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/docker/cli v0.0.0-20190711175710-5b38d82aa076 // indirect
	github.com/docker/docker v1.4.2-0.20191101170500-ac7306503d23 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/docker/libcompose v0.4.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/fsouza/go-dockerclient v1.6.5 // indirect
	github.com/go-logr/logr v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/imdario/mergo v0.3.10 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/moby/sys/mount v0.1.1 // indirect
	github.com/moby/sys/mountinfo v0.1.0 // indirect
	github.com/moby/term v0.0.0-20200915141129-7f0af18e79f2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/novln/docker-parser v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/openshift/api v0.0.0-20200803131051-87466835fcc0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0 // indirect
	go.opencensus.io v0.22.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.1.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	k8s.io/api v0.19.0-rc.2 // indirect
	k8s.io/apimachinery v0.19.0-rc.2 // indirect
	k8s.io/klog/v2 v2.2.0 // indirect
	sigs.k8s.io/structured-merge-diff/v3 v3.0.1-0.20200706213357-43c19bbb7fba // indirect
)
