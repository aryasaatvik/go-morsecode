 /*************************************************************************
 * go-morsecode
 * Saatvik Arya
 * <aryasaatvik@gmail.com>
 */
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var morseMap = map[string]string{
	"a": ". _",
	"b": "_ . . .",
	"c": "_ . _ .",
	"d": "_ . .",
	"e": ".",
	"f": ". . _ .",
	"g": "_ _ .",
	"h": ". . . .",
	"i": ". .",
	"j": ". _ _ _",
	"k": "_ . _",
	"l": ". _ . .",
	"m": "_ _",
	"n": "_ .",
	"o": "_ _ _",
	"p": "._ _ .",
	"q": "_ _ . _",
	"r": ". _ .",
	"s": ". . .",
	"t": "_",
	"u": ". . _",
	"v": ". . . _",
	"w": ". _ _",
	"x": "_ . . _",
	"y": "_ . _ _",
	"z": "_ _..",
	"0": "_ _ _ _ _",
	"1": ". _ _ _ _",
	"2": ". . _ _ _",
	"3": ". . . _ _",
	"4": ". . . . _",
	"5": ". . . . .",
	"6": "_ . . . .",
	"7": "_ _ . . .",
	"8": "_ _ _ . .",
	"9": "_ _ _ _ .",
	" ": " ",
}

var charMap = map[string]string{
	"._":    "a",
	"_...":  "b",
	"_._.":  "c",
	"_..":   "d",
	".":     "e",
	".._.":  "f",
	"__.":   "g",
	"....":  "h",
	"..":    "i",
	".___":  "j",
	"_._":   "k",
	"._..":  "l",
	"__":    "m",
	"_.":    "n",
	"___":   "o",
	".__.":  "p",
	"__._":  "q",
	"._.":   "r",
	"...":   "s",
	"_":     "t",
	".._":   "u",
	"..._":  "v",
	".__":   "w",
	"_.._":  "x",
	"_.__":  "y",
	"__..":  "z",
	"_____": "0",
	".____": "1",
	"..___": "2",
	"...__": "3",
	"...._": "4",
	".....": "5",
	"_....": "6",
	"__...": "7",
	"___..": "8",
	"____.": "9",
	" ":     " ",
}

//morseEncode
func morseEncode(input string) string {
	var buffer bytes.Buffer

	for _, c := range input {
		key := string(c)
		_, exists := morseMap[key]
		if exists {
			buffer.WriteString(morseMap[key])
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

//morseDecode
func morseDecode(input string) string {
	buf := bytes.NewBufferString("")
	for _, c := range strings.Split(input, " ") {
		v, ok := charMap[c]
		if !ok {
			continue
		}

		buf.WriteString(v)
	}

	return strings.TrimSpace(buf.String())
}

func main() {
	app := cli.NewApp()
	app.Name = "go-morsecode"
	app.Author = "Saatvik Arya"
	app.Email = "aryasaatvik@gmail.com"
	app.Version = "1.0.0"
	app.Usage = "morseEncode and morseDecode"
	app.Commands = []cli.Command{
		{
			Name:    "encode",
			Aliases: []string{"en"},
			Usage:   "outputs text to morsecode",
			Action: func(c *cli.Context) {
				args := os.Args[2:]
				fmt.Println("encoding to morseCode: ", os.Args[2:])
				fmt.Println()
				inputString := strings.ToLower(strings.Join(args, " "))
				fmt.Printf("%s\n", morseEncode(inputString))
				fmt.Println()
			},
		},
		{
			Name:    "decode",
			Aliases: []string{"de"},
			Usage:   "outputs morsecode in english",
			Action: func(c *cli.Context) {
				args := os.Args[2:]
				fmt.Println("decoding morseCode: ", os.Args[2:])
				fmt.Println()
				decodeString := strings.ToLower(strings.Join(args, " "))
				fmt.Printf("%s\n", morseDecode(decodeString))
				fmt.Println()
			},
		},
	}
	app.Action = func(c *cli.Context) {
		println("### go-morsecode by Saatvik Arya ###\n")
		println("usage:\n 	help, -h, --help: prints help\n 	encode: encodes text to morsecode\n 	decode: decodes morsecode to text\n")
	}
	app.Run(os.Args)
}
