package main

import (
	"flag"
	"os"

	"ktea/internal/ktea"
)

func main() {
	strFlagPtr := flag.String("type", "link", "env|link ('env' sets KUBECONFIG|'link' creates link)")
	flag.Parse()

	ktea.Ktea(*strFlagPtr, os.Getenv("HOME")+"/.kube")
}
