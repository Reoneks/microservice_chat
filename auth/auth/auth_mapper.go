package auth

func FromAuthUserDTO(authUserDTO *AuthUserDTO) *AuthUser {
	return &AuthUser{
		ID:       authUserDTO.ID,
		Email:    authUserDTO.Email,
		Password: authUserDTO.Password,
	}
}

func FromAuthUserDTOs(AuthUserDTOs []AuthUserDTO) (users []AuthUser) {
	for _, dto := range AuthUserDTOs {
		users = append(users, *FromAuthUserDTO(&dto))
	}
	return
}

func ToAuthUserDTO(user *AuthUser) *AuthUserDTO {
	return &AuthUserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToAuthUserDTOs(users []AuthUser) (AuthUserDTOs []AuthUserDTO) {
	for _, dto := range users {
		AuthUserDTOs = append(AuthUserDTOs, *ToAuthUserDTO(&dto))
	}
	return
}
