package main

import (
	"flag"
	"os"

	"ktea/internal/kfile"
)

func main() {
	strFlagPtr := flag.String("type", "link", "env|link ('env' sets KUBECONFIG|'link' creates link)")
	flag.Parse()

	kfile.Kfile(*strFlagPtr, os.Getenv("HOME") + "/.kube")
}
