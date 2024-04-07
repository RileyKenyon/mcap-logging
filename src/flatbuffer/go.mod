module fbs-encode

go 1.22.1

require (
	github.com/foxglove/mcap/go/mcap v1.3.0
	github.com/google/flatbuffers v24.3.25+incompatible
)

require (
	MyGame/Sample v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/klauspost/compress v1.15.12 // indirect
	github.com/pierrec/lz4/v4 v4.1.12 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace MyGame/Sample => ./MyGame/Sample/
