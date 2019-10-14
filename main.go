package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/antoniocascais/torn/torn"
)

func main() {
	var err error

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 3 {
		log.Panicf("You must provide, at least, 3 arguments to run the program: APIKEY, PROGRAM, PARAMS")
	}

	timeout := 200
	if len(argsWithoutProg) == 4 {
		timeout, err = strconv.Atoi(argsWithoutProg[3])
		if err != nil {
			log.Panicf("Error while parsing %s to a number", argsWithoutProg[3])
		}
	}

	for true {
		to, err := torn.GetChainTimeout(&http.Client{}, argsWithoutProg[0])
		if err != nil {
			log.Panicf("error while getting chain cooldown: %s", err.Error())
		}
		fmt.Printf("Current timeout: %d\n", to)

		if to > 0 && to < timeout {
			cmd := exec.Command(argsWithoutProg[1], argsWithoutProg[2])
			err = cmd.Run()
			if err != nil {
				log.Panic(err)
			}
		}
		time.Sleep(30 * time.Second)
	}
}
