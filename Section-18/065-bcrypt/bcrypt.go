package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"os"
)

// JSONBody struct for the JSON structure of a request body
type JSONBody struct {
	Password       string `json:"password"`
	HashedPassword string `json:"hashedPassword"`
	Cost           int    `json:"cost"`
}

// JSONResponse struct for the structure of basic JSON responses through HTTP
type JSONResponse struct {
	Status         string `json:"status,omitempty"`
	Message        string `json:"message,omitempty"`
	HashedPassword string `json:"hashedPassword,omitempty"`
	PasswordMatch  bool   `json:"passwordMatch"`
	Cost           int    `json:"cost,omitempty"`
}

// JSONError struct for the structure of basic JSON errors through HTTP
type JSONError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", routeHome)
	http.HandleFunc("/bcrypt", routeBcrypt)
	http.HandleFunc("/compare", routeCompare)
	http.HandleFunc("/cost", routeCost)

	http.ListenAndServe(":8080", nil)
}

func routeHome(w http.ResponseWriter, req *http.Request) {
	file, _ := os.Open("./home.html")
	defer file.Close()

	response := make([]byte, 10240)

	n, _ := file.Read(response)

	response = response[:n]

	w.Header().Set("Content-Type", "text/html")
	w.Write(response)
}

func routeBcrypt(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		errorResponse(w, req, http.StatusNotFound, "HTTP method for this resource not available.")
	} else if req.Method == "POST" {
		bsBody := make([]byte, 10240)
		n, err := req.Body.Read(bsBody)

		if err != nil && err != io.EOF {
			errorResponse(w, req, http.StatusBadRequest, "Invalid request body.")
			return
		}

		bsBody = bsBody[:n]

		body := JSONBody{}

		err = json.Unmarshal(bsBody, &body)

		if err != nil {
			errorResponse(w, req, http.StatusBadRequest, "Invalid request body.")
			return
		}

		if body.Cost == 0 {
			body.Cost = 12
		}

		bsHashedPassword, err := hashPassword(body.Password, body.Cost)
		if err != nil {
			serverErrorResponse(w, req)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(JSONResponse{
			Status:         "success",
			HashedPassword: string(bsHashedPassword),
		})

		w.Write(response)
	}
}

func hashPassword(password string, cost int) ([]byte, error) {
	bsPassword := []byte(password)

	bsHashedPassword, err := bcrypt.GenerateFromPassword(bsPassword, cost)
	if err != nil {
		return []byte{}, err
	}

	return bsHashedPassword, nil
}

func routeCompare(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		errorResponse(w, req, http.StatusNotFound, "HTTP method for this resource not available.")
	} else if req.Method == "POST" {
		bsBody := make([]byte, 10240)
		n, err := req.Body.Read(bsBody)

		if err != nil && err != io.EOF {
			errorResponse(w, req, http.StatusBadRequest, "Invalid request body.")
			return
		}

		bsBody = bsBody[:n]

		body := JSONBody{}

		err = json.Unmarshal(bsBody, &body)

		if err != nil {
			errorResponse(w, req, http.StatusBadRequest, "Invalid request body.")
			return
		}

		bsPassword, bsHashedPassword := []byte(body.Password), []byte(body.HashedPassword)

		fail := bcrypt.CompareHashAndPassword(bsHashedPassword, bsPassword)

		if fail != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			response, _ := json.Marshal(JSONResponse{
				Status:        "success",
				PasswordMatch: false,
			})

			w.Write(response)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			response, _ := json.Marshal(JSONResponse{
				Status:        "success",
				PasswordMatch: true,
			})

			w.Write(response)
		}
	}
}

func routeCost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		errorResponse(w, req, http.StatusNotFound, "HTTP method for this resource not available.")
	} else if req.Method == "POST" {
		bsBody := make([]byte, 10240)
		n, err := req.Body.Read(bsBody)

		if err != nil && err != io.EOF {
			errorResponse(w, req, http.StatusBadRequest, "Invalid request body.")
			return
		}

		bsBody = bsBody[:n]

		body := JSONBody{}

		err = json.Unmarshal(bsBody, &body)

		if err != nil {
			errorResponse(w, req, http.StatusBadRequest, "Invalid request body.")
			return
		}

		cost, err := bcrypt.Cost([]byte(body.HashedPassword))
		if err != nil {
			serverErrorResponse(w, req)
			return
		}
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(JSONResponse{
			Status: "success",
			Cost:   cost,
		})

		w.Write(response)
	}
}

func errorResponse(w http.ResponseWriter, req *http.Request, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response, _ := json.Marshal(JSONError{
		Status:  "fail",
		Message: message,
	})

	w.Write(response)
}

func serverErrorResponse(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	response, _ := json.Marshal(JSONError{
		Status:  "error",
		Message: "Something went wrong... Please try again later.",
	})

	w.Write(response)
}
