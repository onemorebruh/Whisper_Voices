package DatabaseResponse

import User "Server/database/models"

type DatabaseResponse struct {
	message      string
	is_succesful bool
	user         User.User
}
