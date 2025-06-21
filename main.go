package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "THIS IS HOME PAGE")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("This is about page"))
}

const PORTNUMBER = ":8080"

func divide(w http.ResponseWriter, r *http.Request) {
	x := 10.0
	y := 0.0

	value, err := divideValues(x, y)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is: %f", x, y, value))
}

func divideValues(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero") // throwing an error, if the value is equal or less then 0
		return 0, err
	}

	// result :=
	return (x / y), nil
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", divide)
	// http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello world")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written : %d", n))
	// })

	fmt.Println(fmt.Sprintf("SERVING ON PORT %s", PORTNUMBER))
	_ = http.ListenAndServe(PORTNUMBER, nil)

}
