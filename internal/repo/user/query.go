package user

const (
	queryGetUser     = "SELECT user_id, username, role, create_time, update_time FROM user"
	queryGetUserAuth = queryGetUser + " WHERE username = ? AND password = ?"
	queryGetUserByID = queryGetUser + " WHERE user_id = ?"

	queryInsertUser = "INSERT INTO user (username, password, role, create_time, update_time) VALUES (?, ?, ?, ?, ?)"

	queryGetUserDetail         = "SELECT detail_user_id, user_id, name, address, school_name, create_time, update_time FROM user_detail"
	queryGetUserDetailByUserID = queryGetUserDetail + " WHERE user_id = ?"

	queryInsertUserDetail = "INSERT INTO user_detail (user_id, name, address, school_name, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?)"
)
