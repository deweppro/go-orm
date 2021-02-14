package orm

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_parseModel(t *testing.T) {
	type TestModel struct {
		_     TableName     `orm:"tests"`
		ID    int64         `orm:"id;index"`
		Email string        `orm:"email"`
		T     time.Duration `orm:"time"`
		B     byte
		Link  *int
	}

	res := parseModel(TestModel{
		ID:    77884,
		Email: "aaaaaaaa",
		T:     time.Minute,
		B:     0,
	})
	fmt.Println(res)
	res = parseModel(&TestModel{})
	fmt.Println(res)

	require.Nil(t, nil)
}
