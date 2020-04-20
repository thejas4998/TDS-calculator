package main

import(
	"fmt"
)
func add()float64{
	var salary float64
	fmt.Print("Enter your salary before TDS:- ")
	if _, err:=fmt.Scan(&salary);err!=nil{
		fmt.Println("Bro enter proper salary bro!")
		salary=add()
	}
	
	return salary
}

func select_tax_bracket(sal float64)rune{
	if sal<0 {
		fmt.Println("You've entered a negative number. Program aborting!")
	}

	if sal<=500000 && sal>0{
		return 'a'
	}

	if sal>500000 && sal<=750000{
		return 'b'
	}

	if sal>750000 && sal<=1000000{
		return 'c'
	}

	if sal>1000000 && sal<=1250000{
		return 'd'
	}

	if sal>1250000 && sal<=1500000{
		return 'e'
	}

	if sal>1500000{
		return 'f'
	}

	return 'z'
}
func main(){
	
	for{

	
	salary:=add()
	bracket:=select_tax_bracket(salary)
	
	switch bracket {
	case 'a':
		fmt.Printf("Your salary after TDS is : %9.2f", salary)
	case 'b':
		salary -= (0.1*(salary-500000))
		fmt.Printf("Your salary after TDS is : %9.2f", salary)
	case 'c':
		salary -= 25000
		salary -= (0.15*(salary-750000))
		fmt.Printf("Your salary after TDS is : %9.2f", salary)
	case 'd':
		salary -= 62000
		salary -= (0.2*(salary-1000000))
		fmt.Printf("Your salary after TDS is : %9.2f", salary)
	case 'e':
		salary -= 112000
		salary -= (0.25*(salary-1250000))
		fmt.Printf("Your salary after TDS is : %9.2f", salary)
	case 'f':
		salary -= 174000
		salary -= (0.3*(salary-1500000))
		fmt.Printf("Your salary after TDS is : %9.2f", salary)
	case 'z':
		fmt.Println("Something went terribly wrong")	
	default:
		fmt.Println("Something went terribly wrong")

	}
	fmt.Println(" ")
	fmt.Printf("Your salary per month is : %9.2f", salary/12)
	fmt.Println(" ")
}

}
