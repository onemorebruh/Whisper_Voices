package DatabaseResponse

import User "DBC/DatabaseModels"

type DatabaseResponse struct {
	Message       string
	Is_successful bool
	User          User.User
}
