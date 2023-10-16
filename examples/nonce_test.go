package examples

import (
	"fmt"

	"github.com/o289697/goutils"
)

func ExampleNonceStr() {
	nonceStr := goutils.NonceStr(16)
	fmt.Println(nonceStr)

	// Output: KbCnyGRPrNouchNm
}

func ExampleNoncePassword() {
	noncePassword := goutils.NoncePassword(12)
	fmt.Println(noncePassword)

	// Output: HF3a1Wc%jclq
}
