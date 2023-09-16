//░░░░░██╗░██████╗░█████╗░███╗░░██╗██████╗░░█████╗░██████╗░██╗░░░██╗
//░░░░░██║██╔════╝██╔══██╗████╗░██║██╔══██╗██╔══██╗██╔══██╗╚██╗░██╔╝
//░░░░░██║╚█████╗░██║░░██║██╔██╗██║██████╦╝██║░░██║██║░░██║░╚████╔╝░
//██╗░░██║░╚═══██╗██║░░██║██║╚████║██╔══██╗██║░░██║██║░░██║░░╚██╔╝░░
//╚█████╔╝██████╔╝╚█████╔╝██║░╚███║██████╦╝╚█████╔╝██████╔╝░░░██║░░░
//░╚════╝░╚═════╝░░╚════╝░╚═╝░░╚══╝╚═════╝░░╚════╝░╚═════╝░░░░╚═╝░░░

//This is a class for using in http requests and responses

package JsonBody

import "Server/database/DatabaseModels"

type DBcommand int8

const (
	Add_user DBcommand = 0
	Get_user DBcommand = 1
)

type JsonBody struct {
	User    DatabaseModels.User `json:"user"`
	Address string              `json:"address"`
	Message string              `json:"message"`
	Command DBcommand           `json:"command"`
}
