module github.com/openshift/osd-network-verifier

go 1.16

require (
	github.com/aws/aws-sdk-go v1.42.20
	github.com/aws/aws-sdk-go-v2 v1.11.2
	github.com/aws/aws-sdk-go-v2/config v1.10.3
	github.com/aws/aws-sdk-go-v2/credentials v1.6.4
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.24.0
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0
	github.com/openshift-online/ocm-sdk-go v0.1.224
	github.com/openshift/api v0.0.0-20211108165917-be1be0e89115
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	google.golang.org/api v0.44.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.22.3 // indirect
	k8s.io/apimachinery v0.22.3
)
