//██╗░░░██╗░██████╗███████╗██████╗░
//██║░░░██║██╔════╝██╔════╝██╔══██╗
//██║░░░██║╚█████╗░█████╗░░██████╔╝
//██║░░░██║░╚═══██╗██╔══╝░░██╔══██╗
//╚██████╔╝██████╔╝███████╗██║░░██║
//░╚═════╝░╚═════╝░╚══════╝╚═╝░░╚═╝

// This is a class of service's user.
// It is used for working with database

package User

import "github.com/google/uuid"

type User struct {
	id               uuid.UUID `example:"id"`
	tag              string    `example:"tag"`
	allow_hostory    bool      `example:"allow_hostory"`
	allow_screenshot bool      `example:"allow_screenshot"`
}
