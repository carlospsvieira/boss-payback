package handlers

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var UpdateUsernameRequest struct {
	UserRequest
	UpdatedUsername string `json:"updatedUsername"`
}

var UpdatePasswordRequest struct {
	UserRequest
	UpdatedPassword string `json:"updatedPassword"`
}

var UpdateUserRoleRequest struct {
	UserRequest
	RoleID uint `json:"roleId"`
}

var GetUsersByRoleRequest struct {
	RoleID uint `json:"roleId"`
}

var UpdateRoleNameRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

var UpdateRoleDescriptionRequest struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

var UpdateExpenseAmountRequest struct {
	ID     uint    `json:"id"`
	Amount float64 `json:"amount"`
}

var UpdateExpenseDescriptionRequest struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

var GetExpensesByUserRequest struct {
	UserID uint `json:"userId"`
}

var DeleteExpenseRequest struct {
	ID uint `json:"id"`
}

var GetWorkflowByApproverRequest struct {
	ApproverID uint `json:"approverId"`
}

var UpdateWorkflowCommentsRequest struct {
	ID       uint   `json:"id"`
	Comments string `json:"comments"`
}

var DeleteWorkflowRequest struct {
	ID uint `json:"id"`
}
