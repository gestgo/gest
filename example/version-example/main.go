package main

import (
	"fmt"

	"github.com/gestgo/gest/package/common/version"
)

// Inject via ldflags:
//
//	-X 'main.env=prod'
//	-X 'main.cluster=blue'
var (
	env     = "local"
	cluster = "none"
)

type AppMeta struct {
	Env     string `json:"env"`
	Cluster string `json:"cluster"`
	Service string `json:"service"`
}

func main() {
	fmt.Println("── 1. No extra ─────────────────────────────────────")
	version.New[any](nil).Print()

	fmt.Println("── 2. Typed struct extra ────────────────────────────")
	version.New(AppMeta{
		Env:     env,
		Cluster: cluster,
		Service: "billing",
	}).Print()

	fmt.Println("── 3. Simulate ldflags + struct ─────────────────────")
	version.AppName = "billing-svc"
	version.Version = "1.2.3"
	version.Branch = "main"
	version.Commit = "abc1234"
	version.BuildTime = "2026-04-21T10:00:00Z"

	version.New(AppMeta{
		Env:     "prod",
		Cluster: "green",
		Service: "billing",
	}).Print()
}
