package main
import ("fmt")
 

func main(){
	 //[] datatype {values} format
		 //slice_name:=[]datatype {values}
	
		//  myslice:=[]string{"Go","Slices","Are","Powerful"}
		//  fmt.Println(myslice)
		//  fmt.Println(cap(myslice))
		//  fmt.Println(len(myslice))



	 //Create slice from an array
		//var arr=[length]datatype {values}
		//slice:=arr[start:end]

		// arr:=[6]int {10,11,12,13,14,15}
		// slice:=arr[0:4]

		// fmt.Println(slice)
		// fmt.Println("length : ",len(slice))
		// fmt.Println("capacity : ",cap(slice))

	 //Using the make() function
	 	//slice:=make([]type,length,capacity)

		// slice:=make([]int,5,10)
		// fmt.Println(slice)
		// fmt.Println("length : ",len(slice))
		// fmt.Println("capacity : ",cap(slice))

		// slice2:=make([]int,5)
		// fmt.Println(slice2)
		// fmt.Println("length : ",len(slice2))
		// fmt.Println("capacity : ",cap(slice2))

		// prices:=[]int{10,20,30}
		// prices[2]=50
		// fmt.Println(prices[0])
		// fmt.Println(prices[2])
		// fmt.Println(prices)
	//slice=append(slice,element1,element2,...)
		// slice1:=[]int{1,2,3,4,5,6}
		// fmt.Println(slice1)
		// fmt.Println("length : ",len(slice1))
		// fmt.Println("capacity : ",cap(slice1))

		// slice1=append(slice1,20,21)
		// fmt.Println(slice1)
		// fmt.Println("length : ",len(slice1))
		// fmt.Println("capacity : ",cap(slice1))

	//slice3=append(slice1,slice2...)
		// slice1:=[]int{1,2,3}
		// slice2:=[]int{4,5,6}
		// slice3:=append(slice1,slice2...)
		// fmt.Println(slice3)
		// fmt.Println("length : ",len(slice3))
		// fmt.Println("capacity : ",cap(slice3))

		// arr1 := [6]int{9, 10, 11, 12, 13, 14}
		// slice1:=arr1[1:5]
		// fmt.Println(slice1)
		// fmt.Println("length : ",len(slice1))
		// fmt.Println("capacity : ",cap(slice1))
		// fmt.Println("______________________")

		// slice1=arr1[1:3]//change length by re-slicing the array
		// fmt.Println(slice1)
		// fmt.Println("length : ",len(slice1))
		// fmt.Println("capacity : ",cap(slice1))

		// fmt.Println("______________________")

		// slice1=append(slice1,20,21,22,23)//change length by appending items 
		// fmt.Println(slice1)
		// fmt.Println("length : ",len(slice1))
		// fmt.Println("capacity : ",cap(slice1))

	//copy(dest,src)
		numbers:=[]int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
		//Original Slice
		fmt.Println(numbers)
		fmt.Println("length : ",len(numbers))
		fmt.Println("capacity : ",cap(numbers))

		//Create Copy With Only Needed Numbers
		neededNumbers:=numbers[:len(numbers)-10]
		numbersCopy:=make([]int,len(neededNumbers))
		copy(numbersCopy,neededNumbers)
		fmt.Println("_____________________________")
		fmt.Println(numbersCopy)
		fmt.Println("length : ",len(numbersCopy))
		fmt.Println("capacity : ",cap(numbersCopy))

	}
