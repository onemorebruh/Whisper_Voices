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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Database_channel chan JsonBody.JsonBody

func Get_message(writer http.ResponseWriter, request *http.Request) {
	var body JsonBody.JsonBody

	//get body
	err := json.NewDecoder(request.Body).Decode(&body)
	Database_channel <- body

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	//TODO try to get data from channel
	fmt.Println("works!")
	io.WriteString(writer, "UwU\n")
}
