package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/antoniocascais/torn/torn"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 3 {
		log.Panicf("You must provide 3 arguments to run the program: APIKEY, PROGRAM, PARAMS")
	}

	for true {
		to, err := torn.GetChainTimeout(&http.Client{}, argsWithoutProg[0])
		if err != nil {
			log.Panicf("error while getting chain cooldown: %s", err.Error())
		}
		fmt.Printf("Current timeout: %d\n", to)

		if to > 0 && to < 200 {
			cmd := exec.Command(argsWithoutProg[1], argsWithoutProg[2])
			err = cmd.Run()
			if err != nil {
				log.Panic(err)
			}
		}
		time.Sleep(30 * time.Second)
	}
}
