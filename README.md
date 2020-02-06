# golines

### sources
* `file://`
* `http://`

```go
package main

import "fmt"
import "github.com/bukowa/golines"

func main() {
	s := &golines.Source{
		Parser: &golines.BasicParser{},
		Source: "file://main.go",
	}
	err := s.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Bytes())
	fmt.Println(s.String())
	fmt.Println(s.StringLines("", ""))
	fmt.Println(s.ByteLines(nil, nil))
	fmt.Println(s.ForLine(func(i int, bytes []byte) error {
		fmt.Println(i, string(bytes))
		return nil
	}))
}
```
```shell script
$ go run .
[112 97 99 107 97 103 101 32 109 97 105 110 10 10 105 109 112 111 114 116 32 34 102 109 116 34 10 105 109 112 111 114 116 32 34 103 105 116 104 117 98 46 99 111 109 47 98 117 107 111 119 97 47 103 111 108 105 110 101 115 34 10 10 102 117 110 99 32 109 97 105 110 40 41 32 123 10 9 115 32 58 61 3
2 38 103 111 108 105 110 101 115 46 83 111 117 114 99 101 123 10 9 9 80 97 114 115 101 114 58 32 38 103 111 108 105 110 101 115 46 66 97 115 105 99 80 97 114 115 101 114 123 125 44 10 9 9 83 111 117 114 99 101 58 32 34 102 105 108 101 58 47 47 109 97 105 110 46 103 111 34 44 10 9 125 10 9 101 1
14 114 32 58 61 32 115 46 80 97 114 115 101 40 41 10 9 105 102 32 101 114 114 32 33 61 32 110 105 108 32 123 10 9 9 112 97 110 105 99 40 101 114 114 41 10 9 125 10 9 102 109 116 46 80 114 105 110 116 108 110 40 115 46 66 121 116 101 115 40 41 41 10 9 102 109 116 46 80 114 105 110 116 108 110 40
 115 46 83 116 114 105 110 103 40 41 41 10 9 102 109 116 46 80 114 105 110 116 108 110 40 115 46 83 116 114 105 110 103 76 105 110 101 115 40 34 34 44 32 34 34 41 41 10 9 102 109 116 46 80 114 105 110 116 108 110 40 115 46 66 121 116 101 76 105 110 101 115 40 110 105 108 44 32 110 105 108 41 41
 10 9 102 109 116 46 80 114 105 110 116 108 110 40 115 46 70 111 114 76 105 110 101 40 102 117 110 99 40 105 32 105 110 116 44 32 98 121 116 101 115 32 91 93 98 121 116 101 41 32 101 114 114 111 114 32 123 10 9 9 102 109 116 46 80 114 105 110 116 108 110 40 105 44 32 115 116 114 105 110 103 40
98 121 116 101 115 41 41 10 9 9 114 101 116 117 114 110 32 110 105 108 10 9 125 41 41 10 125]
package main

import "fmt"
import "github.com/bukowa/golines"

func main() {
        s := &golines.Source{
                Parser: &golines.BasicParser{},
                Source: "file://main.go",
        }
        err := s.Parse()
        if err != nil {
                panic(err)
        }
        fmt.Println(s.Bytes())
        fmt.Println(s.String())
        fmt.Println(s.StringLines("", ""))
        fmt.Println(s.ByteLines(nil, nil))
        fmt.Println(s.ForLine(func(i int, bytes []byte) error {
                fmt.Println(i, string(bytes))
                return nil
        }))
}
[package main  import "fmt" import "github.com/bukowa/golines"  func main() {   s := &golines.Source{           Parser: &golines.BasicParser{},                 Source: "file://main.go",       }       err := s.Parse()        if err != nil {                 panic(err)      }       fmt.Println(s.B
ytes())         fmt.Println(s.String())         fmt.Println(s.StringLines("", ""))      fmt.Println(s.ByteLines(nil, nil))      fmt.Println(s.ForLine(func(i int, bytes []byte) error {                 fmt.Println(i, string(bytes))           return nil      })) }] <nil>
[[112 97 99 107 97 103 101 32 109 97 105 110] [] [105 109 112 111 114 116 32 34 102 109 116 34] [105 109 112 111 114 116 32 34 103 105 116 104 117 98 46 99 111 109 47 98 117 107 111 119 97 47 103 111 108 105 110 101 115 34] [] [102 117 110 99 32 109 97 105 110 40 41 32 123] [9 115 32 58 61 32 3
8 103 111 108 105 110 101 115 46 83 111 117 114 99 101 123] [9 9 80 97 114 115 101 114 58 32 38 103 111 108 105 110 101 115 46 66 97 115 105 99 80 97 114 115 101 114 123 125 44] [9 9 83 111 117 114 99 101 58 32 34 102 105 108 101 58 47 47 109 97 105 110 46 103 111 34 44] [9 125] [9 101 114 114
32 58 61 32 115 46 80 97 114 115 101 40 41] [9 105 102 32 101 114 114 32 33 61 32 110 105 108 32 123] [9 9 112 97 110 105 99 40 101 114 114 41] [9 125] [9 102 109 116 46 80 114 105 110 116 108 110 40 115 46 66 121 116 101 115 40 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 115 46 83 1
16 114 105 110 103 40 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 115 46 83 116 114 105 110 103 76 105 110 101 115 40 34 34 44 32 34 34 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 115 46 66 121 116 101 76 105 110 101 115 40 110 105 108 44 32 110 105 108 41 41] [9 102 109 1
16 46 80 114 105 110 116 108 110 40 115 46 70 111 114 76 105 110 101 40 102 117 110 99 40 105 32 105 110 116 44 32 98 121 116 101 115 32 91 93 98 121 116 101 41 32 101 114 114 111 114 32 123] [9 9 102 109 116 46 80 114 105 110 116 108 110 40 105 44 32 115 116 114 105 110 103 40 98 121 116 101 1
15 41 41] [9 9 114 101 116 117 114 110 32 110 105 108] [9 125 41 41] [125]] <nil>
0 package main
1
2 import "fmt"
3 import "github.com/bukowa/golines"
4
5 func main() {
6       s := &golines.Source{
7               Parser: &golines.BasicParser{},
8               Source: "file://main.go",
9       }
10      err := s.Parse()
11      if err != nil {
12              panic(err)
13      }
14      fmt.Println(s.Bytes())
15      fmt.Println(s.String())
16      fmt.Println(s.StringLines("", ""))
17      fmt.Println(s.ByteLines(nil, nil))
18      fmt.Println(s.ForLine(func(i int, bytes []byte) error {
19              fmt.Println(i, string(bytes))
20              return nil
21      }))
22 }
<nil>


```