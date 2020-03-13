// TODO: Automate creations of pods before nodes
package kubernetescs

import (
	"testing"
)

func TestShutdown(t *testing.T) {
	podName := "sapi-pods" // Pod name for shutdown instance
	if err := Shutdown(host, namespace, podName); err != nil && err != ErrInternalKube {
		t.Errorf("shutdown: Fail to shutdown pod %s\n\t\terror: %s", podName, err.Error())
	}
}

func TestListPods(t *testing.T) {
	list, err := listPods(host, "default")
	t.Log(list)
	if err != nil {
		t.Errorf("listPods: Fail to list pods\n\t\terror: %s", err.Error())
	}
}
