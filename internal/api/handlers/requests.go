package handlers

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var UpdateUsernameRequest struct {
	UserRequest
	NewUsername string `json:"newUsername"`
}

var UpdatePasswordRequest struct {
	UserRequest
	NewPassword string `json:"newPassword"`
}

var UpdateUserRoleRequest struct {
	UserRequest
	NewRoleID uint `json:"newRoleId"`
}

var GetUsersByRoleRequest struct {
	RoleId uint `json:"roleId"`
}
