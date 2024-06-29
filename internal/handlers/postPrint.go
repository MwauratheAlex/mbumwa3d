package handlers

import (
	"fmt"
	"net/http"
)

func PostPrint(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		fmt.Println("Unable to pass form data")
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("No file uploaded")
	} else {
		fmt.Println("file uploaded")
	}

	technology := r.FormValue("technology")
	color := r.FormValue("color")
	buildTime := r.FormValue("time")
	quantity := r.FormValue("quantity")
	price := r.FormValue("price")

	fmt.Println("tech: ", technology, "Color: ", color, "time: ",
		buildTime, "qty: ", quantity, "price: ", price)

	fmt.Println("printing 2", file)

	return nil
}
