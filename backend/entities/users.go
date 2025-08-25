package entities


	type Users struct {
		ID         int    `json:"id"`
		Nama       string `json:"name"`
		Keterangan string `json:"keterangan"`
		Ig         string `json:"ig"`
		Foto       string `json:"foto"`
	}

	type LoginValidator struct {
		Username string 
		Password string 
	}

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`	
	}

	type loginResponse struct {
		Message string `json:"message"`
	}