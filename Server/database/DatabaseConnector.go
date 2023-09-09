//██████╗░░█████╗░████████╗░█████╗░██████╗░░█████╗░░██████╗███████╗░░░░░░░░░█████╗░░█████╗░███╗░░██╗███╗░░██╗██████╗░░█████╗░████████╗░█████╗░██████╗░
//██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝░░░░░░░░██╔══██╗██╔══██╗████╗░██║████╗░██║██╔════╝██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗
//██║░░██║███████║░░░██║░░░███████║██████╦╝███████║╚█████╗░█████╗░░░░░░░░░░██║░░╚═╝██║░░██║██╔██╗██║██╔██╗██║█████╗░░██║░░╚═╝░░░██║░░░██║░░██║██████╔╝
//██║░░██║██╔══██║░░░██║░░░██╔══██║██╔══██╗██╔══██║░╚═══██╗██╔══╝░░░░░░░░░░██║░░██╗██║░░██║██║╚████║██║╚████║██╔══╝░░██║░░██╗░░░██║░░░██║░░██║██╔══██╗
//██████╔╝██║░░██║░░░██║░░░██║░░██║██████╦╝██║░░██║██████╔╝███████╗███████╗╚█████╔╝╚█████╔╝██║░╚███║██║░╚███║███████╗╚█████╔╝░░░██║░░░╚█████╔╝██║░░██║
//╚═════╝░╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░╚═╝╚═════╝░╚═╝░░╚═╝╚═════╝░╚══════╝╚══════╝░╚════╝░░╚════╝░╚═╝░░╚══╝╚═╝░░╚══╝╚══════╝░╚════╝░░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝

//This is a class for connecting to database, adding and changing it's data

package DatabaseConnector

import (
	"Server/database/DatabaseModels"
	"Server/database/DatabaseResponse"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type ConnectionSettings struct { // is not named as DatabaseConnector because one of imports already have such name
	Database string
	Password string
	Host     string
	Port     string
	User     string //database user
}

// NOTE use this before any communication with database
func (connection_settings *ConnectionSettings) is_set() bool { //helps to check if all connection data is filled
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

func (connection_settings *ConnectionSettings) does_user_exist(tag string) DatabaseResponse.DatabaseResponse {
	var response DatabaseResponse.DatabaseResponse
	user := new(DatabaseModels.User) //not actually used for now
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection_settings.User, connection_settings.Password, connection_settings.Host, connection_settings.Port, connection_settings.Database))

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM user WHERE tag = ?", tag)

	if err := row.Scan(&user.Id, &user.Tag, &user.Allow_hostory, &user.Allow_screenshot); err != nil {
		if err == sql.ErrNoRows {
			response.Message = "such user does not exist"
			response.Is_successful = false
		} else {
			response.Is_successful = true
			response.Message = "such user already exist"
			fmt.Println(tag, user.Tag)
		}

		return response
	}
	response.Is_successful = false
	response.Message = "unable to connect to the database. perhaps some settings are not filled"
	return response
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
	var response DatabaseResponse.DatabaseResponse
	//check if sattings are not empty
	if connection_settings.is_set() {
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
				response.Message = DBResponse.Message //"such user exists"
			}

		} else {
			fmt.Println(DBResponse.Message)
		}
	} else {
		fmt.Println("settings are not set")
	}
	response.Is_successful = false
	return response
}
