package main

import (
	"github.com/NoUseFreak/cicd/pkg/server"
	"github.com/cicdpipelines/make-provider/internal/make"
)

func main() {
	server.Serve("make", make.Provider())
}
