package main

import (
	"context"
	"net/http"
	"time"
)

// Context digunakan untuk membatalkan suatu proses apabila proses tersebut sudah melebihi batas waktu terentu
// Context digunakan untuk mengatur timeout pada suatu proses

func main() {
	toContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(toContext, http.MethodGet, "https://www.google.com", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	res.Body.Close()

}
