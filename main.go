package main
import(
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"time"
	"context"
	"fmt"
    "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	 jwt "github.com/golang-jwt/jwt/v4"
	 //"strings"
	 "/database"
	 "/jwt"
	 "/models"
)
var client *mongo.Client

func main(){

	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	http.ListenAndServe(":8080",router)
}

func Login(w http.ResponseWriter, r *http.Request){	
	var username UserLog
	var userstruct UserData
	var usu [] UserData
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		fmt.Println("ERROR", err)
	} else {
	user := username.Username
	fmt.Println("EMAIL", user)
	usu = SearchUser(user)
	fmt.Println("USER BUSCADO", usu)
	fmt.Println("Cantidad Elementos", len(usu))
	if len(usu) > 0 {
	userstruct = UserData {
			Name: usu[0].Name,
			Lastname: usu[0].Lastname,
			Email: usu[0].Email,
			Workunit: usu[0].Workunit,
			Photo: usu[0].Photo,
			Phone:usu[0].Phone,
			Area:usu[0].Area,
			Role:usu[0].Role,
		}
	fmt.Println("USER ESTRUCT", userstruct)
	}
	}
	token, _ := GenerateJWT(userstruct)
	result := ResponseToken{token}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintln(w, "Error al generar json")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}
