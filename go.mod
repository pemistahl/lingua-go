module github.com/pemistahl/lingua-go

go 1.18

require (
	github.com/pemistahl/lingua-go/serialization v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.1
	golang.org/x/exp v0.0.0-20221106115401-f9659909a136
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/pemistahl/lingua-go/serialization => ./serialization
