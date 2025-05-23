package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ResponseCall struct {
	Status int
	Time   time.Duration
	Body   string
}

// Truncate corta el string a la longitud especificada, manejando casos cortos
func (r ResponseCall) Truncate(length int) string {
	if len(r.Body) <= length {
		return r.Body
	}
	return r.Body[:length]
}

// UserCapture captura y valida la URL del usuario
func UserCapture() (string, error) {
	fmt.Print("Inserte la URL a consultar (ej. https://example.com): ")
	var uri string
	_, err := fmt.Scanln(&uri)
	if err != nil {
		return "", fmt.Errorf("error al leer la URL: %v", err)
	}

	// Asegurarse de que la URL tenga un esquema (http o https)
	if !strings.HasPrefix(uri, "http://") && !strings.HasPrefix(uri, "https://") {
		uri = "https://" + uri
	}

	// Validar el formato de la URL
	_, err = url.ParseRequestURI(uri)
	if err != nil {
		return "", fmt.Errorf("URL inválida: %v", err)
	}

	return uri, nil
}

// ExecuteCall realiza la solicitud HTTP y devuelve la respuesta
func ExecuteCall(uri string) (ResponseCall, error) {
	start := time.Now()
	response, err := http.Get(uri)
	if err != nil {
		return ResponseCall{}, fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer response.Body.Close()

	elapsed := time.Since(start)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseCall{
			Status: response.StatusCode,
			Time:   elapsed,
			Body:   "",
		}, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	return ResponseCall{
		Status: response.StatusCode,
		Time:   elapsed,
		Body:   string(body),
	}, nil
}

func main() {
	fmt.Println("Programa para hacer llamadas HTTP GET")

	// Capturar y validar la URL
	uri, err := UserCapture()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ejecutar la solicitud HTTP
	response, err := ExecuteCall(uri)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Mostrar resultados
	fmt.Printf("\nResultados de la solicitud a %s:\n", uri)
	fmt.Printf("Código de estado: %d\n", response.Status)
	fmt.Printf("Tiempo de respuesta: %s\n", response.Time)
	fmt.Printf("Primeros 100 caracteres del cuerpo:\n%s\n", response.Truncate(100))
}
