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

type GetUsersByRoleRequest struct {
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

var GetWorkflowByApproverRequest struct {
	ApproverID uint `json:"approverId"`
}

var GetWorkflowByExpenseRequest struct {
	ExpenseID uint `json:"expenseId"`
}

var UpdateWorkflowCommentsRequest struct {
	ID       uint   `json:"id"`
	Comments string `json:"comments"`
}

var UpdateWorkflowStatusRequest struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}
