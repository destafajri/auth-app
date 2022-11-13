package authentications

import "github.com/destafajri/auth-app/applications/entity"

type authRequestResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Role string `json:"role"`
	Password string `json:"password"`
}

func convertToAuthResponse(user entity.UserEntity) authRequestResponse {
	return authRequestResponse{
		ID: user.ID,
		Name: user.Name,
		Phone: user.Phone,
		Role: user.Role,
		Password: user.Password,
	}
}