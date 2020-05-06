package kubernetescs

import "errors"

var (
	// ErrAPI indicates an error on API
	ErrAPI = errors.New("Internal API server error")
	// ErrEmptyService indicates empty service
	ErrEmptyService = errors.New("Empty Service")
	// ErrInternalKube indicates an API kubernetes error
	ErrInternalKube = errors.New("Internal Kubernetes error")
)
