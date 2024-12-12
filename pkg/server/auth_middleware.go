package server

import (
	"net/http"
	inertia "github.com/romsar/gonertia"
)


type AccountProp struct {
	Name string `json:"name"`
}

type UserProp struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Account   AccountProp`json:"account"`
}

type AuthProp struct {
	User UserProp `json:"user"`
}

type AuthMiddleware struct {
	i *inertia.Inertia
}

func NewAuthMiddleware(i *inertia.Inertia) *AuthMiddleware {
	return &AuthMiddleware{i}
}

func (am AuthMiddleware) sharedPropMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inertia.SetProp(r.Context(), "auth", AuthProp{
			User: UserProp{
				FirstName: "John",
				LastName:  "Doe",
				Account: AccountProp{
					Name: "Acme Corporation",
				},
			},
		})
		next.ServeHTTP(w, r)
	})
}
