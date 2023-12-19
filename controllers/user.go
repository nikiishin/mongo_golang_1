package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo_golang/model"
	"net/http"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	collection := uc.client.Database("mongo-golang").Collection("users")
	var u model.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		w.WriteHeader(404)

	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Создаем новый экземляр каласса Юзер
	u := model.User{}
	//декодируем тело запроса в структуру User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Получаем коллекцию User

	collection := uc.client.Database("mongo-golang").Collection("users")

	// Вставте нового пользователя в коллекцию
	insertResult, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// отправляем ответ

	json.NewEncoder(w).Encode(insertResult.InsertedID)

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}
	collection := uc.client.Database("mongo-golang").Collection("users")
	deleteResults, err := collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(deleteResults.DeletedCount)
}
