package util

import "math/rand"

func GenAnUser() User {
	users, err := Generate(&QueryConfig{1, ""})
	if err == nil && len(users) > 0 {
		return users[0]
	}
	return User{}
}

// RandomMoney generates a random amount of money
func RandomId() int32 {
	return rand.Int31()
}
