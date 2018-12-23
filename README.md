RandomStr
=========

RandomStr is a Go library to generate random strings based on a set of criteria.

Install
=======

```bash
go get -u github.com/jgweir/randstring
```

Usage
=====

To generate a random string with 8 characters:

```golang
rs := randomstr.New(8)
str, err := rs.Build()
```

To generate a random string that only includes lowercase characters and special characters

```golang
rs := randomstr.New(10)
rs.Capitals(false)
rs.Digits(false)
str, err := rs.Build()
```

To generate a random string with a prefix and posfix (note the total length of the string will be the number of characters in the prefix, posfix and the random string length provided)

```golang
prefix := []rune("user-")
posfix := []rune("@email.com")
rs := randomstr.New(6)
rs.Capitals(false)
rs.Lowercase(true)
rs.Digits(false)
rs.Specials(true)

rs.SetPrefix(prefix)
rs.SetPosfix(posfix)
str, err := rs.Build()
```

To generate a random string with the swedish alphabet

```golang
cchars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ")
lchars := []rune("abcdefghijklmnopqrstuvwxyzåäö")

rs := randomstr.New(6)
rs.SetCapitalChars(cchars)
rs.SetLowercaseChars(lchars)
str, err := rs.Build()
```

Running Example
---------------

```golang
package main

import (
    "fmt"
    "github.com/jgweir/randstring"
)

func main() {

    posfix := []rune("@email.com")

    rs := randstring.New(8)
    rs.Capitals(false)
    rs.Specials(false)
    rs.Digits(false)
    rs.SetPosfix(posfix)

    for i := 0; i < 6; i++ {
        str, _ := rs.Build()
        fmt.Println(str)
    }
}

// output
// dweuszrf@email.com
// iybiiaid@email.com
// mdifuicd@email.com
// qudcuiss@email.com
// cmiudspc@email.com
// pincissu@email.com
```


