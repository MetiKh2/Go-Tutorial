package main
import ("fmt")
 

func main(){
 /*
	switch expression{
	case x:
		code block
	case y:
		code block
	case z:
		code block
	...

	default:
		code block
	}
 */

//  day :=8

//  switch day{
// case 1:
// 	fmt.Println("Monday")
// case 2:
// 	fmt.Println("Tuesday")
// case 3:
// 	fmt.Println("Wednesday")
// case 4:
// 	fmt.Println("Thursday")
// case 5:
// 	fmt.Println("Friday")
// case 6:
// 	fmt.Println("Saturday")
// case 7:
// 	fmt.Println("Sunday")
// default:
// 	fmt.Println("Not a weekday")
//  }

 /*
	switch expression{
	case x,y:
		code block
	case w,z:
		code block
	case a:
		code block
	...

	default:
		code block
	}
 */
	day :=6

 
 switch day{
case 1,3,5:
	fmt.Println("Odd Weekday")
case 2,4:
	fmt.Println("Even Weekday")
case 6,7:
	fmt.Println("Weekend")
 
default:
	fmt.Println("Not a weekday")
 }
	}
