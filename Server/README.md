# Server

This is a direcotry with a `Whisper_Voices` server.

## What server does

* connect to database
* save new user's tag and settings in database
* get info about users for adding user to the client's contacts
* update user's tag or settings
* send messages from users to users

## What server does not

* save messages
* save ip address or any other info

## Database

Here you can read about each scheme Server uses

### User

| Field             | Type          | Default   | Extra         | Description                                       |
|-------------------|---------------|-----------|---------------|---------------------------------------------------|
| `id`              | uuid          |           | primary key   | identificator                                     |
| `tag`             | varchar(50)   |           | unique        | used to add user to contacts                      |
| `allow_history`   | bool          | false     |               | saves this user's messages on contact's devices   |
| `allow_screenshot`| bool          | false     |               | allows to do screenshots on contact's devices     |

## Classes

Here you can read about each class in this programm

### DatabaseModels

#### User

User is a class used while communicating with database

| Name              | Field/Method  | Type          | Description       |
|-------------------|---------------|---------------|-------------------|
| `Id`              | Field         | uuid.UUID     | identificator     |
| `Tag`             | Field         | string        | unique user tag   |
| `Allow_history`   | Field         | bool          | check User scheme |
| `Allow_screenshot`| Field         | bool          | check User scheme |

#### Invite

Invite is a class used while communicating with database.
Administrator need it to blocking users and users need it to invite new users to invite new users

| Name              | Field/Method  | Type          | Description                           |
|-------------------|---------------|---------------|---------------------------------------|
| `Id`              | Field         | uuid.UUID     | identificator                         |
| `User`            | Field         | uuid.UUID     | user identificator                    |
| `Value`           | Field         | string        | string user share to invite new one   |

### DatabaseResponse

DatabaseResponse is used in DatabaseConnector to get data from methods

| Name              | Field/Method  | Type          | Description                   |
|-------------------|---------------|---------------|-------------------------------|
| `Message`         | Field         | string        | status message                |
| `Is_successful`   | Field         | bool          | was communication successful  |
| `User`            | Field         | User.User     | user gotten by tag            |

### DatabaseConnector - ConnectionSettings

ConnectionSettings is used for communication with mysql database
It is not named as DatabaseConnector because there is such class in one of dependencies

| Name              | Field/Method  | Type              | Description                           |
|-------------------|---------------|-------------------|---------------------------------------|
| `Database`        | Field         | string            | database name                         |
| `Password`        | Field         | string            | database user password                |
| `Host`            | Field         | string            | database host                         |
| `Port`            | Field         | string            | database port                         |
| `User`            | Field         | string            | database user                         |
| `is_set`          | Method        | bool              | checks if all fields are set          |
| `does_user_exists`| Method        | DatabaseResponse  | checks if user with given tag exists  |
| `insert_user`     | Method        | DatabaseResponse  | inserts user into database            |
| `Create_user`     | Method        | DatabaseResponse  | inserts user into database            |
| `Get_user`        | Method        | DatabaseResponse  | returns user from database            |
| `insert_invite`   | Method        | DatabaseResponse  | inserts invite into database          |
| `Create_invite`   | Method        | DatabaseResponse  | inserts user into database            |

> Attention! It is not secure to use `insert_user` and `insert_invite`. you should rather use `Add_user` and `Create_invite`

usage example:
```go
database_connector := DatabaseConnector.ConnectionSettings{
    Database: "dbname",
    Password: "dbpass",
    Host: "localhost",
    Port: "3306",
    User: "root",
}
response := database_connector.Add_user("foo")
fmt.Println(response.Message)
```

expexted output:
```
user successfully registred
```

### JsonBody

JsonBody is a class for parsing and sending json via

| Name      | Type                  | Tag       | Description                   |
|-----------|-----------------------|-----------|-------------------------------|
| `User`    | DatabaseModels.User   | user      | database user                 |
| `Address` | string                | address   | address of request sender     |
| `Message` | string                | message   | status message                |
| `Command` | DBcommand             | command   | command to do on server       |

#### DBcommand

DBcommand is a Enum. used for getting data from database and creating new records in it.

| Possible values   |
|-------------------|
| `Create_user`     |
| `Get_user`        |
| `Create_invite`   |

### Controller

Controller is a compilation of buissness logic functions

| Name          | Description       |
|---------------|-------------------|
| `Get_message` | call to the Server|

usage example:
```go
func main() {
	http.HandleFunc("/", Controller.Get_message)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
```

example of request:
```js
{
    "user": {
        "tag": "username"
    },
    "command": 0
}
```

example of response:
```
user successfully registred
```

### PasswordGenerator

This is a class used for generating unique like combinations of words. later used and called as passwords

| Name              | Field/Method  | Type              | Description                           |
|-------------------|---------------|-------------------|---------------------------------------|
| `dictionary`      | Field         | [1000]int         | array of words used while generating  |
| `Init`            | Static method | PasswordGenerator | returns PasswordGenerator object      |
| `charge`          | Method        | string            | generates password                    |
| `Create_password` | Method        | string            | generates password                    |

> Attention! you would rather use `Create_password` but not `charge`.

example of usage:
```go
import (
    PasswordGenerator "Server/database/PasswordGeneration"
)

passwordGenerator = PasswordGenerator.Init()

fmt.Println(passwordGenerator.Create_password())
```

example of output:
```
numbernumberdrivefulltwentyread
```

## TODO

Database
- [x] move user model to separate file
- [x] build DbConnector as a class but not a script as it is now
- [x] add ability to add new user to DbConnector
- [x] add ability to get user by tag to DbConnector
- [x] create invite table
- [x] dump updated database
- [x] add ability to create Invite by user
- [x] add ability to get Invite
- [x] implement function for creating word combinations for invites
- [x] add public key column in user table
- [ ] insert correct user id when `Create_invite`

Controller
- [x] design
- [x] add function for getting invites

main.go
- [x] gather all the components to the programm
- [ ] add https support
- [ ] make start script

tests
