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

var UpdateRoleNameRequest struct {
	ID      uint   `json:"id"`
	NewName string `json:"newName"`
}

var UpdateRoleDescriptionRequest struct {
	ID             uint   `json:"id"`
	NewDescription string `json:"newDescription"`
}

var UpdateExpenseAmountRequest struct {
	ID     uint    `json:"id"`
	Amount float64 `json:"amount"`
}

var UpdateExpenseDescriptionRequest struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

var DeleteExpenseRequest struct {
	ID uint `json:"id"`
}
