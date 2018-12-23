package randstring_test

import (
	"fmt"
	"github.com/jgweir/randstring"
)

func Example() {

	posfix := []rune("@email.com")

	rs := randstring.New(8)
	rs.Capitals(false)
	rs.Specials(false)
	rs.Digits(false)
	rs.SetPosfix(posfix)

	for i := 0; i < 6; i++ {
		str, _ := rs.Build()
		fmt.Println(str)
	}
}
