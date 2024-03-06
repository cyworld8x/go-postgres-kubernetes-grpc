package util

func GenAnUser() User {
	users, err := Generate(&QueryConfig{1, ""})
	if err == nil && len(users) > 0 {
		return users[0]
	}
	return User{}
}
