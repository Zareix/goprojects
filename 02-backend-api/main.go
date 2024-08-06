package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"zareix/goprojects/02-backend-api/internal/middleware"

	"github.com/rs/cors"
)

type POSTTwoNumbers struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type POSTDivision struct {
	Dividend float32 `json:"dividend"`
	Divisor  float32 `json:"divisor"`
}

func readStructFromJSONBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.Unmarshal(body, v)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		var input POSTTwoNumbers
		err := readStructFromJSONBody(r, &input)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res := input.Number1 + input.Number2

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})
	router.HandleFunc("POST /subtract", func(w http.ResponseWriter, r *http.Request) {
		var input POSTTwoNumbers
		err := readStructFromJSONBody(r, &input)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res := input.Number1 - input.Number2

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})
	router.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		var input POSTTwoNumbers
		err := readStructFromJSONBody(r, &input)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res := input.Number1 * input.Number2

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})
	router.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		var input POSTDivision
		err := readStructFromJSONBody(r, &input)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if input.Divisor == 0 {
			http.Error(w, "Divisor cannot be zero", http.StatusBadRequest)
			return
		}

		res := input.Dividend / input.Divisor

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})
	router.HandleFunc("POST /sum", func(w http.ResponseWriter, r *http.Request) {
		var input []int
		err := readStructFromJSONBody(r, &input)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var res int = 0
		for _, i := range input {
			res += i
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	server := http.Server{
		Addr: ":3000",
		// Handler: cors.Default().Handler(middleware.LoggingMiddleware(router)),
		Handler: middleware.CreateStack(
			cors.Default().Handler,
			middleware.LoggingMiddleware,
		)(router),
	}

	fmt.Println("Listening on port 3000...")
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("An error occurred:", err)
	}
}
