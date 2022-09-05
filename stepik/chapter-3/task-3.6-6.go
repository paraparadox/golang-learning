package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type GroupStat struct {
	Average float64
}

func main() {
	var group Group
	dec := json.NewDecoder(os.Stdin)
	dec.Decode(&group)
	totalScores := 0
	for _, student := range group.Students {
		totalScores += len(student.Rating)
	}
	groupStat := GroupStat{float64(totalScores) / float64(len(group.Students))}
	result, _ := json.MarshalIndent(groupStat, "", "    ")
	fmt.Printf("%s", result)
}

/*

{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            "LastName":"Ien",
            "FirstName":"ccc",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[5,2]
        }
    ]
}

*/
