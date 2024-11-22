package main

import (
	"fmt"

	v1 "github.com/pecolynx/bazel-applications/proto/hello/v1"
)

func main() {
	request := v1.HelloRequest{}
	fmt.Println(request)
}
