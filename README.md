# gson

You can use spread json tag by using this package.

## Usage

## Thanks

inspired by `encoding/json`
https://github.com/golang/go/tree/master/src/encoding/json

## Origin

ex)

```go
type good[T any] struct {
    Name string
    Free T `json:"..."`
}

type more struct {
    Age int
    Country string
}

var foo = good[more]{
    Name: "Foo",
    more: more{
        Age: 20,
        Country: "U.S.A",
    },
}

```

If you encode it in json. It'll be like this.

```
{"Name": "Foo", "Age": 20, "Country": "U.S.A"}
```

https://github.com/golang/go/issues/52138
