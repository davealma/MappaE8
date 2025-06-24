package main

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var encodedStrings []string

func DecodeString(encodedString string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}

	return string(decodedBytes)
}

func HoldDoor(value string) string {
	if value != "" {
		encodedStrings = append(encodedStrings, value)
	}
	cookie, resp := HitDors(value)
	fmt.Println("Resp: ", resp)
	time.Sleep(500 * time.Millisecond)
	if strings.Contains(resp, "revelio") {
		fmt.Println("DETENER")
		return cookie
	}else {
		return HoldDoor(cookie)
	}
}

func main() {
	godotenv.Load()

	HoldDoor("")
	var hiddenMessage []string
	for _, value  := range encodedStrings {
		hiddenMessage = append(hiddenMessage, DecodeString(value))	 	
	}
	fmt.Println("Mensage: ", strings.Join(hiddenMessage, " "))
	PostSolution(strings.Join(hiddenMessage, " "))
}