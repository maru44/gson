package gson

import (
	"testing"
)

func TestMarshalSpread(t *testing.T) {
	type base[T any] struct {
		Name     string
		Free     T `json:"..."`
		NickName string
	}

	type additional struct {
		Age     int
		IsValid bool
	}

	tests := []struct {
		name string
		in   any
		want string
	}{
		{
			name: "ok",
			in: base[additional]{
				Name: "Foo",
				Free: additional{
					Age: 20,
				},
				NickName: "Bar",
			},
			want: "{\"Name\":\"Foo\",\"Age\":20,\"IsValid\":false,\"NickName\":\"Bar\"}",
		},
		{
			name: "ok ptr",
			in: base[*additional]{
				Name: "Foo",
				Free: &additional{
					Age:     40,
					IsValid: true,
				},
				NickName: "Bar",
			},
			want: "{\"Name\":\"Foo\",\"Age\":40,\"IsValid\":true,\"NickName\":\"Bar\"}",
		},
		{
			name: "ok ignored",
			in: base[string]{
				Name:     "Foo",
				Free:     "ignored...",
				NickName: "Bar",
			},
			want: "{\"Name\":\"Foo\",\"NickName\":\"Bar\"}",
		},
	}

	for _, tt := range tests {
		tt := tt
		b, err := Marshal(tt.in)
		if err != nil {
			t.Errorf("test %s, unexpected failure: %v", tt.name, err)
		}

		if got := string(b); got != tt.want {
			t.Errorf("test %s, Marshal(%#v) = %q, want %q", tt.name, tt.in, got, tt.want)
		}
	}
}
