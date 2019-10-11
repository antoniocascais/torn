package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
	"torn/torn"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 3 {
		log.Panicf("You must provide 3 arguments to run the program: APIKEY, PROGRAM, PARAMS")
	}

	for true {
		cd, err := torn.GetChainCooldown(&http.Client{}, argsWithoutProg[0])
		if err != nil {
			log.Panic("error while getting chain cooldown")
		}
		fmt.Printf("Current cooldown: %d\n", cd)

		if cd > 0 && cd < 200 {
			cmd := exec.Command(argsWithoutProg[1], argsWithoutProg[2])
			err = cmd.Run()
			if err != nil {
				log.Panic(err)
			}
		}
		time.Sleep(30 * time.Second)
	}
}
