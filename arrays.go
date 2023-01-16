package main
import ("fmt")
 

func main(){
	//var array_name = [length] datatype{values}
	//length is defined 

	//var array_name = [...] datatype{values}
	//length is inferred 

	//array_name:=[length]datatype{values}
	//length is defined 

	//array_name:=[...]datatype{values}
	//length is inferred  


	// var arr1=[3] int {1,2,3}
	// arr2:=[5]int {4,5,6,7,8}

	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// var arr1=[...] int {1,2,3}
	// arr2:=[...]int {4,5,6,7,8}

	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// var cars=[4] string {"Volvo","BMW","Ford","Mazda"}

	// fmt.Println(cars)

	// prices:=[3] int {10,20,30}
	// fmt.Println(prices[0])
	// fmt.Println(prices[1])
	// fmt.Println(prices[2])
	// prices[1]=50
	// fmt.Println("__________________________________")
	// fmt.Println(prices)

	// arr1:=[5]int {}//not initialized
	// arr2:=[5]int{1,2}//partially initialized
	// arr3:=[5]int{1,2,3,4,5}//fully initialized
	// fmt.Println(arr1)
	// fmt.Println(arr2)
	// fmt.Println(arr3)

	// arr1:=[5]int{1:10,2:40}
	// fmt.Println(arr1)

	var arr1=[4] string {"Volvo","BMW","Ford","Mazda"}
	var arr2=[...] int {1,2,3,4,5,6}
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))

}
