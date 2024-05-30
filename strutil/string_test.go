package strutil

import (
	"fmt"
	"testing"
)

func TestToLowerFirstEachWord(t *testing.T) {
	fmt.Println(ToLowerFirstEachWord("UserInfoModel"))
	fmt.Println(ToLowerFirstEachWord("My Name Is Equal"))
}
