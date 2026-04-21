run:
	go run .

build:
	go build -ldflags "\
		-X 'github.com/gestgo/gest/package/common/version.AppName=billing-svc' \
		-X 'github.com/gestgo/gest/package/common/version.Version=1.2.3' \
		-X 'github.com/gestgo/gest/package/common/version.Branch=main' \
		-X 'github.com/gestgo/gest/package/common/version.Commit=abc1234' \
		-X 'github.com/gestgo/gest/package/common/version.BuildTime=2026-04-21T10:00:00Z' \
		-X 'github.com/gestgo/gest/package/common/version.ExtraJSON={\"env\":\"prod\",\"cluster\":\"blue\"}'" \
		-o bin/app .

run-built: build
	./bin/app
