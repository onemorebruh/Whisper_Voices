//██████╗░░█████╗░████████╗░█████╗░██████╗░░█████╗░░██████╗███████╗░░░░░░░░░█████╗░░█████╗░███╗░░██╗███╗░░██╗██████╗░░█████╗░████████╗░█████╗░██████╗░
//██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝░░░░░░░░██╔══██╗██╔══██╗████╗░██║████╗░██║██╔════╝██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗
//██║░░██║███████║░░░██║░░░███████║██████╦╝███████║╚█████╗░█████╗░░░░░░░░░░██║░░╚═╝██║░░██║██╔██╗██║██╔██╗██║█████╗░░██║░░╚═╝░░░██║░░░██║░░██║██████╔╝
//██║░░██║██╔══██║░░░██║░░░██╔══██║██╔══██╗██╔══██║░╚═══██╗██╔══╝░░░░░░░░░░██║░░██╗██║░░██║██║╚████║██║╚████║██╔══╝░░██║░░██╗░░░██║░░░██║░░██║██╔══██╗
//██████╔╝██║░░██║░░░██║░░░██║░░██║██████╦╝██║░░██║██████╔╝███████╗███████╗╚█████╔╝╚█████╔╝██║░╚███║██║░╚███║███████╗╚█████╔╝░░░██║░░░╚█████╔╝██║░░██║
//╚═════╝░╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░╚═╝╚═════╝░╚═╝░░╚═╝╚═════╝░╚══════╝╚══════╝░╚════╝░░╚════╝░╚═╝░░╚══╝╚═╝░░╚══╝╚══════╝░╚════╝░░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝

//This is a class for connecting to database, adding and changing it's data

package main

import (
	database_response "Server/database"
	"database/sql"
	"fmt"

	//user "Server/database/models"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectionSettings struct {
	database string
	password string
	host     string
	port     string
	user     string //database user
}

func (connection_settings *ConnectionSettings) is_set() bool { //helps to check if all connection data is filled
	if connection_settings.database != "" &&
		connection_settings.password != "" &&
		connection_settings.host != "" &&
		connection_settings.port != "" &&
		connection_settings.user != "" {
		return true
	} else {
		return false
	}
}

func (connection_settings *ConnectionSettings) does_user_exist(tag string) DatabaseResponse {
	response := database_response.DatabaseResponse
	user := User //not actually used for now
	if connection_settings.is_set() {
		db, error := sql.Open("mysql", fmt.Sprintf("?:?@tcp(?:?)/?", connection_settings.user, connection_settings.password, connection_settings.host, connection_settings.port, connection_settings.database)) //NOTE i am not sure will it work or not. if it doesn't just change ? to %U in this line

		if error != nil {
			panic(error.Error())
		}
		defer db.Close()

		row := db.QueryRow("SELECT * FROM user WHERE tag = ?", tag)

		if error := row.Scan(&user.id, &user.tag, &user.allow_hostory, &user.allow_screenshot); error != nil {
			if error == sql.ErrNoRows {
				response.message = "such user does not exist"
				response.is_successful = false
			} else {
				response.is_successful = true
				response.message = "such user already exist"
			}
		}

		return response
	}
	response.is_successful = false
	response.message = "unable to connect to the database. perhaps some settings are not filled"
	return response
}

func (connection_settings *ConnectionSettings) addUser(tag string) database_response.DatabaseResponse {
	response := database_response.DatabaseResponse
	//check if sattings are not empty
	if connection_settings.is_set() {
		//get user by tag to check if tag is available

		//init User object

		// insert user into database

		//return result
		return
	} else {
		response.message = "database is not connected"
		response.is_successful = false
		return response
	}
}

/*

func main() {
	db, err := sql.Open("mysql", "whisper_voices:wh15p3r_v01c35@tcp(localhost:3306)/whisper_voices")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	id := uuid.New()

	insert, err := db.Query("INSERT INTO user (id, tag) VALUES (?, ?)", id, "cutie")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.id, &user.tag, &user.allow_hostory, &user.allow_screenshot)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user)
	}

}

*/
