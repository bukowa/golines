# golines

### sources
* `file://`
* `http://`

```go
package main

import (
	"fmt"
)
import "github.com/bukowa/golines"

func main() {
	fmt.Println("Add!")
	lines := &golines.Lines{}
	fmt.Println(lines.Add("1\n"))
	fmt.Println(lines.Add([]byte("2\n")))
	fmt.Println(lines.Add([][]byte{[]byte("3\n"), []byte("4\n")}))
	fmt.Println(lines.Add([]string{"5\n", "6\n", "7"}))
	fmt.Println("Show!")
	fmt.Println(lines.String())
	fmt.Println("StringLines!")
	fmt.Println(lines.StringLines())
	fmt.Println("ByteLines!")
	fmt.Println(lines.ByteLines())

	checkBytes := []byte("123")
	fmt.Println("How many lines in Bytes? ", checkBytes)
	fmt.Println(lines.CountBytes(checkBytes))
	checkStrings := "123456"
	fmt.Println("How many lines in String?", checkStrings)
	fmt.Println(lines.CountString(checkStrings))
	fmt.Println("Map count to Bytes!", checkBytes)
	fmt.Println(lines.CountBytesMapNLines(checkBytes))
	checkStrings = "123456123456123"
	fmt.Println("Map count to Strings!", checkStrings)
	fmt.Println(lines.CountStringMapNLines(checkStrings))
	fmt.Println("Map Strings to count!", checkStrings)
	fmt.Println(lines.CountStringMapLineN(checkStrings))
	//fmt.Println("appears:", buff.CountBytes([]byte(str)))
	//fmt.Println("appears:", buff.CountString(str))

	// `lines` is `bytes.Buffer`
	lines = &golines.Lines{
		//Buffer: bytes.Buffer{},
		Parser: &golines.BasicParser{},
		Source: "file://main.go",
		//Source: "http://golang.org",
	}
	fmt.Println("Parser.PrefixMap!")
	fmt.Println(lines.Parser.PrefixMap())
	fmt.Println("Parse!")
	fmt.Println(lines.Parse())
	fmt.Println("ForLine!")
	lines.ForLine(func(i int, bytes []byte) error {
		fmt.Println(i, string(bytes))
		return nil
	})
	fmt.Println("StringLines!")
	fmt.Println(lines.StringLines())
	fmt.Println("ByteLines!")
	fmt.Println(lines.ByteLines())

}
````
````shell script
$ go run .
Add!
2 <nil>
2 <nil>
4 <nil>
5 <nil>
Show!
1
2
3
4
5
6
7
StringLines!
[1 2 3 4 5 6 7]
ByteLines!
[[49] [50] [51] [52] [53] [54] [55]]
How many lines in Bytes?  [49 50 51]
3
How many lines in String? 123456
6
Map count to Bytes! [49 50 51]
map[0:[[52] [53] [54] [55]] 1:[[49] [50] [51]]]
Map count to Strings! 123456123456123
map[0:[7] 2:[4 5 6] 3:[1 2 3]]
Map Strings to count! 123456123456123
map[1:3 2:3 3:3 4:2 5:2 6:2 7:0]
Parser.PrefixMap!
map[file://:0x630970 http://:0x630a90 https://:0x630a90]
Parse!
<nil>
ForLine!
0
1 package main
2
3 import (
4       "fmt"
5 )
6 import "github.com/bukowa/golines"
7
8 func main() {
9       fmt.Println("Add!")
10      lines := &golines.Lines{}
11      fmt.Println(lines.Add("1\n"))
12      fmt.Println(lines.Add([]byte("2\n")))
13      fmt.Println(lines.Add([][]byte{[]byte("3\n"), []byte("4\n")}))
14      fmt.Println(lines.Add([]string{"5\n", "6\n", "7"}))
15      fmt.Println("Show!")
16      fmt.Println(lines.String())
17      fmt.Println("StringLines!")
18      fmt.Println(lines.StringLines())
19      fmt.Println("ByteLines!")
20      fmt.Println(lines.ByteLines())
21
22      checkBytes := []byte("123")
23      fmt.Println("How many lines in Bytes? ", checkBytes)
24      fmt.Println(lines.CountBytes(checkBytes))
25      checkStrings := "123456"
26      fmt.Println("How many lines in String?", checkStrings)
27      fmt.Println(lines.CountString(checkStrings))
28      fmt.Println("Map count to Bytes!", checkBytes)
29      fmt.Println(lines.CountBytesMapNLines(checkBytes))
30      checkStrings = "123456123456123"
31      fmt.Println("Map count to Strings!", checkStrings)
32      fmt.Println(lines.CountStringMapNLines(checkStrings))
33      fmt.Println("Map Strings to count!", checkStrings)
34      fmt.Println(lines.CountStringMapLineN(checkStrings))
35      //fmt.Println("appears:", buff.CountBytes([]byte(str)))
36      //fmt.Println("appears:", buff.CountString(str))
37
38      // `lines` is `bytes.Buffer`
39      lines = &golines.Lines{
40              //Buffer: bytes.Buffer{},
41              Parser: &golines.BasicParser{},
42              Source: "file://main.go",
43              //Source: "http://golang.org",
44      }
45      fmt.Println("Parser.PrefixMap!")
46      fmt.Println(lines.Parser.PrefixMap())
47      fmt.Println("Parse!")
48      fmt.Println(lines.Parse())
49      fmt.Println("ForLine!")
50      lines.ForLine(func(i int, bytes []byte) error {
51              fmt.Println(i, string(bytes))
52              return nil
53      })
54      fmt.Println("StringLines!")
55      fmt.Println(lines.StringLines())
56      fmt.Println("ByteLines!")
57      fmt.Println(lines.ByteLines())
58
59 }
StringLines!
[ package main  import (        "fmt" ) import "github.com/bukowa/golines"  func main() {       fmt.Println("Add!")     lines := &golines.Lines{}       fmt.Println(lines.Add("1\n"))   fmt.Println(lines.Add([]byte("2\n")))   fmt.Println(lines.Add([][]byte{[]byte("3\n"), []byte("4\n")}))  fmt.Pri
ntln(lines.Add([]string{"5\n", "6\n", "7"}))    fmt.Println("Show!")    fmt.Println(lines.String())     fmt.Println("StringLines!")     fmt.Println(lines.StringLines())        fmt.Println("ByteLines!")       fmt.Println(lines.ByteLines())          checkBytes := []byte("123")     fmt.Println("Ho
w many lines in Bytes? ", checkBytes)   fmt.Println(lines.CountBytes(checkBytes))       checkStrings := "123456"        fmt.Println("How many lines in String?", checkStrings)  fmt.Println(lines.CountString(checkStrings))    fmt.Println("Map count to Bytes!", checkBytes)  fmt.Println(lines.Count
BytesMapNLines(checkBytes))     checkStrings = "123456123456123"        fmt.Println("Map count to Strings!", checkStrings)      fmt.Println(lines.CountStringMapNLines(checkStrings))   fmt.Println("Map Strings to count!", checkStrings)      fmt.Println(lines.CountStringMapLineN(checkStrings))
        //fmt.Println("appears:", buff.CountBytes([]byte(str)))         //fmt.Println("appears:", buff.CountString(str))        // `lines` is `bytes.Buffer`    lines = &golines.Lines{                 //Buffer: bytes.Buffer{},               Parser: &golines.BasicParser{},                 Source:
 "file://main.go",              //Source: "http://golang.org",  }       fmt.Println("Parser.PrefixMap!")        fmt.Println(lines.Parser.PrefixMap())   fmt.Println("Parse!")   fmt.Println(lines.Parse())      fmt.Println("ForLine!")         lines.ForLine(func(i int, bytes []byte) error {
                fmt.Println(i, string(bytes))           return nil      })      fmt.Println("StringLines!")     fmt.Println(lines.StringLines())        fmt.Println("ByteLines!")       fmt.Println(lines.ByteLines())  }]
ByteLines!
[[] [112 97 99 107 97 103 101 32 109 97 105 110] [] [105 109 112 111 114 116 32 40] [9 34 102 109 116 34] [41] [105 109 112 111 114 116 32 34 103 105 116 104 117 98 46 99 111 109 47 98 117 107 111 119 97 47 103 111 108 105 110 101 115 34] [] [102 117 110 99 32 109 97 105 110 40 41 32 123] [9 10
2 109 116 46 80 114 105 110 116 108 110 40 34 65 100 100 33 34 41] [9 108 105 110 101 115 32 58 61 32 38 103 111 108 105 110 101 115 46 76 105 110 101 115 123 125] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 65 100 100 40 34 49 92 110 34 41 41] [9 102 109 116 46 80 11
4 105 110 116 108 110 40 108 105 110 101 115 46 65 100 100 40 91 93 98 121 116 101 40 34 50 92 110 34 41 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 65 100 100 40 91 93 91 93 98 121 116 101 123 91 93 98 121 116 101 40 34 51 92 110 34 41 44 32 91 93 98 121 116 1
01 40 34 52 92 110 34 41 125 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 65 100 100 40 91 93 115 116 114 105 110 103 123 34 53 92 110 34 44 32 34 54 92 110 34 44 32 34 55 34 125 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 83 104 111 119 33 34 41]
[9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 83 116 114 105 110 103 40 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 83 116 114 105 110 103 76 105 110 101 115 33 34 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 83 116 114 105 11
0 103 76 105 110 101 115 40 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 66 121 116 101 76 105 110 101 115 33 34 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 66 121 116 101 76 105 110 101 115 40 41 41] [] [9 99 104 101 99 107 66 121 116 101 115 32 58 6
1 32 91 93 98 121 116 101 40 34 49 50 51 34 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 72 111 119 32 109 97 110 121 32 108 105 110 101 115 32 105 110 32 66 121 116 101 115 63 32 34 44 32 99 104 101 99 107 66 121 116 101 115 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105
 110 101 115 46 67 111 117 110 116 66 121 116 101 115 40 99 104 101 99 107 66 121 116 101 115 41 41] [9 99 104 101 99 107 83 116 114 105 110 103 115 32 58 61 32 34 49 50 51 52 53 54 34] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 72 111 119 32 109 97 110 121 32 108 105 110 101 115 32 105
 110 32 83 116 114 105 110 103 63 34 44 32 99 104 101 99 107 83 116 114 105 110 103 115 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 67 111 117 110 116 83 116 114 105 110 103 40 99 104 101 99 107 83 116 114 105 110 103 115 41 41] [9 102 109 116 46 80 114 105 110 11
6 108 110 40 34 77 97 112 32 99 111 117 110 116 32 116 111 32 66 121 116 101 115 33 34 44 32 99 104 101 99 107 66 121 116 101 115 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 67 111 117 110 116 66 121 116 101 115 77 97 112 78 76 105 110 101 115 40 99 104 101 99 107
 66 121 116 101 115 41 41] [9 99 104 101 99 107 83 116 114 105 110 103 115 32 61 32 34 49 50 51 52 53 54 49 50 51 52 53 54 49 50 51 34] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 77 97 112 32 99 111 117 110 116 32 116 111 32 83 116 114 105 110 103 115 33 34 44 32 99 104 101 99 107 83 11
6 114 105 110 103 115 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 67 111 117 110 116 83 116 114 105 110 103 77 97 112 78 76 105 110 101 115 40 99 104 101 99 107 83 116 114 105 110 103 115 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 77 97 112 32 83 11
6 114 105 110 103 115 32 116 111 32 99 111 117 110 116 33 34 44 32 99 104 101 99 107 83 116 114 105 110 103 115 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 67 111 117 110 116 83 116 114 105 110 103 77 97 112 76 105 110 101 78 40 99 104 101 99 107 83 116 114 105 11
0 103 115 41 41] [9 47 47 102 109 116 46 80 114 105 110 116 108 110 40 34 97 112 112 101 97 114 115 58 34 44 32 98 117 102 102 46 67 111 117 110 116 66 121 116 101 115 40 91 93 98 121 116 101 40 115 116 114 41 41 41] [9 47 47 102 109 116 46 80 114 105 110 116 108 110 40 34 97 112 112 101 97 114
 115 58 34 44 32 98 117 102 102 46 67 111 117 110 116 83 116 114 105 110 103 40 115 116 114 41 41] [] [9 47 47 32 96 108 105 110 101 115 96 32 105 115 32 96 98 121 116 101 115 46 66 117 102 102 101 114 96] [9 108 105 110 101 115 32 61 32 38 103 111 108 105 110 101 115 46 76 105 110 101 115 123]
 [9 9 47 47 66 117 102 102 101 114 58 32 98 121 116 101 115 46 66 117 102 102 101 114 123 125 44] [9 9 80 97 114 115 101 114 58 32 38 103 111 108 105 110 101 115 46 66 97 115 105 99 80 97 114 115 101 114 123 125 44] [9 9 83 111 117 114 99 101 58 32 34 102 105 108 101 58 47 47 109 97 105 110 46
103 111 34 44] [9 9 47 47 83 111 117 114 99 101 58 32 34 104 116 116 112 58 47 47 103 111 108 97 110 103 46 111 114 103 34 44] [9 125] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 80 97 114 115 101 114 46 80 114 101 102 105 120 77 97 112 33 34 41] [9 102 109 116 46 80 114 105 110 116 108
110 40 108 105 110 101 115 46 80 97 114 115 101 114 46 80 114 101 102 105 120 77 97 112 40 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 80 97 114 115 101 33 34 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 80 97 114 115 101 40 41 41] [9 102 109 116 46 8
0 114 105 110 116 108 110 40 34 70 111 114 76 105 110 101 33 34 41] [9 108 105 110 101 115 46 70 111 114 76 105 110 101 40 102 117 110 99 40 105 32 105 110 116 44 32 98 121 116 101 115 32 91 93 98 121 116 101 41 32 101 114 114 111 114 32 123] [9 9 102 109 116 46 80 114 105 110 116 108 110 40 10
5 44 32 115 116 114 105 110 103 40 98 121 116 101 115 41 41] [9 9 114 101 116 117 114 110 32 110 105 108] [9 125 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 83 116 114 105 110 103 76 105 110 101 115 33 34 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 83 1
16 114 105 110 103 76 105 110 101 115 40 41 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 34 66 121 116 101 76 105 110 101 115 33 34 41] [9 102 109 116 46 80 114 105 110 116 108 110 40 108 105 110 101 115 46 66 121 116 101 76 105 110 101 115 40 41 41] [] [125]]

```