package main

import "fmt"

func main() {
	//接受用户输入的选项
	var key = ""
	var bool = true
	var details string = "收支\t账务金额\t收支金额\t说明\n"
	var money float64
	var balance float64
	var note string

	for {
		fmt.Println("-----------家庭收入记账软件-------------")
		fmt.Println("            1.收支明细")
		fmt.Println("            2.登记收入")
		fmt.Println("            3.登记支出")
		fmt.Println("            4.退出软件")
		fmt.Println("请选择（1-4）:")

		_, _ = fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("------------------当前收支明细--------------")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金额")
			fmt.Scanln(&money)
			if money > 10000 {
				fmt.Println("存储金额单次上限10000,请重新再试")
				break
			}
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			balance += money
			details += fmt.Sprintf("收入\t%v\t%v\t%v\n", balance, money, note)
		case "3":
			fmt.Println("本次支出金额")
			fmt.Scanln(&money)
			if balance-money < 0 {
				fmt.Println("账务余额不足,请重新再试")
				break
			}
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			balance -= money
			details += fmt.Sprintf("支出\t%v\t%v\t%v\n", balance, money, note)
		case "4":
			bool = false
		default:
			fmt.Println("请输入正确选项")
		}
		if !bool {
			break
		}
	}
	fmt.Println("已退出...")
}
