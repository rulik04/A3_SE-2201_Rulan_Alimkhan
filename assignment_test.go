package greenlight

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

type Movie struct {
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Runtime  string   `json:"runtime"`
	Genres   []string `json:"genres"`
}

func TestUsersCreationEndpoint(t *testing.T) {
	dataToSend, _ := json.Marshal(map[string]string{
		"Name": "Rulan",
		"Email":    "RulanA@alimhan.kz",
		"Password": "6666666666",
	})
	data := bytes.NewBuffer(dataToSend)

	response, err := http.Post("http://localhost:4000/v1/users", "application/json", data)

	if err != nil {
		log.Fatalf("%v", err)
	}

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
	defer response.Body.Close()
}

func TestInvalidUsersCreationEndpoint(t *testing.T) {
	dataToSend, _ := json.Marshal(map[string]string{
		"Name": "",
		"Email":    "@",
		"Password": "",
	})
	data := bytes.NewBuffer(dataToSend)

	response, err := http.Post("http://localhost:4000/v1/users", "application/json", data)

	if err != nil {
		log.Fatalf("%v", err)
	}

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))

	defer response.Body.Close()
}

func TestUserActivation(t *testing.T) {
	client := &http.Client{}
	dataToSend, _ := json.Marshal(map[string]string{
		"Token": "YDRKUQ4HFR5A5XF62OUUANCZM4",
	})
	data := bytes.NewBuffer(dataToSend)

	request, err := http.NewRequest(http.MethodPut, "http://localhost:4000/v1/users/activated", data)
	if err != nil {
		log.Fatalf("An Error Occurred while creating request: %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("An Error Occurred while sending request: %v", err)
	}
	defer response.Body.Close()

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
}

func TestInvalidUserActivation(t *testing.T){
	client := &http.Client{}
	dataToSend, _ := json.Marshal(map[string]string{
		"Token": "YDRKUQ4HFR5A5XF62OUUANCZM5",
	})
	data := bytes.NewBuffer(dataToSend)

	request, err := http.NewRequest(http.MethodPut, "http://localhost:4000/v1/users/activated", data)
	if err != nil {
		log.Fatalf("An Error Occurred while creating request: %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("An Error Occurred while sending request: %v", err)
	}
	defer response.Body.Close()

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
}

func TestGetUserAuthentificationToken(t *testing.T){
	dataToSend, _ := json.Marshal(map[string]string{
		"Email": "RulanA@alimhan.kz",
		"Password": "6666666666",
	})

	data := bytes.NewBuffer(dataToSend)

	response, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", data)
	if err != nil {
		log.Fatalf("%v", err)
	}

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
	defer response.Body.Close()
}

func TestInvalidGetUserAuthentificationToken(t *testing.T){
	dataToSend, _ := json.Marshal(map[string]string{
		"Email": "RulanA@alimhan.kz",
		"Password": "66666666666666666",
	})

	data := bytes.NewBuffer(dataToSend)

	response, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", data)
	if err != nil {
		log.Fatalf("%v", err)
	}

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
	defer response.Body.Close()
}

func TestMovieCreation(t *testing.T){
	client := &http.Client{}
	Movie := Movie{
		Title:   "Godfather",
		Year:    1972,
		Runtime: "170 mins",
		Genres:  []string{"Drama", "Crime"},
	}
	data, err := json.Marshal(Movie)
	if err != nil {
		log.Fatalf("%v", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("%v", err)
	}

	bearerToken := "EVO7EUZQH2ZYOXYHW4LQGDFGVI"
	request.Header.Set("Authorization", "Bearer "+bearerToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer response.Body.Close()

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
	defer response.Body.Close()
}

func TestInvalidYearTitleMovieCreation(t *testing.T){
	client := &http.Client{}
	Movie := Movie{
		Title:   "",
		Year:    30000,
		Runtime: "170 mins",
		Genres:  []string{},
	}
	data, err := json.Marshal(Movie)
	if err != nil {
		log.Fatalf("%v", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("%v", err)
	}

	bearerToken := "EVO7EUZQH2ZYOXYHW4LQGDFGVI"
	request.Header.Set("Authorization", "Bearer "+bearerToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer response.Body.Close()

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
	defer response.Body.Close()
}

func TestInvalidRuntimeMovieCreation(t *testing.T){
	client := &http.Client{}
	Movie := Movie{
		Title:   "Godfather",
		Year:    1972,
		Runtime: "170",
		Genres:  []string{"Drama", "Crime"},
	}
	data, err := json.Marshal(Movie)
	if err != nil {
		log.Fatalf("%v", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("%v", err)
	}

	bearerToken := "EVO7EUZQH2ZYOXYHW4LQGDFGVI"
	request.Header.Set("Authorization", "Bearer "+bearerToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer response.Body.Close()

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
	defer response.Body.Close()
}

func TestPatchMovieEndPoint(t *testing.T){
	client := &http.Client{}
	
	dataToSend, _ := json.Marshal(map[string]int32{
		"year": 1978,
	})

	request, err := http.NewRequest("PATCH", "http://localhost:4000/v1/movies/1", bytes.NewBuffer(dataToSend))
	if err != nil {
		log.Fatalf("%v", err)
	}

	bearerToken := "EVO7EUZQH2ZYOXYHW4LQGDFGVI"
	request.Header.Set("Authorization", "Bearer "+bearerToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer response.Body.Close()

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Print(string(answer))
	defer response.Body.Close()
}