package main

import "gitlab.com/samkomarov/uploads-svc.git/internal"

func main() {
	cfg := internal.ReadConfigFromEnv()
	internal.InitializeAndStart(cfg)
}
