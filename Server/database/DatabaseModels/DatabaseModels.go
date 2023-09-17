//███╗░░░███╗░█████╗░██████╗░███████╗██╗░░░░░░██████╗
//████╗░████║██╔══██╗██╔══██╗██╔════╝██║░░░░░██╔════╝
//██╔████╔██║██║░░██║██║░░██║█████╗░░██║░░░░░╚█████╗░
//██║╚██╔╝██║██║░░██║██║░░██║██╔══╝░░██║░░░░░░╚═══██╗
//██║░╚═╝░██║╚█████╔╝██████╔╝███████╗███████╗██████╔╝
//╚═╝░░░░░╚═╝░╚════╝░╚═════╝░╚══════╝╚══════╝╚═════╝░

// This is a class of service's user.
// It is used for working with database

package DatabaseModels

import "github.com/google/uuid"

type User struct { //NOTE all properites are public which is not really secure
	Id               uuid.UUID `json:"id"`
	Tag              string    `json:"tag"`
	Allow_hostory    bool      `json:"allow_hostory"`
	Allow_screenshot bool      `json:"allow_screenshot"`
}
