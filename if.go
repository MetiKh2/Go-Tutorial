package main
import ("fmt")
 

func main(){
	  //<
	  //<=
	  //>
	  //>=
	  //==
	  //!=
	  //&&=> AND
	 //|| =. OR
	 //! =>Not

	 //if
	 //else 
	 //else if 
	 //switch 

	 //if condition {
		//code to be executed if condition is true 
	//}

	x:=18
	y:=20
	if y<x{
		fmt.Printf("20 is greater than 18")
	}

	/*
	if condition{
		code to be executed if condition is true 
	}
	else{
		code to be executed if condition is false 
	}
	*/

	// time:=10
	// if (time<18){
	// 	fmt.Printf("Good day")
	// }else{
	// 	fmt.Printf("Good evening")
	// }
	/*
	if condition{
		code to be executed if condition1 is true 
	}else if condition2{
		code to be executed if condition1 is false and
		condition2 is true
	}
	else{
		code to be executed if condition1 and 
		condition2 are both false 
	}
	*/
	// time:=9
	// if (time<10){
	// 	fmt.Printf("Good morning")
	// }else if time<20{
	// 	fmt.Printf("Good day")
	// }else{
	// 	fmt.Printf("Good evening")
	// }

	/*
	if condition1{
		  code to be executed if condition1 is true
		if condition2{
			code to be executed if both
			 condition1 and condition2 are true
		}
	}
	*/

	num:=12
	if num>=10{
		fmt.Println("Num is more than 10 ")
		if num>15{
		fmt.Println("Num is also more than 15 ")
		}
	}else{
		fmt.Println("Num is less than 10 ")
	}
	}
