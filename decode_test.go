package gson

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type base[T any] struct {
		Name     string
		Free     T `json:"..."`
		NickName string
	}

	type additional struct {
		Age     int
		IsValid bool
	}

	in := "{\"Name\":\"Foo\",\"Age\":20,\"IsValid\":false,\"NickName\":\"Bar\"}"
	want := base[additional]{
		Name: "Foo",
		Free: additional{
			Age: 20,
		},
		NickName: "Bar",
	}

	got := base[additional]{}
	err := Unmarshal([]byte(in), &got)
	if err != nil {
		t.Errorf("cannot unmarshal %v", err)
	}
	if got != want {
		t.Errorf("not equal: got %v, want %v", got, want)
	}
}
