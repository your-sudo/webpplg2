package homecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/your-sudo/webkelaspplg2/model/usersmodel"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	users := usersmodel.TampilUser()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

