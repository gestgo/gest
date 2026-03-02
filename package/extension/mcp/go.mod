module github.com/gestgo/gest/package/extension/mcp

go 1.23.0

toolchain go1.24.2

require (
	github.com/gestgo/gest/package/core/lifecycle v0.0.0-00010101000000-000000000000
	github.com/mark3labs/mcp-go v0.44.1
	go.uber.org/fx v1.23.0
)

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/invopop/jsonschema v0.13.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	go.uber.org/dig v1.18.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gestgo/gest/package/core/lifecycle => ../../core/lifecycle
