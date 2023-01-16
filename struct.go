package main
import ("fmt")
type Person struct {
	name string
	age int
	job string
	salary int
}

func main(){
	// type struct_name struct {
	// 	memeber1 datatype;
	// 	memeber2 datatype;
	// 	memeber3 datatype;
	// }

	
	var pers1 Person 
	var pers2 Person
	
	pers1.name = "Mahdi"
	pers1.age=17
	pers1.job = "Programmer"
	pers1.salary =6000

	pers2.name = "Cecilie"
	pers2.age=20
	pers2.job = "Teacher"
	pers2.salary =10000

	 printPerson(pers1)
	 fmt.Println("____________________")
	 printPerson(pers2)
	

	}

func printPerson(pers Person){
	fmt.Println("Name" ,pers.name)
	fmt.Println("Age",pers.age)
	fmt.Println("Job",pers.job)
	fmt.Println("Salary",pers.salary)
} 
