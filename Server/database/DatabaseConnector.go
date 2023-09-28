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
	PasswordGenerator "Server/database/PasswordGeneration"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var DatabaseConnection ConnectionSettings

func init() {
	PasswordGenerator.Password_generator = PasswordGenerator.Init()
}

type ConnectionSettings struct { // is not named as DatabaseConnector because one of imports already have such name
	Database string
	Password string
	Host     string
	Port     string
	User     string //database user
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
	var user DatabaseModels.User
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
			db_response.User = user
			db_response.Is_successful = true
			db_response.Message = "such user already exists"
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

func (connection_settings *ConnectionSettings) Create_user(tag string) DatabaseResponse.DatabaseResponse {
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
		fmt.Println(connection_settings.Is_set())
		fmt.Println("settings are not set")
	}
	db_response.Is_successful = false
	return db_response
}

func (connection_settings *ConnectionSettings) Get_user(tag string) DatabaseResponse.DatabaseResponse {
	var user DatabaseModels.User
	var db_response DatabaseResponse.DatabaseResponse

	if connection_settings.Is_set() {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection_settings.User, connection_settings.Password, connection_settings.Host, connection_settings.Port, connection_settings.Database))

		if err != nil {
			db_response.Is_successful = false
			db_response.Message = fmt.Sprintf("error while connecting to database: %s", err.Error())
			return db_response
		}
		defer db.Close()

		row := db.QueryRow("SELECT * FROM user WHERE tag = ?", tag)

		if err := row.Scan(&user.Id, &user.Tag, &user.Allow_hostory, &user.Allow_screenshot); err != nil {
			if err == sql.ErrNoRows {
				db_response.Message = "such user does not exist"
				db_response.Is_successful = false
			} else {
				db_response.Is_successful = true
				db_response.Message = "such user already exists"
				fmt.Println(tag, user.Tag)
			}

			return db_response
		}
	}

	db_response.Is_successful = false
	db_response.Message = "unable to connect to the database. perhaps some settings are not filled"

	return db_response
}

//communication with invite table

func (connection_settings *ConnectionSettings) insert_invite(user uuid.UUID) DatabaseResponse.DatabaseResponse {
	var word = PasswordGenerator.Password_generator.Create_password()
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
