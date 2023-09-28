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
	var stringified_response []byte

	//get body
	err := json.NewDecoder(request.Body).Decode(&body)
	var db_response DatabaseResponse.DatabaseResponse
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	switch body.Command { // run according the command
	case JsonBody.Create_user:
		{
			//communicate with database
			db_response = DatabaseConnector.DatabaseConnection.Create_user(body.User.Tag)
			//init response
			http_response.Command = JsonBody.Create_user
			http_response.Message = db_response.Message
			http_response.User = db_response.User
		}
	case JsonBody.Get_user:
		{
			//communicate with database
			db_response = DatabaseConnector.DatabaseConnection.Get_user(body.User.Tag)
			//init response
			http_response.Command = JsonBody.Get_user
			http_response.Message = db_response.Message
			http_response.User = db_response.User
		}
	case JsonBody.Create_invite:
		{
			//communicate with database
			db_response = DatabaseConnector.DatabaseConnection.Add_invite(body.User.Tag)
			fmt.Println("key is: ", db_response)
			//init response
			http_response.Command = JsonBody.Create_invite
			http_response.Message = db_response.Message
			http_response.Invite = db_response.Invite
		}
	}
	//send message
	stringified_response, err = json.Marshal(http_response)
	if err != nil {
		fmt.Println(err.Error())
		panic("error while encoding json")
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	writer.Write(stringified_response)
}
