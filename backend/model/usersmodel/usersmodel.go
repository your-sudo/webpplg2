package usersmodel

import (
	"github.com/your-sudo/webkelaspplg2/config"
	"github.com/your-sudo/webkelaspplg2/entities"
)

func TampilUser() []entities.Users {
	rows, err := config.DB.Query("SELECT id, nama, keterangan, ig, foto FROM users")
	if err != nil {
		panic(err)

	}
	defer rows.Close()

	var users []entities.Users
	for rows.Next() {
		var u entities.Users
		if err := rows.Scan(&u.ID, &u.Nama, &u.Keterangan, &u.Ig, &u.Foto); err != nil {
			panic(err)
		}
		users = append(users, u)
	}
	return users
}

func LoginUser() []entities.Login {
	rows, err := config.DB.Query("SELECT username, password FROM users")
	if err != nil {
		panic(err)

	}
	defer rows.Close()

	var login []entities.Login

	for rows.Next() {
		var l entities.Login
		if err := rows.Scan(&l.Username, &l.Password); err != nil {
			panic(err)
		}
		login = append(login, l)
	}
	return login

}