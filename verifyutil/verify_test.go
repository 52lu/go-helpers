package verifyutil

import (
	"fmt"
	"testing"
)

func TestParseIdCard(t *testing.T) {
	idCard := "341221198903048135"
	card, err := ParseIdCard(idCard)
	fmt.Println(card)
	fmt.Println(err)
}
