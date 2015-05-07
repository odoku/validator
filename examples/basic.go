package main

import (
	"fmt"
	"net/http"

	"github.com/odoku/validator"
)

type Book struct {
	Title         string `form:"title"          valid:"required(), max_length(50)"`
	Discount      bool   `form:"discount"       valid:"required(), bool()"`
	Price         int    `form:"price"          valid:"if(discount, false, required()), int(), min(0)"`
	DiscountPrice int    `form:"discount_price" valid:"if(discount, true,  required()), int(), min(0)"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	v := validator.NewWith(Book{})

	r.ParseForm()
	result, err := v.Validate(r.Form)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	if !result.IsValid {
		fmt.Fprintf(w, "Validation failed.\n")
		for field, errors := range result.Errors {
			fmt.Fprintf(w, "%s: %v\n", field, errors)
		}
		return
	}

	fmt.Fprintf(w, "Validation succeeded.")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
