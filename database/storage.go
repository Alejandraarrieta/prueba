package storage
import(
	m "../../models"
	"fmt"
	"net/http"
	"encoding/json"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json")
	var user m.UserData
	_ = json.NewDecoder(r.Body).Decode(&user)
	collection := client.Database("UserDB").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _:= collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)
}


func SearchUser(username string) []m.UserData {
	var users []m.UserData
	collection := client.Database("UserDB").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"email":username}
	options := options.Find()
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
	var user m.UserData
	cursor.Decode(&user)
	users = append(users, user)
	fmt.Println("Error", users)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("SearchUser", err)
	}
	return users
}
