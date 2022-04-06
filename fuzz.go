package gson

import (
	"fmt"
)

func Fuzz(data []byte) (score int) {
	for _, ctor := range []func() any{
		func() any { return new(any) },
		func() any { return new(map[string]any) },
		func() any { return new([]any) },
	} {
		v := ctor()
		err := Unmarshal(data, v)
		if err != nil {
			continue
		}
		score = 1

		m, err := Marshal(v)
		if err != nil {
			fmt.Printf("v=%#v\n", v)
			panic(err)
		}

		u := ctor()
		err = Unmarshal(m, u)
		if err != nil {
			fmt.Printf("v=%#v\n", v)
			fmt.Printf("m=%s\n", m)
			panic(err)
		}
	}

	return
}
