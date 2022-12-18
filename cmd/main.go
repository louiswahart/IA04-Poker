package main

import (
	"fmt"
	"sync"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/serveragent"
)

func main() {
	const urlserv = ":8080"
	var wg sync.WaitGroup
	serv := serveragent.NewServerAgent(urlserv, 1, 1, 2, &wg)
	go serv.Start()
	fmt.Scanln()
}
