package notaryserverresources

import (
	goharborv1alpha2 "github.com/goharbor/harbor-operator/api/v1alpha2"
)

type Manager struct {
	Clair *goharborv1alpha2.Clair
}