package main

import "fmt"

type TechDets struct {
	tech string
	exp  float64
}

type Tech struct {
	id       int
	techdets []TechDets
}

type ContactDets struct {
	email string
	phone string
}

type Contact struct {
	id int
	ContactDets
}

type Address struct {
	area    string
	country string
}

type User struct {
	id   int
	name string
	Address
}

type MergeUser struct {
	id   int
	name string
	Address
	techdets []TechDets
	email    string
	phone    string
}

func addContactPrefix(users []User, techDetails []Tech, contactDetails []Contact) map[int]MergeUser {
	countrymap := map[string]string{
		"IND": "+91",
		"UK":  "+41",
	}
	usersMap := make(map[int]User)
	for _, user := range users {
		usersMap[user.id] = user
	}
	merged := make([]MergeUser, len(users))
	for i, user := range users {

		mergeuser := MergeUser{}

		if valuser, ok := usersMap[user.id]; ok {
			mergeuser.id = valuser.id
			mergeuser.name = valuser.name
			mergeuser.Address = valuser.Address
		}

		for _, tech := range techDetails {

			if user.id == tech.id {
				mergeuser.techdets = tech.techdets
			}

		}

		for _, contact := range contactDetails {

			if user.id == contact.id {
				mergeuser.email = contact.email
				if valprefix, ok := countrymap[mergeuser.Address.country]; ok {
					mergeuser.phone = valprefix + " " + contact.phone
				}

			}

		}

		merged[i] = mergeuser
	}
	mapuser := make(map[int]MergeUser)
	for _, val := range merged {
		mapuser[val.id] = val
	}
	return mapuser
}

func display(alluser map[int]MergeUser) {
	for key, val := range alluser {
		fmt.Println(key, val)
	}
}

func main() {

	user1 := User{1, "Radha", Address{"Bopal", "IND"}}
	tech1 := Tech{1, []TechDets{{"react", 3}, {"Golang", 1.5}}}
	contact1 := Contact{1, ContactDets{"radha.kotech@bacancy.com", "12345 7890"}}

	user2 := User{2, "Aniket", Address{"Maninagar", "UK"}}
	tech2 := Tech{2, []TechDets{{"Vue", 0.9}, {"Golang", 1.5}}}
	contact2 := Contact{2, ContactDets{"aniket.amin@bacancy.com", "09876 54321"}}

	user3 := User{3, "Jaimin", Address{"Rajkot", "IND"}}
	tech3 := Tech{3, []TechDets{{"Golang", 0.2}}}
	contact3 := Contact{3, ContactDets{"jaimin.parmar@bacancy.com", "96628 73324"}}

	user4 := User{4, "Anshuman", Address{"Baroda", "UK"}}
	tech4 := Tech{4, []TechDets{{"Golang", 0.3}}}
	contact4 := Contact{4, ContactDets{"anshuman.aich@bacancy.com", "99885 22335"}}

	users := []User{user1, user2, user3, user4}
	techDetails := []Tech{tech1, tech2, tech3, tech4}
	contactDetails := []Contact{contact1, contact2, contact3, contact4}

	mergealluser := addContactPrefix(users, techDetails, contactDetails)
	fmt.Println(mergealluser)

}
