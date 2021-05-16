package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/oscarneiva/apilesson/apirepositories"
	"github.com/oscarneiva/apilesson/apiservices"
	"github.com/oscarneiva/apilesson/domain/entities"
	"github.com/oscarneiva/apilesson/domain/services"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running...")
	repository := apirepositories.NewUserCSVRepository()
	service := apiservices.NewUserService(repository)

	router := mux.NewRouter()

	createEndPoint := httptransport.NewServer(MakeCreateEndpoint(service), DecodeUserFromBody, EncodeResponse)
	router.Handle("/api/users", createEndPoint).Methods(http.MethodPost)

	getByIdEndPoint := httptransport.NewServer(MakeGetByIdEndPoint(service), DecodeIDFromURL, EncodeResponse)
	router.Handle("/api/users/{id}", getByIdEndPoint).Methods(http.MethodGet)

	searchEndPoint := httptransport.NewServer(MakeSearchEndpoint(service), DecodeNothing, EncodeResponse)
	router.Handle("/api/users/search", searchEndPoint).Methods(http.MethodGet)





	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err.Error())
	}

}

// Decode user
func DecodeUserFromBody(ctx context.Context, r *http.Request) (interface{}, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	user := &entities.User{}

	err = json.Unmarshal(bytes, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Create user end point
func MakeCreateEndpoint(userService services.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		user := request.(*entities.User)
		return userService.Create(ctx, user)
	}
}

// Get user by id end point
func MakeGetByIdEndPoint(userService services.UserService) func(ctx context.Context, request interface{}) (interface{}, error){
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		return userService.GetByID(ctx, id)
	}
}

// Search user end point
func MakeSearchEndpoint(userService services.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return userService.Search(ctx)
	}
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}


func DecodeNothing(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeIDFromURL(ctx context.Context, r *http.Request) (interface{}, error){
	uriParams := mux.Vars(r)

	id := uriParams["id"]
	if id == "" {
		err := errors.New("Error: blank id")
		return nil,err
	}

	return id, nil
}
