package entities


	type Users struct {
		ID         int    `json:"id"`
		Nama       string `json:"name"`
		Keterangan string `json:"keterangan"`
		Ig         string `json:"ig"`
		Foto       string `json:"foto"`
	}
