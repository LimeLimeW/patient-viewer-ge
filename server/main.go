package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/rs/cors"
)

type Patient struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

var Patients []Patient

var jwtKey = []byte("toz")

var hashKey = []byte("very-secret")
var blockKey = []byte("a-lot-secret")
var s = securecookie.New(hashKey, blockKey)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requete : root")
}

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
}

func Login(w http.ResponseWriter, r *http.Request) {

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(tokenString)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Secure:  true,
	})
	fmt.Println("Endpoint Hit: token")
}

func createNewPatient(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var patient Patient

	json.Unmarshal(reqBody, &patient)

	Patients = append(Patients, patient)

	json.NewEncoder(w).Encode(patient)

	fmt.Fprintf(w, "%+v", string(reqBody))
}

func returnPatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["email"]

	for _, patient := range Patients {
		if patient.Email == key {
			json.NewEncoder(w).Encode(patient)
		}
	}
}

func Refresh(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/patients", returnAllPatients)
	router.HandleFunc("/patient/{email}", returnPatient)
	router.HandleFunc("/patient", createNewPatient).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":10000", handler))
}

func returnAllPatients(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requete : returnAllPatients")
	json.NewEncoder(w).Encode(Patients)
}

func main() {
	Patients = []Patient{
		Patient{Firstname: "Janine", Lastname: "LAPUNAISE", Birthday: "06/08/1999", Height: "168cm", Weight: "54kg", Email: "me@ethereal.pw", Gender: "F"},
	}
	handleRequests()
}
