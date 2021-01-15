package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"text/template"
)

type User struct {
	Userid       int
	Firstname    string
	Lastname     string
	DataofBirth  string
	Email        string
	MobileNumber string
}

func display(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.html")
		t.Execute(w, nil)
	}
}

func formhandling(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		userformdetails := User{
			Userid:       rand.Intn(10000000),
			Firstname:    r.FormValue("firstname"),
			Lastname:     r.FormValue("lastname"),
			DataofBirth:  r.FormValue("dateofbirth"),
			Email:        r.FormValue("email"),
			MobileNumber: r.FormValue("mobilenumber"),
		}
		// _, err := json.Marshal(details)
		// if err != nil {
		// 	fmt.Println("Error in marshal")
		// }

		usersfiledata, err := ioutil.ReadFile("user.json")
		if err != nil {
			fmt.Println(err)
		}

		var alluserdetails []User
		err = json.Unmarshal([]byte(usersfiledata), &alluserdetails)
		if err != nil {
			fmt.Println("Error JSON Unmarshling for user file")
			fmt.Println(err)

		}

		alluserdetails = append(alluserdetails, userformdetails)
		file, _ := json.MarshalIndent(alluserdetails, "", " ")
		_ = ioutil.WriteFile("user.json", file, 0644)
		//w.Write(file)
		t, _ := template.ParseFiles("success.html")
		t.Execute(w, nil)

	}
}

func main() {
	fmt.Println("User Form Handling")
	http.HandleFunc("/", display)
	http.HandleFunc("/dataprocess", formhandling)
	fmt.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)

}
