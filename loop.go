package main
import ("fmt")
 

func main(){
	/*
	for statement1;statement2;statement3{
		code to be executed for each iteration
	}
	*/
	//i++=>i=i+1
	// for i:=0;i<5;i++{
	// 	fmt.Println(i)
	// }
	//i+=10=>i=i+10
	// for i:=0;i<=100;i+=10{
	// 	fmt.Println(i)
	// }
	// for i:=0;i<=5;i++{
	// 	if i ==3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// for i:=0;i<=5;i++{
	// 	if i ==3{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// adj:=[2]string {"big","tasty"}
	// fruits:=[3]string {"apple","orange","banana"}
	// for i:=0;i<len(adj);i++{
	// 	for j:=0;j<len(fruits);j++{
	// 		fmt.Println(adj[i],fruits[j])
	// 	}
	// }

	/*
	for index,value :=range array|slice|map{
		code to be executed for each iteration
	}
	*/
	fruits:=[3]string {"apple","orange","banana"}
	for index,_:=range fruits{
		fmt.Println(index)
	}
	}
