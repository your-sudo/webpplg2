package logincontroller

import (
	"net/http"
	"github.com/your-sudo/webkelaspplg2/model/usersmodel"
)

func Login(w http.ResponseWriter, r *http.Request) {
	

}

func LoginValidation(w http.ResponseWriter, r *http.Request) {
	usersmodel.LoginUser()

}