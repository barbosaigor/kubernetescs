package main

import (
	"log"

	"github.com/barbosaigor/april/destroyer/cli"
	"github.com/barbosaigor/kubernetescs"
)

var (
	// Kubernetes host endpoint
	endpoint string
	// Kubernetes namespace
	namespace string
)

func init() {
	cli.RootCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "127.0.0.1:8001", "Enpoint of Kubernetes API")
	cli.RootCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Namespace of Kubernetes for CS to operate")
}

func main() {
	// Change cli long description
	cli.RootCmd.Long = "kubernetescs (Chaos server) shutdowns kubernetes pods via an API."
	cli.Cs = kubernetescs.ChaosServerKubernetes{Host: endpoint, Namespace: namespace}
	if err := cli.Execute(); err != nil {
		log.Printf("Error: %v\n", err)
	}
}
