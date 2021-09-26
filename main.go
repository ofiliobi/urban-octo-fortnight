package main

import (
	_"encoding/json"
	"github.com/ofiliobi/urban-octo-fortnight/infrastructure"
)

func main() {
	infrastructure.NewHTTPServer().Start()
}
