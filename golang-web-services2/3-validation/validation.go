package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
	"net/http"
)

// http://localhost:8080/?to=v.romanov@corp&priority=low&subject=Hello!&inner=ignored&id=12&flag=23

type SendMessage struct {
	Id        int    `valid:"optional"`
	Priority  string `valid:"in(low|normal|high)"`
	Recipient string `schema:"to" valid:"email"`
	Subject   string `valid:"msgSubject"`
	Inner     string `chema:"-" valid:"-"`
	flag      int
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("request " + r.URL.String() + "\n\n"))

	msg := &SendMessage{}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err := decoder.Decode(msg, r.URL.Query())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal", 500)
		return
	}
	w.Write([]byte(fmt.Sprintf("Msg: %#v\n\n", msg)))

	_, err = govalidator.ValidateStruct(msg)

	if err != nil {
		if allErrs, ok := err.(govalidator.Errors); ok {
			for _, fld := range allErrs.Errors() {
				data := []byte(fmt.Sprintf("field: %#v\n\n", fld))
				w.Write(data)
			}
		}

		w.Write([]byte(fmt.Sprintf("error: %s\n\n", err)))
	} else {
		w.Write([]byte(fmt.Sprintf("msg is correct\n\n")))
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("starting server at :4011")
	http.ListenAndServe(":4011", nil)
}

func init() {
	govalidator.CustomTypeTagMap.Set("msgSubject", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		if len(subject) == 0 || len(subject) > 10 {
			return false
		}
		return true
	}))
}
