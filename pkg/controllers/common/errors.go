package common

import (
	serrors "github.com/goharbor/harbor-operator/pkg/controllers/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (c *Controller) HandleError(err error) (ctrl.Result, error) {
	if err == nil {
		return ctrl.Result{}, nil
	}

	if err, ok := err.(*serrors.RetryLaterError); ok {
		return err.Result(), nil
	}

	if _, ok := err.(*serrors.UnrecoverrableError); ok {
		// TODO Mark error in status
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, err
}
