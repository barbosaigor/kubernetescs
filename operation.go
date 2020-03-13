package kubernetescs

// Shutdown a especific pod
func Shutdown(host, namespace, pod string) error {
	if pod == "" {
		return ErrEmptyService
	}
	return shutdown(host, namespace, pod)
}
