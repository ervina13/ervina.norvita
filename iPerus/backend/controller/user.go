package api

import (
	"encoding/json"
	"net/http"
	"IPERPUS/repository"
)

type ProfilUserErrorResponse struct {
	Error string `json:"error"`
}

type ProfiladminErrorResponse struct {
	Error string `json:"error"`
}

type ProfilUser struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
	Address   string `json:"address"`
	NoHp      string `json:"nohp"`
}

type AddProfilUserSuccessResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
	Address   string `json:"address"`
	NoHp      string `json:"nohp"`
}

type profiladmin struct{
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
	Address   string `json:"address"`
	NoHp      string `json:"nohp"`
}

type AddProfilAdminSuccessResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
	Address   string `json:"address"`
	NoHp      string `json:"nohp"`
}

func (api *controller) addProfilUser(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var profilUser ProfilUser
	encoder := json.NewEncoder(w)

	err := json.NewDecoder(r.Body).Decode(&profilUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("id").(int64)

	err = api.profilUserRepo.InsertProfilUser(repository.ProfilUser{
		Name:      profilUser.Name,
		Email:     profilUser.Email,
		Gender:    profilUser.Gender,
		Birthdate: profilUser.Birthdate,
		Address:   profilUser.Address,
		NoHp:      profilUser.NoHp,
		UserID:    userID,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ProfilUserErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AddProfilUserSuccessResponse{
		Name:      profilUser.Name,
		Email:     profilUser.Email,
		Gender:    profilUser.Gender,
		Birthdate: profilUser.Birthdate,
		Address:   profilUser.Address,
		NoHp:      profilUser.NoHp,
	})

}

func (api *controller) editProfiladmin(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var profiladmin profiladmin
	encoder := json.NewEncoder(w)

	err := json.NewDecoder(r.Body).Decode(&profiladmin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("id").(int64)

	err = api.profiladminRepo.UpdateProfiladmin(repository.Profiladmin{
		Name:      profiladmin.Name,
		Email:     profiladmin.Email,
		Address:   profiladmin.Address,
		NoHp:      profiladmin.NoHp,
		UserID:    userID,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ProfiladminErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AddProfilUserSuccessResponse{
		Name:      profiladmin.Name,
		Email:     profiladmin.Email,
		Gender:    profiladmin.Gender,
		Birthdate: profiladmin.Birthdate,
		Address:   profiladmin.Address,
		NoHp:      profiladmin.NoHp,
	})
}
func (api *controller) addProfilAdmin(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var profiladmin profiladmin
	encoder := json.NewEncoder(w)

	err := json.NewDecoder(r.Body).Decode(&profiladmin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("id").(int64)

	err = api.profiladminRepo.InsertProfiladmin(repository.Profiladmin{
		Name:      profiladmin.Name,
		Email:     profiladmin.Email,
		Address:   profiladmin.Address,
		NoHp:      profiladmin.NoHp,
		UserID:    userID,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ProfiladminErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AddProfilUserSuccessResponse{
		Name:      profiladmin.Name,
		Email:     profiladmin.Email,
		Gender:    profiladmin.Gender,
		Birthdate: profiladmin.Birthdate,
		Address:   profiladmin.Address,
		NoHp:      profiladmin.NoHp,
	})

}

func (api *controller) editProfilUser(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var profilUser ProfilUser
	encoder := json.NewEncoder(w)

	err := json.NewDecoder(r.Body).Decode(&profilUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("id").(int64)

	err = api.profilUserRepo.UpdateProfilUser(repository.ProfilUser{
		Name:      profilUser.Name,
		Email:     profilUser.Email,
		Gender:    profilUser.Gender,
		Birthdate: profilUser.Birthdate,
		Address:   profilUser.Address,
		NoHp:      profilUser.NoHp,
		UserID:    userID,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ProfilUserErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AddProfilUserSuccessResponse{
		Name:      profilUser.Name,
		Email:     profilUser.Email,
		Gender:    profilUser.Gender,
		Birthdate: profilUser.Birthdate,
		Address:   profilUser.Address,
		NoHp:      profilUser.NoHp,
	})
}
