package decimal_wrap1

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func TFunc() {
	v, _ := decimal.NewFromString("-123.4567")
	fmt.Printf("v1.3.1: %s\n", v)
}
