package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/serveragent"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	log.Println("Seed:", t)
	var wg sync.WaitGroup
	serv := serveragent.NewServerAgent(1, 1, 2, &wg)
	go serv.Start()
	fmt.Scanln()
}
