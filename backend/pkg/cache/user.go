package cache

import (
	"nyctaxi_mapup/pkg/model"
	"os"
)

var (
	cache = make(map[string]model.UserCache)
)

func GetUserCredentials(username string) (model.UserCache, bool) {
	return cache[username], true
}

func InitUserCredentials() {
	cache[os.Getenv("ADMIN_USERNAME")] = model.UserCache{
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Role:     os.Getenv("ADMIN_ROLE"),
	}
	cache[os.Getenv("MANAGER_USERNAME")] = model.UserCache{
		Username: os.Getenv("MANAGER_USERNAME"),
		Password: os.Getenv("MANAGER_PASSWORD"),
		Role:     os.Getenv("MANAGER_ROLE"),
	}
	cache[os.Getenv("USER_USERNAME")] = model.UserCache{
		Username: os.Getenv("USER_USERNAME"),
		Password: os.Getenv("USER_PASSWORD"),
		Role:     os.Getenv("USER_ROLE"),
	}

}

func SetJWTToken(username, token string) {
	usercache := cache[username]
	usercache.Token = token

	cache[username] = usercache
}
func GetPassword(username string) string {
	return cache[os.Getenv("USER_USERNAME")].Password
}

// GetRole returns the role.
func GetRole() string {
	return cache[os.Getenv("USER_USERNAME")].Role
}
