package main

import "fmt"

func arrayOperation()  {
	
	var hobbies [3]string = [3]string{"drawing", "playing", "watching movies"}
	fmt.Println(hobbies)

	fmt.Println("First element = ", hobbies[0])
	processedHobbies := [2]string{hobbies[1], hobbies[2]}
	fmt.Println(processedHobbies)
	fmt.Println("Remaining element = ", hobbies[1:])

	courses := []string{"golaang masterclass", "rest api with golaang"}
	fmt.Println(courses)
	fmt.Println(len(courses))
	courses = append(courses, "connecting databse with golaang")
	fmt.Println(courses)

	mobile1 := New(1, "Samsung", 80000.0)
	mobile2 := New(2, "Pixel", 60000.0)

	mobiles := []Product{mobile1, mobile2}
	fmt.Println(mobiles)
}