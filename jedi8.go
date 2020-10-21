package main

import (
	"fmt"
	"sort"
)

type user struct { //ninja 8.5 (start)
	First   string
	Last    string
	Age     int
	Sayings []string
}

func displayUser(s string, u []user, b bool) {
	fmt.Println(s)
	for i, person := range u {
		fmt.Println("\tPerson #", i+1)
		fmt.Println("\t\t", person.First, person.Last, person.Age)
		if b {
			sort.Strings(person.Sayings)
		}
		for _, saying := range person.Sayings {
			fmt.Println("\t\t\t", saying)
		}
	}
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByFirstName []user

func (a ByFirstName) Len() int           { return len(a) }
func (a ByFirstName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirstName) Less(i, j int) bool { return a[i].First < a[j].First }

type ByLastName []user

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	showUnsorted := false
	displayUser("UnSorted", users, showUnsorted)
	showUnsorted = true
	sort.Sort(ByAge(users))

	displayUser("Sorted by Age:", users, showUnsorted)

	sort.Sort(ByFirstName(users))
	displayUser("Sorted by First name:", users, showUnsorted)

	sort.Sort(ByLastName(users))
	displayUser("Sorted by Last name:", users, showUnsorted)

} //ninja 8.5 (end)

// func main() { //ninja 8.4 (start)
// 	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
// 	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

// 	fmt.Println(xi)
// 	// sort xi
// 	sort.Ints(xi)
// 	fmt.Println(xi)

// 	fmt.Println(xs)
// 	// sort xs
// 	sort.Strings(xs)
// 	fmt.Println(xs)
// } // ninja 8.4 (end)

// type user struct { //ninja 8.3 (start)
// 	First   string
// 	Last    string
// 	Age     int
// 	Sayings []string
// }

// func main() {
// 	u1 := user{
// 		First: "James",
// 		Last:  "Bond",
// 		Age:   32,
// 		Sayings: []string{
// 			"Shaken, not stirred",
// 			"Youth is no guarantee of innovation",
// 			"In his majesty's royal service",
// 		},
// 	}

// 	u2 := user{
// 		First: "Miss",
// 		Last:  "Moneypenny",
// 		Age:   27,
// 		Sayings: []string{
// 			"James, it is soo good to see you",
// 			"Would you like me to take care of that for you, James?",
// 			"I would really prefer to be a secret agent myself.",
// 		},
// 	}

// 	u3 := user{
// 		First: "M",
// 		Last:  "Hmmmm",
// 		Age:   54,
// 		Sayings: []string{
// 			"Oh, James. You didn't.",
// 			"Dear God, what has James done now?",
// 			"Can someone please tell me where James Bond is?",
// 		},
// 	}

// 	users := []user{u1, u2, u3}

// 	fmt.Println(users, "\n")

// 	err := json.NewEncoder(os.Stdout).Encode(users)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// } //ninja 8.3 (end)

// type person struct {
// 	First   string   `json:"First"`
// 	Last    string   `json:"Last"`
// 	Age     int      `json:"Age"`
// 	Sayings []string `json:"Sayings"`
// }

// func main() {

// 	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
// 	fmt.Println(s)

// 	var people []person
// 	err := json.Unmarshal([]byte(s), &people)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(people)

// for i, person := range people {
// 	fmt.Println("Person #", i+1)
// 	fmt.Println("\t", person.First, person.Last, person.Age)
// 	for _, saying := range person.Sayings {
// 		fmt.Println("\t\t", saying)
// 	}
// }

// } //ninja 8.2 (end)

// type user struct { //ninja 8.1 (start)
// 	First string
// 	Age   int
// }

// func main() {
// 	u1 := user{
// 		First: "James",
// 		Age:   32,
// 	}

// 	u2 := user{
// 		First: "Moneypenny",
// 		Age:   27,
// 	}

// 	u3 := user{
// 		First: "M",
// 		Age:   54,
// 	}

// 	users := []user{u1, u2, u3}

// 	fmt.Println(users)

// 	toJson, err := json.Marshal(users)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(toJson))
// } //ninja 8.1 (end)
