package controller

import "net/http"

func getTokenJWT(r *http.Request) (int,string){
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return http.StatusUnauthorized,"token não informado"
	}
	tokenString = tokenString[len("Bearer "):]
	return http.StatusOK,tokenString
}