package bindings

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lr LoginRequest) Validate() error {
	errs := new(RequestErrors)
	if lr.Username == "" {
		errs.Append(ErrUsernameEmpty)
	}
	if lr.Password == "" {
		errs.Append(ErrPasswordEmpty)
	}
	if errs.Len() == 0 {
		return nil
	}
	return errs
}
