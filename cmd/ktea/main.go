package main

import (
	"flag"

	"ktea/internal/ktea"
)

func main() {
	var strFlag string

	flag.StringVar(&strFlag, "c", "link", "env|link ('env' sets KUBECONFIG|'link' creates link)")
	flag.Parse()

	ktea.Ktea(strFlag, "/home/cmalfait/.kube")
}
