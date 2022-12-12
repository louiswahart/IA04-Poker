package main

import (
	"fmt"
	"sync"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/serveragent"
)

func main() {
	var wg sync.WaitGroup
	serv := serveragent.NewServerAgent(1, 1, 1, &wg)
	go serv.Start()
	fmt.Scanln()
}
