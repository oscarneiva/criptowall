package main

import (
	"context"
	"encoding/json"
	"fmt"
	httptransport 	"github.com/go-kit/kit/transport/http"
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
	repository := apirepositories.NewUserCSVRepository()
	service := apiservices.NewUserService(repository)

	router := mux.NewRouter()

	createEndPoint := httptransport.NewServer(MakeCreateEndpoint(service), DecodeUserFromBody, EncodeResponse)
	router.Handle("/api/users", createEndPoint).Methods(http.MethodPost)

	searchEndPoint := httptransport.NewServer(MakeSearchEndpoint(service), DecodeNothing, EncodeResponse)
	router.Handle("/api/users/search", searchEndPoint).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

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

func DecodeNothing(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func MakeCreateEndpoint(userService services.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		user := request.(*entities.User)
		return userService.Create(ctx, user)
	}
}

func MakeSearchEndpoint(userService services.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return userService.Search(ctx)
	}
}

func Things()  {
	fmt.Println("hello world")

	repository := apirepositories.NewUserCSVRepository()
	service := apiservices.NewUserService(repository)

	service.Create(context.TODO(), &entities.User{
		ID: "01",
		Name: "Elton Tusk",
		Wallets: nil,
	})
	service.Create(context.TODO(), &entities.User{
		ID: "02",
		Name: "Jiz Wezos",
		Wallets: nil,
	})
	user, err := service.GetByID(context.TODO(), "02")
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	fmt.Println(user)
}