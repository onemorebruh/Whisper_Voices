package DatabaseResponse

import User "Server/database/DatabaseModels"

type DatabaseResponse struct {
	Message       string
	Is_successful bool
	User          User.User
}
