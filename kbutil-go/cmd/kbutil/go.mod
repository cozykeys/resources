module main

go 1.15

replace github.com/cozykeys/resources/kbutil-go/pkg/kbutil => ../../pkg/kbutil

require (
	github.com/cozykeys/resources/kbutil-go/pkg/kbutil v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.25.0
)
