package spark

import (
	"fmt"
	"testing"
)

func TestGetMyself(t *testing.T) {
	s = getSpark(t)
	me, err := s.GetMyself()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Me: ", me)
}
