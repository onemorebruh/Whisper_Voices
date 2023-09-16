//███╗░░░███╗░█████╗░██████╗░███████╗██╗░░░░░░██████╗
//████╗░████║██╔══██╗██╔══██╗██╔════╝██║░░░░░██╔════╝
//██╔████╔██║██║░░██║██║░░██║█████╗░░██║░░░░░╚█████╗░
//██║╚██╔╝██║██║░░██║██║░░██║██╔══╝░░██║░░░░░░╚═══██╗
//██║░╚═╝░██║╚█████╔╝██████╔╝███████╗███████╗██████╔╝
//╚═╝░░░░░╚═╝░╚════╝░╚═════╝░╚══════╝╚══════╝╚═════╝░

// Here are classes used for working with database responses

package DatabaseModels

import "github.com/google/uuid"

// This is a class of service's user.
// It is used for working with database

type User struct { //NOTE all properites are public which is not really secure
	Id               uuid.UUID `json:"id"`
	Tag              string    `json:"tag"`
	Allow_hostory    bool      `json:"allow_hostory"`
	Allow_screenshot bool      `json:"allow_screenshot"`
}

// This is a class of user's key.
// User have to use key each time they do any action in via Server.
// User can invite new user they need to create key for each one.

type Key struct {
	Id    uuid.UUID `json:"id"`
	User  uuid.UUID `json:"user"`
	Value string    `json:"value"`
}
