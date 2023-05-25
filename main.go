package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/optclblast/filetagger/api"
	"github.com/optclblast/filetagger/auth"
	"github.com/optclblast/filetagger/logger"
)

const (
	SERVER_ADDRESS = "localhost"
)

func main() {
	server := api.NewServer()
	logger.Logger(logger.LogEntry{
		DateTime: time.Now(),
		Level:    logger.INFO,
		Location: "MAIN/main() SCOPE",
		Content:  "Server started",
	})

	err := auth.GenerateKeyPair(askAboutRegenetationRSAKeys())
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "MAIN/main() SCOPE",
			Content:  fmt.Sprintf("An error occured at keys generation step:%s", err),
		})
	}

	err = http.ListenAndServe(SERVER_ADDRESS+portDeclaration(), server)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.FATAL,
			Location: "MAIN/main() SCOPE",
			Content:  fmt.Sprintf("An error occured at server runtime:%s", err),
		})
	}
}

func portDeclaration() string {
	fmt.Println("Please, enter a port the server will be listening to")
	reader := bufio.NewReader(os.Stdin)
	port, _ := reader.ReadString('\n')
	port = port[:len(port)-2]
	logger.Logger(logger.LogEntry{
		DateTime: time.Now(),
		Level:    logger.INFO,
		Location: "MAIN/portDeclaration() SCOPE",
		Content:  fmt.Sprintf("Listening to :%s", port),
	})
	return fmt.Sprintf(":%s", port)
}

func askAboutRegenetationRSAKeys() bool {
	fmt.Println("Do you want to re-generate RSA keys? [Y/N]")
	reader := bufio.NewReader(os.Stdin)
	mode, _ := reader.ReadString('\n')
	mode = mode[:len(mode)-2]
	logger.Logger(logger.LogEntry{
		DateTime: time.Now(),
		Level:    logger.INFO,
		Location: "MAIN/portDeclaration() SCOPE",
		Content:  fmt.Sprintf("Re-generation mode: %s", mode),
	})
	if mode == "Y" || mode == "y" {
		return true
	} else if mode == "N" || mode == "n" {
		return false
	}
	logger.Logger(logger.LogEntry{
		DateTime: time.Now(),
		Level:    logger.INFO,
		Location: "MAIN/portDeclaration() SCOPE",
		Content:  fmt.Sprintf("Got unexpected input:%s || Mod set as FALSE", mode),
	})
	return false
}
