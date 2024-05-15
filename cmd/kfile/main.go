package main

import (
	"flag"

	"ktea/internal/kfile"
)

func main() {
	var strFlag string

	flag.StringVar(&strFlag, "c", "link", "env|link ('env' sets KUBECONFIG|'link' creates link)")
	flag.Parse()

	kfile.Kfile(strFlag, "/home/cmalfait/.kube")
}
