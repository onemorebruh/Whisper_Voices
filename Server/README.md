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
| `Id`              | field         | uuid.UUID     | identificator     |
| `Tag`             | field         | string        | unique user tag   |
| `Allow_history`   | field         | bool          | check User scheme |
| `Allow_screenshot`| field         | bool          | check User scheme |

#### DBcommand

DBcommand is a Enum

| Possible values   |
|-------------------|
| `Add_user`        |
| `Get_user`        |

### DatabaseResponse

DatabaseResponse is used in DatabaseConnector to get data from methods

| Name              | Field/Method  | Type          | Description                   |
|-------------------|---------------|---------------|-------------------------------|
| `Message`         | field         | string        | status message                |
| `Is_successful`   | field         | bool          | was communication successful  |
| `User`            | field         | User.User     | user gotten by tag            |

### DatabaseConnector - ConnectionSettings

ConnectionSettings is used for communication with mysql database
It is not named as DatabaseConnector because there is such class in one of dependencies

| Name              | Field/Method  | Type              | Description                           |
|-------------------|---------------|-------------------|---------------------------------------|
| `Database`        | field         | string            | database name                         |
| `Password`        | field         | string            | database user password                |
| `Host`            | field         | string            | database host                         |
| `Port`            | field         | string            | database port                         |
| `User`            | field         | string            | database user                         |
| `is_set`          | method        | bool              | checks if all fields are set          |
| `does_user_exists`| method        | DatabaseResponse  | checks if user with given tag exists  |
| `insert_user`     | method        | DatabaseResponse  | inserts user into database            |
| `Add_user`        | method        | bool              | inserts user into database            |

> Attention! It is not secure to use `insert_user`. you should rather use `Add_user`

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
- [ ] implement function for creating word combinations for invites
- [ ] make start script

Controller
- [x] design
- [ ] add https support
- [ ] add function for getting invites

main.go
- [ ] gather all the components to the programm

tests
