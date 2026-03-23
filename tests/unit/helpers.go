package helpers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func GetTimeFromRFC3339(dateStr string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	if err!= nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func ExtractRouteParams(r *http.Request) map[string]string {
	vars := mux.Vars(r)
	return vars
}

func GetStatusCode(status string) int {
	switch strings.ToLower(status) {
	case "ok":
		return http.StatusOK
	case "created":
		return http.StatusCreated
	case "not_found":
		return http.StatusNotFound
	case "unauthorized":
		return http.StatusUnauthorized
	case "forbidden":
		return http.StatusForbidden
	case "unprocessable_entity":
		return http.StatusUnprocessableEntity
	case "internal_server_error":
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func ExtractTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", fmt.Errorf("token not found in request header")
	}
	return strings.TrimSpace(strings.Replace(token, "Bearer ", "", 1)), nil
}

func GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func ValidateToken(token string) (*jwt.Token, error) {
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if tokenClaims == nil ||!tokenClaims.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	return tokenClaims, nil
}

func GetIntFromQuery(r *http.Request, param string) (int, error) {
	query := r.URL.Query()
	value := query.Get(param)
	if value == "" {
		return 0, fmt.Errorf("parameter %s not found in query string", param)
	}
	intValue, err := strconv.Atoi(value)
	if err!= nil {
		return 0, err
	}
	return intValue, nil
}