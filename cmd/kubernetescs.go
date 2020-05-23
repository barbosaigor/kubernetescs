package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/barbosaigor/april/destroyer/cli"
	"github.com/barbosaigor/kubernetescs"
)

var namespace string
var confPath string

func init() {
	cli.RootCmd.Flags().StringVarP(&namespace, "namespace", "", "default", "Kubernetes namespace")
	cli.RootCmd.Flags().StringVarP(&confPath, "configuration", "", filepath.Join(homeDir(), ".kube", "config"), "Kubernetes cluster configuration file")
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func main() {
	// CLI long description
	cli.RootCmd.Long = "Kubernetescs (Chaos Server) shutdowns kubernetes pods via an API."
	cli.Cs = kubernetescs.New(confPath, namespace)
	if err := cli.Execute(); err != nil {
		log.Printf("Error: %v\n", err)
	}
}
