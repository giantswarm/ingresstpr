package ingresstpr

import (
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/api/v1"
)

// CustomObject represents the ingress TPR's custom object. It holds the
// specifications of the resource the ingress operator is interested in.
type CustomObject struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty"`

	Spec Spec `json:"spec" yaml:"spec"`
}
