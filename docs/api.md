# Boss Payback API Documentation

This documentation provides an overview of the API routes, models, and handlers for the Boss Payback system.

## Table of Contents

1. [Routes](#routes)
    - [User Routes](#user-routes)
    - [Role Routes](#role-routes)
    - [Expense Routes](#expense-routes)
    - [Workflow Routes](#workflow-routes)
2. [Models](#models)
    - [User Model](#user-model)
    - [Role Model](#role-model)
    - [Expense Model](#expense-model)
    - [Workflow Model](#workflow-model)
3. [Handlers](#handlers)
4. [Request Structs](#request-structs)

---

## Routes

### User Routes

- **GET `/users`**: Retrieve users based on role.
  
- **POST `/login`**: Login with username and password.

#### Admin Routes:

- All routes under `/admin/user` require an admin token for authorization.
  - **POST `/register`**: Register a new user.
  - **PUT `/username`**: Update username.
  - **PUT `/password`**: Update user password.
  - **PUT `/role`**: Update user role.
  - **DELETE `/delete`**: Delete user.

### Role Routes

- **GET `/roles`**: Retrieve all roles.

#### Admin Routes:

- All routes under `/admin/role` require an admin token for authorization.
  - **POST `/role/new`**: Create a new role.
  - **PUT `/name`**: Update role name.
  - **PUT `/description`**: Update role description.
  - **DELETE `/delete`**: Delete a role.

### Expense Routes

- **GET `/expenses`**: Retrieve all expenses.
- **GET `/expenses/user`**: Retrieve expenses by user.

#### General Token Routes:

- All routes under `/expense` require a general token for authorization.
  - **POST `/new`**: Create a new expense.
  - **PUT `/amount`**: Update expense amount.
  - **PUT `/description`**: Update expense description.

#### Admin Routes:

- All routes under `/admin/expense` require an admin token for authorization.
  - **DELETE `/delete`**: Delete an expense.

### Workflow Routes

- **GET `/workflows`**: Retrieve all workflows.
- **GET `/workflow/approver`**: Retrieve workflows by approver.
- **GET `/workflow/expense`**: Retrieve workflows by expense.

#### Admin and Approver Token Routes:

- All routes under `/workflow` and `/admin/workflow` require admin and approver tokens for authorization.
  - **POST `/new`**: Create a new workflow.
  - **PUT `/status`**: Update workflow status.
  - **PUT `/comments`**: Update workflow comments.

#### Admin Routes:

- All routes under `/admin/workflow` require an admin token for authorization.
  - **DELETE `/delete`**: Delete a workflow.

---

## Models

### User Model

```go
type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `json:"username" gorm:"unique;not null"`
    Password string `json:"password" gorm:"not null"`
    Email    string `json:"email" gorm:"unique;not null"`
    RoleID   uint   `json:"roleId"`
    Role     Role
}
```

### Role Model

```go
type Role struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `json:"name" gorm:"unique;not null"`
    Description string `json:"description"`
}
```

### Expense Model

```go
type Expense struct {
    ID           uint    `gorm:"primaryKey"`
    UserID       uint    `json:"userId"`
    Description  string  `json:"description" gorm:"not null"`
    Amount       float64 `json:"amount" gorm:"not null"`
    Status       string  `json:"status" gorm:"default:'pending'"`
    ApproverID   uint    `json:"approverId"`
    ReceiptImage string  `json:"receiptImage"`
}
```

### Workflow Model

```go
type Workflow struct {
    ID         uint   `gorm:"primaryKey"`
    ExpenseID  uint   `json:"expenseId"`
    ApproverID uint   `json:"approverId"`
    Status     string `json:"status" gorm:"default:'pending'"`
    Comments   string `json:"comments"`
}
```

## Handlers

- User Handlers: Include registration, login, and user management.
- Role Handlers: Manage role creation, update, and deletion.
- Expense Handlers: Handle expense creation, retrieval, and management.
- Workflow Handlers: Handle workflow creation, retrieval, and management.

## Request Structs

- UserRequest: Contains username and password for user-related operations.
- UpdateUsernameRequest: Includes updated username along with the user request.
- UpdatePasswordRequest: Includes updated password along with the user request.
- UpdateUserRoleRequest: Contains role ID for updating user role.
- GetUsersByRoleRequest: Contains role ID for retrieving users by role.
- UpdateRoleNameRequest: Includes role ID and new role name.
- UpdateRoleDescriptionRequest: Includes role ID and updated description.
- UpdateExpenseAmountRequest: Contains expense ID and updated amount.
- UpdateExpenseDescriptionRequest: Contains expense ID and updated description.
- GetExpensesByUserRequest: Contains user ID for retrieving expenses.
- GetWorkflowByApproverRequest: Contains approver ID for retrieving workflows.
- GetWorkflowByExpenseRequest: Contains expense ID for retrieving workflows.
- UpdateWorkflowCommentsRequest: Includes workflow ID and updated comments.
- UpdateWorkflowStatusRequest: Includes workflow ID and updated status.