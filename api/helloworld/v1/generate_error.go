package v1

//go:generate protoc -I . -I ../../../third_party --go_out=paths=source_relative:. --go-errors-kitx_out=paths=source_relative:. ./error_reason.proto
