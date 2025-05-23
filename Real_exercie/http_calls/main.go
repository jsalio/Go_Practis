package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type StringBody string

func (s StringBody) Truncate(length int) string {
	return string(s[:length])
}

type ResponseCall struct {
	Status int
	Time   time.Duration
	Body   StringBody
}

func UserCapture() string {
	var uri string = ""
	fmt.Print("inserte la url a consultar :")
	_, err := fmt.Scanln(&uri)

	if err != nil {
		fmt.Errorf("Invalid url")
		return ""
	}
	return uri
}

func ExecuteCall(uri string) ResponseCall {
	start := time.Now()
	response, error := http.Get(uri)

	if error != nil {
		elapsed := time.Since(start)
		return ResponseCall{
			Status: 500,
			Time:   elapsed,
			Body:   "",
		}
	}
	defer response.Body.Close()
	elapsed := time.Since(start)

	if response.StatusCode != 200 {
		return ResponseCall{
			Status: response.StatusCode,
			Time:   elapsed,
			Body:   "",
		}
	}

	content, error := io.ReadAll(response.Body)

	if error != nil {
		return ResponseCall{
			Status: response.StatusCode,
			Time:   elapsed,
			Body:   "Response can't decode",
		}
	}

	return ResponseCall{
		Status: response.StatusCode,
		Time:   elapsed,
		Body:   StringBody(content),
	}
}

func main() {
	fmt.Println("Hacer lla,ada a apis")
	uri := UserCapture()
	response := ExecuteCall(uri)
	fmt.Printf("La llamada a %s respondio con %d duro %s y su contenido es \n %s", uri, response.Status, response.Time, response.Body.Truncate(25))
}
