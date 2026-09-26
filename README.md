# Testify template

Provisioned from [`Qode-Platform/fleet-template-v1`](https://github.com/Qode-Platform/fleet-template-v1) - the fleet
lifecycle contract with a Testify starter on top.

## Verified

Built and tested locally on Go 1.23.4 (toolchain auto-upgraded to 1.25):
`go build ./...` and `go test ./...` both pass.

## Fleet lifecycle

| step | command |
|---|---|
| install | `go mod download` |
| build | `go vet ./... && go build ./...` |
| start | `(none - not a service)` |

## Notes

- NOT A SERVICE: a testing library. START_CMD is empty; the image is a test runner (ENTRYPOINT go test ./... -v).
- Run the suite with: go test ./...
