//░█████╗░░█████╗░███╗░░██╗████████╗██████╗░░█████╗░██╗░░░░░██╗░░░░░███████╗██████╗░
//██╔══██╗██╔══██╗████╗░██║╚══██╔══╝██╔══██╗██╔══██╗██║░░░░░██║░░░░░██╔════╝██╔══██╗
//██║░░╚═╝██║░░██║██╔██╗██║░░░██║░░░██████╔╝██║░░██║██║░░░░░██║░░░░░█████╗░░██████╔╝
//██║░░██╗██║░░██║██║╚████║░░░██║░░░██╔══██╗██║░░██║██║░░░░░██║░░░░░██╔══╝░░██╔══██╗
//╚█████╔╝╚█████╔╝██║░╚███║░░░██║░░░██║░░██║╚█████╔╝███████╗███████╗███████╗██║░░██║
//░╚════╝░░╚════╝░╚═╝░░╚══╝░░░╚═╝░░░╚═╝░░╚═╝░╚════╝░╚══════╝╚══════╝╚══════╝╚═╝░░╚═╝

// This is a compilation of Server logic it uses while working with http requests

package Controller

import (
	"Server/JsonBody"
	DatabaseConnector "Server/database"
	"Server/database/DatabaseResponse"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//var DatabaseConnection DatabaseConnector.ConnectionSettings

func init() {
	fmt.Println("initialized")
	DatabaseConnector.DatabaseConnection = DatabaseConnector.ConnectionSettings{
		Database: "whisper_voices",
		Password: "wh15p3r_v01c35", // NOTE password have to be read form configuration file
		Host:     "localhost",
		Port:     "3306",
		User:     "whisper_voices",
	}
	fmt.Println("database connection:", DatabaseConnector.DatabaseConnection.Is_set())
}

func Get_message(writer http.ResponseWriter, request *http.Request) {
	var body JsonBody.JsonBody
	var http_response JsonBody.JsonBody

	//get body
	err := json.NewDecoder(request.Body).Decode(&body)
	var db_response DatabaseResponse.DatabaseResponse
	switch body.Command {
	case JsonBody.Add_user:
		{
			//communicate with database
			fmt.Println(DatabaseConnector.DatabaseConnection.Database)
			db_response = DatabaseConnector.DatabaseConnection.Add_user(body.User.Tag)
			//init response
			http_response.Command = JsonBody.Add_user
			http_response.Message = db_response.Message
			//NOTE here could db_response.User be copied to the http_response.User but Add_user does not return it
			//send response
		}
	case JsonBody.Get_user:
		{
			//TODO implement
		}
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	io.WriteString(writer, http_response.Message)
}
