test:
	@ export GO_ENV=TEST
	@ export GOCACHE=off
	@ go test ./nats -v
	@ go test ./utils -v
	@ go test ./cmd -v
