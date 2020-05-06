package kubernetescs

import "github.com/barbosaigor/april/destroyer"

// ChaosServerKubernetes implements chaos server interface from april/destroyer
type ChaosServerKubernetes struct {
	// Host of Kubernetes Cluster
	Host string
	// Namespace value for Kubernetes operations
	Namespace string
}

// Shutdown turns off a pod
func (k ChaosServerKubernetes) Shutdown(svc string) error {
	return Shutdown(k.Host, k.Namespace, svc)
}

// ListInstances lists all instances with corresponding status
func (k ChaosServerKubernetes) ListInstances(status destroyer.Status) ([]destroyer.Instance, error) {
	pods, err := listPods(k.Host, k.Namespace)
	if err != nil {
		return nil, err
	}
	inst := make([]destroyer.Instance, len(pods))
	for i, pod := range pods {
		inst[i] = destroyer.Instance{Name: pod, Sts: destroyer.Up}
	}
	return inst, nil
}

// OnStart is a life cycle routine when stating the chaos server
func (k ChaosServerKubernetes) OnStart() {
}
