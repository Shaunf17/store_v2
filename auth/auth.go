package auth

import "net/http"

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func guest() *User {
	return &User{
		Username: "guest",
	}
}

func BasicAuth(r *http.Request) *User {
	u, p, ok := r.BasicAuth()
	if !ok {
		// handle no auth
	}

	_ = p // discard password for now

	return &User{
		Username: u,
	}
}
