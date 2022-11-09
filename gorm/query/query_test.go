package query

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {
	q := New().Equal("a", "A").StartWith("b", "B").EndWith("c", "C").Contain("d", "D")
	j, _ := json.Marshal(q)
	fmt.Println(string(j))
}
