package zidane

import (
	"fmt"

	"github.com/lastzidanemagaba/GoRefact/main"
)

func DoSomething() {
	fmt.Println("The time in Jakarta is:", main.Now.Format("15:04"))
}
