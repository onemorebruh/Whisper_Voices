//███╗░░░███╗░█████╗░██╗███╗░░██╗
//████╗░████║██╔══██╗██║████╗░██║
//██╔████╔██║███████║██║██╔██╗██║
//██║╚██╔╝██║██╔══██║██║██║╚████║
//██║░╚═╝░██║██║░░██║██║██║░╚███║
//╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝╚═╝░░╚══╝

//This is the main file of server, so here everything loads and runs

package main

import (
	"Server/Controller"
	DatabaseConnector "Server/database"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func database_goroutine() {
	database_connector := DatabaseConnector.ConnectionSettings{
		Database: "whisper_voices",
		Password: "wh15p3r_v01c35", // NOTE password have to be read form configuration file
		Host:     "localhost",
		Port:     "3306",
		User:     "whisper_voices",
	}
	response := database_connector.Add_user("foo")
	fmt.Println(response.Message)
}

func main() {
	http.HandleFunc("/", Controller.Send_message)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
