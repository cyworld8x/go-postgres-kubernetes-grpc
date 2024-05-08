//go:build tools
// +build tools

//Use this to keep package tools in the final binary

package tools

import (
	_ "github.com/golang/mock/mockgen/model"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
)
