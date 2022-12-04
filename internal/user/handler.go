package user

import (
	"Rest-Api-App/internal/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURl      = "/users"
	userURL       = "/users/:uuid"
	createUserURL = "/create"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURl, h.GetList)
	router.POST(userURL, h.GetUserByUUID)
	router.GET(createUserURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(usersURl, h.PartiallyUpdateUser)
	router.DELETE(usersURl, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user by UUid"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Create User"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Update User"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Partially Update User"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("delete user"))
}
