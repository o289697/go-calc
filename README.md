calc := &calc.Calc{}


result, _ := calc.Calc("0.1*50+0.2+0.3/0.2")


fmt.Println("计算结果:", result)

result, _ = calc.Calc("(10.5*2)+(21.5*4)")


fmt.Println("计算结果:", result)
