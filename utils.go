package go_cc

import (
	"strconv"
)

func Float64(str string) (num float64) {
	num, _ = strconv.ParseFloat(str, 64)
	return
}

func Int(str string) (num int) {
	num, _ = strconv.Atoi(str)
	return
}

func hcf(a, b int) int {
	if b == 0 {
		return a
	} else {
		return hcf(b, a%b)
	}
}

func lcm(a, b int) int {
	return a * b / hcf(a, b)
}

/*
expectations

If(condition, value_or_callback).
	Elif(elseCond, value_or_callback).
	Elif(elseCond, value_or_callback).
	Else(value_or_callback)
*/
type Iffy struct {
	matched bool
	result  interface{}
}

func If(condition bool, then...interface{}) Iffy {
	iffy := Iffy{matched: condition}
  if len(then) == 1 && condition{
    iffy.result = handleExpr(then[0])
  }
	return iffy
}

func handleExpr(expr interface{}) interface{} {
	if callback, ok := expr.(func() interface{}); ok {
		return callback()
	} else {
		return expr
	}
}

func (this Iffy) Elif(condition bool, expr interface{}) Iffy {
	if !this.matched && condition {
		this.matched = true
		this.result = handleExpr(expr)
	}
	return this
}

func (this Iffy) Else(expr interface{}) interface{} {
	if !this.matched {
		this.result = handleExpr(expr)
	}
	return this.result
}
