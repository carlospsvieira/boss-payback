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
		Name string `json:"name"`
	}

	UpdateRoleDescriptionRequest struct {
		Description string `json:"description"`
	}

	UpdateExpenseAmountRequest struct {
		Amount float64 `json:"amount"`
	}

	UpdateExpenseDescriptionRequest struct {
		Description string `json:"description"`
	}

	UpdateWorkflowCommentsRequest struct {
		Comments string `json:"comments"`
	}

	UpdateWorkflowStatusRequest struct {
		Status string `json:"status"`
	}
)
