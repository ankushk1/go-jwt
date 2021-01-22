package main

import (
	"./model"
	"fmt"
	"./controller"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"context"
	"encoding/json"

	// "github.com/gorilla/mux"

	jwt "github.com/dgrijalva/jwt-go"
	
)

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	var result model.User
	var res model.ResponseResult
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.Username = claims["username"].(string)
		result.FirstName = claims["firstname"].(string)
		result.LastName = claims["lastname"].(string)

		json.NewEncoder(w).Encode(result)
		return
	} else {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	// r.HandleFunc("/login", controller.getUser).Methods("POST")
	r.HandleFunc("/user", getUser).Methods("GET")

	fmt.Println("server running on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
	
}