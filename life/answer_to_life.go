package main

import "fmt"

var WhatIsThe int = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe = 0
}

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	} else {
		fmt.Println("It's all true???")
	}
}
