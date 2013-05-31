package github

type User struct {
	Name string
}

func createUser(j map[string]interface{}) User {
	return User{
		Name: j["name"].(string),
	}
}
