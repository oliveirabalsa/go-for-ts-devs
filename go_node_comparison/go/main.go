package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/fibonacci/") {
		numStr := strings.TrimPrefix(r.URL.Path, "/fibonacci/")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		resultChan := make(chan int)

		go func() {
			resultChan <- fibonacci(num)
		}()

		result := <-resultChan

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"result": result})
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}
