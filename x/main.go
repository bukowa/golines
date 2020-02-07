
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
