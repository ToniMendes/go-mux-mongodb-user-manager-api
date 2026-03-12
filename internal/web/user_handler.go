package web

import (
	"encoding/json"
	"go-mux-mongodb-user-manager-api/internal/usecases/repository"
	"go-mux-mongodb-user-manager-api/internal/usecases/users_manager"
	"net/http"
)

type UserUseCases interface {
	repository.WriterOnlyRepositoryUseCases
	repository.ReadOnlyRepositoryUseCases
}

type UserUseCasesRepository struct {
	usecase UserUseCases
}

func NewUserUseCasesRepository(uc UserUseCases) *UserUseCasesRepository {
	return &UserUseCasesRepository{
		usecase: uc,
	}
}

func (repo *UserUseCasesRepository) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var inputDto users_manager.UserCreateInput

	if err := json.NewDecoder(r.Body).Decode(&inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := users_manager.Validate(inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := repo.usecase.ExecCreate(inputDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"new_user": response,
	})
}

func (repo *UserUseCasesRepository) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	response, err := repo.usecase.ExecGetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users": response,
	})
}

func (repo *UserUseCasesRepository) LoginUser(w http.ResponseWriter, r *http.Request) {
	var inputDto users_manager.UserLoginInput

	if err := json.NewDecoder(r.Body).Decode(&inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := users_manager.Validate(inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := repo.usecase.ExecLogin(inputDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"login_success": response,
	})

}

func (repo *UserUseCasesRepository) UpdateName(w http.ResponseWriter, r *http.Request) {
	var inputDto users_manager.UserUpdateNameInput

	if err := json.NewDecoder(r.Body).Decode(&inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := users_manager.Validate(inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := repo.usecase.ExecUpdateName(inputDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (repo *UserUseCasesRepository) UpdateEmail(w http.ResponseWriter, r *http.Request) {
	var inputDto users_manager.UserUpdateEmailInput

	if err := json.NewDecoder(r.Body).Decode(&inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := users_manager.Validate(inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := repo.usecase.ExecUpdateEmail(inputDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
