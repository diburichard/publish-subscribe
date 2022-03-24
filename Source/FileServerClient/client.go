package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func validateArguments(action string, channel string, path string) (err error) {

	if action != "send" && action != "receive" {
		// las accion no reconocida
		return errors.New("-action Error: Only process 'send' or 'receive' commands")
	}

	if action == "send" && (channel == "" || path == "") {
		// la anccion de envio requiere los parametros de channel y path
		return errors.New("send Error: channel and path is not empty")
	}

	if action == "receive" && channel == "" {
		return errors.New("receive Error: channel is not empty")
	}

	if action == "send" && channel != "" && path != "" {
		// Valida existncia de path
		if _, err := os.Stat(path); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	action := flag.String("action", "", "Accion a realizar con el cliente [send] [receive]")
	channel := flag.String("channel", "", "Canal para envio de archivos o escucha de archivos")
	path := flag.String("path", "", "Url del archivo a enviar")
	flag.Parse()
	err := validateArguments(*action, *channel, *path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*action)
	fmt.Println(*channel)
	fmt.Println(*path)
}
