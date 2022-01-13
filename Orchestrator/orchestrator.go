package Orchestrator

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func ReadFile(searched string) string {

	var a string
	
	content, err := ioutil.ReadFile("Orchestrator/Config.txt")

	if err != nil {
		log.Fatal(err)
	}

	s := strings.Fields(string(content))

	for i, j := range s {
		if searched == j {
			a = strings.TrimSpace(s[i+1])
			fmt.Println(a)
			return a
		}
	}
	return "error, no result"
}

func Spawn(conn net.Conn, data []byte) {

	/*
		r = NewRequest(data, conn, SERVICE_ADDRESS, AUTH_ADDRESS)
		DoRequest()*/
}
