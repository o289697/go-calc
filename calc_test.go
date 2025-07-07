package calc

import (
	"testing"
)

//go test -v -run TestCalcAll$

// 表格驱动测试
func TestCalcAll(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  float64
	}{
		{"case 0", "0.1+0.2", 0.3},
		{"case 1", "1 + 2 * (3 - 4)", -1},
		{"case 2", "(10.5*10)+(55.5*10)", 660},
		{"case 3", "(10.5*10)+(55.5*10)+40", 700},
		{"case 4", "(11.5*10)+(3.5*10)", 150},
		{"case 5", "0.1+0.2+0.3", 0.6},
		{"case 6", "0.3/0.2", 1.5},
		{"case 7", "0.1+0.2+0.3/0.2", 1.8},
	}
	calc := &Calc{}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := calc.Calc(tt.input)
			if got != tt.want {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}
