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

type (
	GetUsersByRoleRequest struct {
		RoleID uint `json:"roleId"`
	}

	UpdateRoleNameRequest struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	UpdateRoleDescriptionRequest struct {
		ID          uint   `json:"id"`
		Description string `json:"description"`
	}

	UpdateExpenseAmountRequest struct {
		ID     uint    `json:"id"`
		Amount float64 `json:"amount"`
	}

	UpdateExpenseDescriptionRequest struct {
		ID          uint   `json:"id"`
		Description string `json:"description"`
	}

	GetExpensesByUserRequest struct {
		UserID uint `json:"userId"`
	}

	GetWorkflowByApproverRequest struct {
		ApproverID uint `json:"approverId"`
	}

	GetWorkflowByExpenseRequest struct {
		ExpenseID uint `json:"expenseId"`
	}

	UpdateWorkflowCommentsRequest struct {
		ID       uint   `json:"id"`
		Comments string `json:"comments"`
	}

	UpdateWorkflowStatusRequest struct {
		ID     uint   `json:"id"`
		Status string `json:"status"`
	}
)
