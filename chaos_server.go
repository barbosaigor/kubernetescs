package kubernetescs

import (
	"context"

	"github.com/barbosaigor/april/destroyer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// ChaosServerKubernetes implements chaos server interface from april/destroyer
type ChaosServerKubernetes struct {
	// ConfPath is the Kubernetes configuration file
	ConfPath string
	// Namespace value for Kubernetes operations
	Namespace string
	// client is a kubernetes client
	client *kubernetes.Clientset
}

// New creates a instance of ChaosServerKubernetes
func New(conf, namespace string) *ChaosServerKubernetes {
	return &ChaosServerKubernetes{ConfPath: conf, Namespace: namespace, client: nil}
}

// Shutdown turns off a pod
func (k ChaosServerKubernetes) Shutdown(pod string) error {
	deletePolicy := metav1.DeletePropagationForeground
	return k.client.AppsV1().ReplicaSets(k.Namespace).Delete(context.TODO(), pod, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

// ListInstances lists all instances with corresponding status
func (k ChaosServerKubernetes) ListInstances(status destroyer.Status) ([]destroyer.Instance, error) {
	pods, err := k.listPods()
	if err != nil {
		return nil, err
	}
	inst := make([]destroyer.Instance, len(pods))
	for i, pod := range pods {
		inst[i] = destroyer.Instance{Name: pod, Sts: destroyer.Up}
	}
	return inst, nil
}

// OnStart is a life cycle routine when stating the chaos server.
func (k *ChaosServerKubernetes) OnStart() {
	// Make connection with k8 api
	config, err := clientcmd.BuildConfigFromFlags("", k.ConfPath)
	if err != nil {
		panic(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	k.client = clientSet
}

func (k ChaosServerKubernetes) listPods() ([]string, error) {
	pods, err := k.client.CoreV1().Pods(k.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	p := make([]string, len(pods.Items))
	for i, pod := range pods.Items {
		p[i] = pod.GetName()
	}
	return p, nil
}
