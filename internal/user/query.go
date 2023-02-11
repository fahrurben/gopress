package user

const (
	InsertUser     = "INSERT INTO users(email, name, password, type, created_at, updated_at) values (?, ?, ?, ?, ?, ?)"
	UpdateUser     = "UPDATE users SET name=?, updated_at=? WHERE id=?"
	FindUserById   = "SELECT id, email, name, password, type, created_at, updated_at, deleted_at FROM users WHERE id=?"
	FindAllUser    = "SELECT id, email, name, password, type, created_at, updated_at, deleted_at FROM users ORDER BY name limit ? offset ?"
	deleteUserById = "DELETE FROM users WHERE id=?"
)
