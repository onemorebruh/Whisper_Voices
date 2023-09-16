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

func main() {
	http.HandleFunc("/", Controller.Get_message)

	go DatabaseConnector.Goroutine(Controller.Database_channel)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
