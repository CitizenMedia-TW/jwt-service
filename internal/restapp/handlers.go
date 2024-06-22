package restapp

import (
	"encoding/json"
	"jwt-service/internal/helper"
	"jwt-service/internal/models"
	"jwt-service/protobuffs/jwt-service"
	"net/http"
	"strings"
)

func (s *RestServer) GenerateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.httpError(&w, "Method not allows", http.StatusBadRequest)
		return
	}

	var signContent models.JWTContent // Declare the content to sign

	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			s.httpError(&w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, v := range []string{"mail", "name"} {
			if !r.PostForm.Has(v) {
				s.httpError(&w, v+" is required", http.StatusBadRequest)
				return
			}
		}

		signContent = models.JWTContent{
			Mail: r.PostFormValue("mail"),
			Name: r.PostFormValue("name"),
		}
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&signContent)
		if err != nil {
			s.httpError(&w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		s.httpError(&w, "Content-Type not allowed", http.StatusBadRequest)
		return
	}

	tokenString, err := helper.JWTSignContent(signContent, s.config.Secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// res := struct {
	// 	Message string `json:"message"`
	// 	Token   string `json:"token"`
	// }{Message: "Success", Token: "Bearer " + tokenString}
	res := &jwt.GenerateTokenResponse{
		Message: "Success",
		Token:   "Bearer " + tokenString,
	}
	json.NewEncoder(w).Encode(res)

	return
}

func (s *RestServer) VerifyToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.httpError(&w, "Method not allows", http.StatusBadRequest)
		return
	}
	if strings.Split(r.Header.Get("Authorization"), " ")[0] != "Bearer" {
		s.httpError(&w, "Authorization header is required", http.StatusBadRequest)
		return
	}

	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	claims, err := helper.JWTParseToken(tokenString, s.config.Secret)
	if err != nil {
		s.httpError(&w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// res := struct {
	// 	Message    string             `json:"message"`
	// 	JWTContent *models.JWTContent `json:"jwtContent"`
	// }{Message: "Success", JWTContent: claims}
	res := &jwt.VerifyTokenResponse{
		Message: "Success",
		JwtContent: &jwt.JwtContent{
			Mail: claims.Mail,
			Name: claims.Name,
		}}
	json.NewEncoder(w).Encode(res)

	return
}
