package src

import "fmt"

type FamilyAccount struct {
	key     string
	bool    bool
	balance float64
	money   float64
	note    string
	details string
}

func NewFalilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		bool:    true,
		balance: 0,
		money:   0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说明\n",
	}
}

func (this *FamilyAccount) ShowMainMenu() {
	for {
		fmt.Println("-----------家庭收入记账软件-------------")
		fmt.Println("            1.收支明细")
		fmt.Println("            2.登记收入")
		fmt.Println("            3.登记支出")
		fmt.Println("            4.退出软件")
		fmt.Println("请选择（1-4）:")

		_, _ = fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.showIncome()
		case "3":
			this.showOutput()
		case "4":
			this.bool = false
		default:
			fmt.Println("请输入正确选项")
		}
		if !this.bool {
			break
		}
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("------------------当前收支明细--------------")
	fmt.Println(this.details)
}

func (this *FamilyAccount) showIncome() {
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.money)
	if this.money > 10000 {
		fmt.Println("存储金额单次上限10000,请重新再试")
		return
	}
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.balance += this.money
	this.details += fmt.Sprintf("收入\t%v\t%v\t%v\n", this.balance, this.money, this.note)
}

func (this *FamilyAccount) showOutput() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	if this.balance-this.money < 0 {
		fmt.Println("账务余额不足,请重新再试")
		return
	}
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.balance -= this.money
	this.details += fmt.Sprintf("支出\t%v\t%v\t%v\n", this.balance, this.money, this.note)
}
