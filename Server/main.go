//███╗░░░███╗░█████╗░██╗███╗░░██╗
//████╗░████║██╔══██╗██║████╗░██║
//██╔████╔██║███████║██║██╔██╗██║
//██║╚██╔╝██║██╔══██║██║██║╚████║
//██║░╚═╝░██║██║░░██║██║██║░╚███║
//╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝╚═╝░░╚══╝

//This is the main file of server, so here everything loads and runs

package main

import (
	DatabaseConnector "Server/database"
	"fmt"
)

func main() {
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
