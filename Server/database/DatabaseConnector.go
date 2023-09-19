//██████╗░░█████╗░████████╗░█████╗░██████╗░░█████╗░░██████╗███████╗░░░░░░░░░█████╗░░█████╗░███╗░░██╗███╗░░██╗██████╗░░█████╗░████████╗░█████╗░██████╗░
//██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝░░░░░░░░██╔══██╗██╔══██╗████╗░██║████╗░██║██╔════╝██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗
//██║░░██║███████║░░░██║░░░███████║██████╦╝███████║╚█████╗░█████╗░░░░░░░░░░██║░░╚═╝██║░░██║██╔██╗██║██╔██╗██║█████╗░░██║░░╚═╝░░░██║░░░██║░░██║██████╔╝
//██║░░██║██╔══██║░░░██║░░░██╔══██║██╔══██╗██╔══██║░╚═══██╗██╔══╝░░░░░░░░░░██║░░██╗██║░░██║██║╚████║██║╚████║██╔══╝░░██║░░██╗░░░██║░░░██║░░██║██╔══██╗
//██████╔╝██║░░██║░░░██║░░░██║░░██║██████╦╝██║░░██║██████╔╝███████╗███████╗╚█████╔╝╚█████╔╝██║░╚███║██║░╚███║███████╗╚█████╔╝░░░██║░░░╚█████╔╝██║░░██║
//╚═════╝░╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░╚═╝╚═════╝░╚═╝░░╚═╝╚═════╝░╚══════╝╚══════╝░╚════╝░░╚════╝░╚═╝░░╚══╝╚═╝░░╚══╝╚══════╝░╚════╝░░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝

//This is a class for connecting to database, adding and changing it's data

package DatabaseConnector

import (
	"Server/JsonBody"
	"Server/database/DatabaseModels"
	"Server/database/DatabaseResponse"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var HttpResponseChan chan JsonBody.JsonBody

type ConnectionSettings struct { // is not named as DatabaseConnector because one of imports already have such name
	Database string
	Password string
	Host     string
	Port     string
	User     string //database user
}

func Goroutine(database_channel chan JsonBody.JsonBody) {
	database_connector := ConnectionSettings{
		Database: "whisper_voices",
		Password: "wh15p3r_v01c35", // NOTE password have to be read form configuration file
		Host:     "localhost",
		Port:     "3306",
		User:     "whisper_voices",
	}
	for request := range database_channel {
		var db_response DatabaseResponse.DatabaseResponse
		switch request.Command {
		case JsonBody.Add_user:
			{
				//communicate with database
				db_response = database_connector.Add_user(request.User.Tag)
				//init response
				var http_response JsonBody.JsonBody
				http_response.Command = JsonBody.Add_user
				http_response.Message = db_response.Message
				//NOTE here could db_response.User be copied to the http_response.User but Add_user's returns it empty
				//send response
				HttpResponseChan <- http_response
			}
		case JsonBody.Get_user:
			{
				//TODO implement
			}
		}
	}
}

// NOTE use this before any communication with database
func (connection_settings *ConnectionSettings) Is_set() bool { //helps to check if all connection data is filled
	if connection_settings.Database != "" &&
		connection_settings.Password != "" &&
		connection_settings.Host != "" &&
		connection_settings.Port != "" &&
		connection_settings.User != "" {
		return true
	} else {
		return false
	}
}

// communication with user table
func (connection_settings *ConnectionSettings) does_user_exist(tag string) DatabaseResponse.DatabaseResponse {
	var db_response DatabaseResponse.DatabaseResponse
	user := new(DatabaseModels.User) //not actually used for now
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection_settings.User, connection_settings.Password, connection_settings.Host, connection_settings.Port, connection_settings.Database))

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM user WHERE tag = ?", tag)

	if err := row.Scan(&user.Id, &user.Tag, &user.Allow_hostory, &user.Allow_screenshot); err != nil {
		if err == sql.ErrNoRows {
			db_response.Message = "such user does not exist"
			db_response.Is_successful = false
		} else {
			db_response.Is_successful = true
			db_response.Message = "such user already exist"
			fmt.Println(tag, user.Tag)
		}

		return db_response
	}
	db_response.Is_successful = false
	db_response.Message = "unable to connect to the database. perhaps some settings are not filled"
	return db_response
}

func (connection_settings *ConnectionSettings) insert_user(user DatabaseModels.User) DatabaseResponse.DatabaseResponse {
	var result DatabaseResponse.DatabaseResponse
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection_settings.User, connection_settings.Password, connection_settings.Host, connection_settings.Port, connection_settings.Database))

	if err != nil {
		result.Is_successful = false
		result.Message = fmt.Sprintf("error while connecting to database: %s", err.Error())
		return result
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO user (id, tag) VALUES (?, ?)", user.Id, user.Tag)
	if err != nil {
		result.Is_successful = false
		result.Message = fmt.Sprintf("error while insert: %s", err.Error())
		return result
	}
	result.Message = "user successfully registred"
	result.Is_successful = true
	defer insert.Close()
	return result
}

func (connection_settings *ConnectionSettings) Add_user(tag string) DatabaseResponse.DatabaseResponse {
	var db_response DatabaseResponse.DatabaseResponse
	//check if sattings are not empty
	if connection_settings.Is_set() {
		//get user by tag to check if tag is available
		var DBResponse DatabaseResponse.DatabaseResponse
		DBResponse = connection_settings.does_user_exist(tag)
		if DBResponse.Is_successful == false { //true means that user exists
			if DBResponse.User.Id == uuid.Nil {
				//init User object
				var user DatabaseModels.User
				user.Id = uuid.New()
				user.Tag = tag
				// insert user into database
				insert_response := connection_settings.insert_user(user)
				//return result
				return insert_response
			} else {
				fmt.Println(DBResponse.User.Tag, DBResponse.User.Id)
				db_response.Message = DBResponse.Message //"such user exists"
			}

		} else {
			fmt.Println(DBResponse.Message)
		}
	} else {
		fmt.Println("settings are not set")
	}
	db_response.Is_successful = false
	return db_response
}

//communication with invite table

func (connection_settings *ConnectionSettings) insert_invite(user uuid.UUID) DatabaseResponse.DatabaseResponse {
	var word = "TODO" //TODO implement function i did in python for generating passwords
	var result DatabaseResponse.DatabaseResponse
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection_settings.User, connection_settings.Password, connection_settings.Host, connection_settings.Port, connection_settings.Database))

	if err != nil {
		result.Is_successful = false
		result.Message = fmt.Sprintf("error while connecting to database: %s", err.Error())
		return result
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO invite (id, user, value) VALUES (?, ?, ?)", uuid.New(), user, word)
	if err != nil {
		result.Is_successful = false
		result.Message = fmt.Sprintf("error while insert: %s", err.Error())
		return result
	}
	result.Message = "invite successfully created"
	result.Is_successful = true
	result.Invite = word
	defer insert.Close()
	return result

}

func (connection_settings *ConnectionSettings) Add_invite(tag string) DatabaseResponse.DatabaseResponse {
	var db_response DatabaseResponse.DatabaseResponse
	var user DatabaseModels.User

	if connection_settings.Is_set() {

		db_response = connection_settings.does_user_exist(tag)
		if db_response.Is_successful { // if user exists

			user = db_response.User
			db_response = connection_settings.insert_invite(user.Id)

		}
	} else {
		db_response.Is_successful = false
		db_response.Message = "settings are not set"
	}
	return db_response
}
