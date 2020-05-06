package main

import (
	"log"

	"github.com/barbosaigor/april/destroyer/cli"
	"github.com/barbosaigor/kubernetescs"
)

func main() {
	// Change cli long description
	cli.RootCmd.Long = "Kubernetescs (Chaos server) shutdowns kubernetes pods via an API."
	cli.Cs = kubernetescs.ChaosServerKubernetes{Host: "127.0.0.1:8001", Namespace: "lab"}
	if err := cli.Execute(); err != nil {
		log.Printf("Error: %v\n", err)
	}
}
