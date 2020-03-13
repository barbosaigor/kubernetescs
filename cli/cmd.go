// package cli contain console line interface that hosts an API
// which shutdown pods (at this moment).
//
// TODO: expand to services and deploys 
package cli

import (
	"github.com/barbosaigor/april/destroyer"
	ksd "github.com/barbosaigor/kubernetescs"
	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

var username string
var password string
var namespace string
var host string
var port int

func init() {
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	rootCmd.Flags().StringVarP(&password, "password", "s", "", "Password")
	rootCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:8001", "Kubernetes Host")
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Kubernetes Namespace")
	rootCmd.Flags().IntVarP(&port, "port", "p", 7071, "Server port number")
	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("password")
}

var rootCmd = &cobra.Command{
	Use:   "kubernetescs",
	Short: "Kubernetes chaos server terminates instances via API.",
	Long:  "Kubernetes chaos server terminates instances via API.",
	Run: func(cmd *cobra.Command, args []string) {
		serv := destroyer.New(port, ksd.KubernetesDestryr{host, namespace})
		serv.Cred.Register(username, password)
		serv.Serve()
	},
	Version: VERSION,
}

func Execute() error {
	return rootCmd.Execute()
}
