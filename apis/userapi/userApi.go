package userapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PhongVX/golang-rest-api/entities"
)

func Login(response http.ResponseWriter, request *http.Request) {
	fmt.Println("redirected")
	var user entities.Owner
	_ = json.NewDecoder(request.Body).Decode(&user)
	fmt.Println(user.Password)
	fmt.Println(user.Username)

	requestBody, err := json.Marshal(map[string]string{
		"name":          user.Username,
		"password":      user.Password,
		"grant-type":    "password",
		"client-secret": "secret-key",
	})

	if err != nil {
		fmt.Println("test 1")
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		resp, err := http.Post("http://127.0.0.1:3000/authorize", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("test 2")
			responseWithError(response, http.StatusBadRequest, err.Error())
		}
		defer resp.Body.Close()

	}
}

func GetResource(response http.ResponseWriter, request *http.Request) {
	var user entities.Response
	_ = json.NewDecoder(request.Body).Decode(&user)
	url := "http://127.0.0.1:3000/api/resource"
	var bearer = "Bearer " + user.AccessToken
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string([]byte(body)))
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
