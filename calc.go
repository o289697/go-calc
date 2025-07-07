package goCalc

import (
	"fmt"
	"strings"
	"text/scanner"

	"github.com/shopspring/decimal"
)

type Calc struct {
}

func (c Calc) Calc(s string) (float64, error) {
	postfix, _ := c.infixToPostfix(s)
	result, _ := c.evaluatePostfix(postfix)
	return result, nil
}

// 判断是否是操作符
func (c Calc) isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// 获取操作符优先级
func (c Calc) precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

// 中缀表达式转后缀表达式（逆波兰表达式）
func (c Calc) infixToPostfix(expression string) ([]string, error) {
	var output []string
	var stack []string

	var s scanner.Scanner
	src := strings.NewReader(strings.ReplaceAll(expression, " ", ""))
	s.Init(src)

	for {
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}
		token := s.TokenText()
		switch {
		case token == "(":
			stack = append(stack, token)
		case token == ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1] // 弹出左括号
			} else {
				return nil, fmt.Errorf("括号不匹配")
			}
		case c.isOperator(token):
			for len(stack) > 0 && c.precedence(stack[len(stack)-1]) >= c.precedence(token) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		default:
			// 数字或变量
			output = append(output, token)
		}
	}
	// 将栈中剩余操作符加入输出
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top == "(" || top == ")" {
			return nil, fmt.Errorf("括号不匹配")
		}
		output = append(output, top)
	}
	return output, nil
}

// 计算后缀表达式的值
func (c Calc) evaluatePostfix(postfix []string) (float64, error) {
	var stack []decimal.Decimal

	for _, token := range postfix {
		if c.isOperator(token) && len(stack) >= 2 {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result decimal.Decimal
			switch token {
			case "+":
				result = a.Add(b)
			case "-":
				result = a.Sub(b)
			case "*":
				result = a.Mul(b)
			case "/":
				if b.IsZero() {
					return 0, fmt.Errorf("除数不能为零")
				}
				result = a.Div(b)
			}
			stack = append(stack, result)
		} else {
			stack = append(stack, decimal.RequireFromString(token))
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("表达式不合法")
	}
	f, _ := stack[0].Float64()
	return f, nil
}
