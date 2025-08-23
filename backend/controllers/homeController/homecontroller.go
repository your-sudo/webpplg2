package homecontroller

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/your-sudo/webkelaspplg2/config"
)

type user struct {
	ID         int    `json:"id"`
	Nama       string `json:"name"`
	Keterangan string `json:"keterangan"`
	Ig         string `json:"ig"`
	Foto       string `json:"foto"`
}

func Welcome(w http.ResponseWriter, r *http.Request) {

	rows, err := config.DB.Query("SELECT id, nama, keterangan, ig, foto FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Nama, &u.Keterangan, &u.Ig, &u.Foto); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
