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

## TODO

Database
- [ ] move user model to separate file
- [ ] build DbConnector as a class but not a script as it is now
- [ ] add abbility to add new user to DbConnector
- [ ] add abbility to get user by tag to DbConnector

Listener
- [ ] build

Responder
- [ ] build

main.go
- [ ] gather all the components to the programm

tests
