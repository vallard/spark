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

func TestGetPerson(t *testing.T) {
	s = getSpark(t)
	p, err := s.GetPerson("Y2lzY29zcGFyazovL3VzL1BFT1BMRS83MDE2MDRmMS04ZTk2LTRiNzEtOTE1Mi0wODY0YTkxYmM2MTM")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Person: ", p)
}
