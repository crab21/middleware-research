module main

go 1.13

require crab.com/rbmq v0.0.1

require (
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/mitchellh/gox v1.0.1
)

replace crab.com/rbmq v0.0.1 => ../crab.com/rbmq
