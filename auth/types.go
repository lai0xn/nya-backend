package auth

type UserType struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserType) validate() error {
	validator := Validator{}
	if err := validator.ValidateEmail(u.Email); err != nil {
		return err
	}
	if err := validator.ValidateEmail(u.Email); err != nil {
		return err
	}
	return nil
}

type LoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
