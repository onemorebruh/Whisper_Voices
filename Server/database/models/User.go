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
	id               uuid.UUID
	tag              string
	allow_hostory    bool
	allow_screenshot bool
}
