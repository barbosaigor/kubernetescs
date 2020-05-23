// TODO: Instantiate k8s pods for testing
package kubernetescs

import (
	"testing"

	"github.com/barbosaigor/april/destroyer"
)

func TestShutdown(t *testing.T) {
	cs := ChaosServerKubernetes{"/home/igor/.kube/config", "lab", nil}
	cs.OnStart()
	err := cs.Shutdown("keychaindb-deployment-77dfdbf68f-vx76s")
	if err != nil {
		t.Errorf("Shutdown got an error: %v", err)
	}
}

func TestListInstances(t *testing.T) {
	cs := ChaosServerKubernetes{"/home/igor/.kube/config", "lab", nil}
	cs.OnStart()
	pods, err := cs.ListInstances(destroyer.Any)
	if err != nil {
		t.Errorf("ListInstances got an error: %v", err)
	}
	t.Log(pods)
}
