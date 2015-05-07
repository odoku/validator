package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/odoku/validator"
)

type Book struct {
	Title string `form:"title" valid:"required(), max_length(50)"`
	Price int    `form:"price" valid:"if(discount, false, required()), int(), min(0)"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root@/test?charset=utf8")
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	f := validator.NewFunction(func(value string) error {
		var count int
		query := db.Where("title = ?", value).Find(&Book{}).Count(&count)
		if query.Error == gorm.RecordNotFound {
			return nil
		} else if query.Error != nil {
			return query.Error
		}

		message := fmt.Sprintf("Title '%v' is already exists.", value)
		return &validator.ValidationError{Message: message}
	})

	v := validator.NewWith(Book{})
	v.AddRule("title", f)

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
