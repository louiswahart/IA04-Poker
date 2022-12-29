package main

import (
	"fmt"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/serveragent"
)

func main() {
	const urlserv = ":8080"
	serv := serveragent.NewServerAgent(urlserv, 1)
	go serv.Start()
	fmt.Scanln()
}
