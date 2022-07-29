package people

func populateDatabase() {
	people = append(people, Person{Id: "1", FirstName: "John", LastName: "Doe", Addres: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{Id: "2", FirstName: "Petter", LastName: "Tomson"})
}
