package bankoperation

import "fmt"
import "example/App/FileOperation"

func Bank() {

	fmt.Println("Welcome to my bank !!!")

	balance, err := fileoperation.ReadFromFile("result.txt")
	if (err != nil) {
		fmt.Println("Error" , err)
		balance = 1000.0
	}

	for {
		displayOptions()

		option := getOption()

		if (option < 1 || option > 4) {
			fmt.Println("Enter the valid option")
			continue
		}

		switch option {
		case 1 : fmt.Println("Balance = ", balance)
		case 2 : {
				depositAmount := getAmount("deposit")
				if(depositAmount <= 0.0) {
					fmt.Println("Enter the valid amount")
					continue
				}
				balance += depositAmount
				fmt.Println("Amount deposited successfully")
				fileoperation.WriteToFile("result.txt", balance)
			}
		case 3 : {
				withdrawAmount := getAmount("withdraw")
				if (withdrawAmount <= 0.0) {
					fmt.Println("Enter the valid amount")
					continue
				}

				if (withdrawAmount > balance) {
					fmt.Println("Insufficient balance")
					continue
				}
				balance -= withdrawAmount
				fmt.Println("Amount withdrawn successfully")
				fileoperation.WriteToFile("result.txt", balance)
			}
		case 4 : fmt.Println("Thanks for visiting our bank") 
				 return
		}
	}

}

func getOption() int{

	var option int;
	fmt.Print("Enter the option = ")
	fmt.Scan(&option)

	return option
}

func getAmount(operation string) float64{

	var amount float64;
	fmt.Println("Enter the amount for " + operation + " = ")
	fmt.Scan(&amount)

	return amount
}