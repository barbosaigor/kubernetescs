package kubernetescs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// kubeGetResp is a response body of Kubernetes get pods API
type kubeGetResp struct {
	Items []struct {
		Metadata struct {
			Name string `json:"name"`
		} `json:"metadata"`
	} `json:"items"`
}

// shutdown terminates Kubernetes pods using K8s HTTP API
func shutdown(host, namespace, pod string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s/api/v1/namespaces/%s/pods/%s", host, namespace, pod), nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	return ErrInternalKube
}

// listPods list Kubernetes pods using K8s HTTP API
func listPods(host, namespace string) (pods []string, err error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/v1/namespaces/%s/pods", host, namespace))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var kresp kubeGetResp
	json.Unmarshal(body, &kresp)
	for _, item := range kresp.Items {
		pods = append(pods, item.Metadata.Name)
	}
	return
}
