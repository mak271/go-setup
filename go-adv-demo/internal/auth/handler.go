package auth

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/request"
	"go/adv-demo/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Read body
		payload, err := request.HandleBody[LoginRequest](w, req)
		if err != nil {
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, err := request.HandleBody[RegisterRequest](w, req)
		if err != nil {
			return
		}
		fmt.Println(payload)
		data := RegisterResponse{
			Token: "321",
		}
		res.Json(w, data, http.StatusOK)
	}
}
