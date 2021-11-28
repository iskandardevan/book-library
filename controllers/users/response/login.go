package response

import "github.com/iskandardevan/book-library/business/users"

func UserLogin(domain users.Domain, token string) JWTResponse{
	Response := UserResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		// Password:  domain.Password,
		Token:     domain.Token,
	}

	TokenResponse := JWTResponse{}
	TokenResponse.Token = token
	TokenResponse.User = Response
	return TokenResponse

}

