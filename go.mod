module github.com/mokhae/opcua

go 1.24

require (
	github.com/google/uuid v1.6.0
	github.com/pascaldekloe/goe v0.1.1
	github.com/pkg/errors v0.9.1
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394
	golang.org/x/term v0.30.0
)

require golang.org/x/sys v0.31.0 // indirect

retract (
	v0.2.5 // https://github.com/mokhae/opcua/issues/538
	v0.2.4 // https://github.com/mokhae/opcua/issues/538
)
