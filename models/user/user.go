package user

type User struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"-"`
	AccessToken string `json:"access_token"`
}

func GetTableName() string {
	return "users"
}

func GetLoginField() []string {
	return []string{
		"id", "name", "password",
	}
}
