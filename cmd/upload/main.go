package main

import "gitlab.com/samkomarov/upload-svc.git/internal"

func main() {
	cfg := internal.ReadConfigFromEnv()
	internal.InitializeAndStart(cfg)
}
