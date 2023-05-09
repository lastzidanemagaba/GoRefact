package app

import (
	"fmt"
	"time"

	"github.com/lastzidanemagaba/GoRefact/main"
	"github.com/lastzidanemagaba/GoRefact/zidane"
)

func Run() {
	fmt.Println("The current time is:", main.Now.Format(time.RFC3339))
	zidane.DoSomething()
}
