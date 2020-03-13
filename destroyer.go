package kubernetescs

type KubernetesDestryr struct {
	// Host of Kubernetes Cluster
	Host string
	// Namespace value for Kubernetes operations
	Namespace string
}

func (d KubernetesDestryr) Destroy(nodes []string) error {
	if nodes == nil {
		return ErrEmptyService
	}
	for _, n := range nodes {
		if err := Shutdown(d.Host, d.Namespace, n); err != nil {
			return err
		}
	}
	return nil
}
