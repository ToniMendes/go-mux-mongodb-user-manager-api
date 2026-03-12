package users_manager

import (
	"go-mux-mongodb-user-manager-api/internal/domain"
)

type UpdateEmailServices struct {
	Collection      domain.MongoRepository
	HashingServices domain.HashingRepository
}

func NewUpdateEmailServices(collection domain.MongoRepository, hashingServices domain.HashingRepository) *UpdateEmailServices {
	return &UpdateEmailServices{
		Collection:      collection,
		HashingServices: hashingServices,
	}
}

func (r *UpdateNameServices) ExecUpdateEmail(inputDto UserUpdateEmailInput) (map[string]interface{}, error) {
	model := domain.NewUpdateEmail(
		inputDto.Email,
		inputDto.Password,
		inputDto.NewEmail,
	)

	result, err := r.Collection.GetByEmail(model.Email)
	if err != nil {
		return map[string]interface{}{}, err
	}

	err = r.HashingServices.ComparePassword(result.Password, model.Password)
	if err != nil {
		return map[string]interface{}{}, err
	}

	err = r.Collection.UpdateEmail(model.NewEmail, model.Email)
	if err != nil {
		return map[string]interface{}{}, err
	}

	response := UserUpdateEmailResponse{
		Id:       result.Id,
		Email:    result.Email,
		NewEmail: model.NewEmail,
	}

	newResponse := map[string]interface{}{
		"update_success": map[string]interface{}{
			"user": map[string]interface{}{
				"id":    response.Id.Hex(),
				"email": response.Email,
				"new_email": map[string]string{
					"email": response.NewEmail,
				},
			},
		},
	}

	return newResponse, nil
}
