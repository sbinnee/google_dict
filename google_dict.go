package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	// "io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
	// "path/filepath"
	"strings"
	// "unicode/utf16"
	"github.com/mitchellh/go-wordwrap"
	"strconv"
)

const (
	// Foreground Color
	fg0 string = "\033[30m"
	fg1 string = "\033[31m" // Red
	fg2 string = "\033[32m" // Green
	fg3 string = "\033[33m" // Yellow Rather maron?
	fg4 string = "\033[34m" // Blue
	fg5 string = "\033[35m" // ?
	fg6 string = "\033[36m"
	fg7 string = "\033[37m"
	// FG Bright
	fg8  string = "\033[90m"
	fg9  string = "\033[91m"
	fg10 string = "\033[92m"
	fg11 string = "\033[93m" // Yellow
	fg12 string = "\033[94m"
	fg13 string = "\033[95m"
	fg14 string = "\033[96m" // Bluish?
	fg15 string = "\033[97m"

	// Background Color
	bg0 string = "\033[40m"
	bg1 string = "\033[41m" // Red
	bg2 string = "\033[42m" // Green
	bg3 string = "\033[43m" // Yellow
	bg4 string = "\033[44m" // Blue
	bg5 string = "\033[45m" // ?
	bg6 string = "\033[46m"
	bg7 string = "\033[47m"
	// BG Bright
	bg8  string = "\033[100m" // Gray, Good with fg15
	bg9  string = "\033[101m"
	bg10 string = "\033[102m"
	bg11 string = "\033[103m" // Yellow
	bg12 string = "\033[104m"
	bg13 string = "\033[105m"
	bg14 string = "\033[106m" // Bluish?
	bg15 string = "\033[107m"

	// Style
	sD string = "\033[0m" // Default
	sB string = "\033[1m" // Bold
	sI string = "\033[3m" // Italic
	sU string = "\033[4m" //Underline

	dd string = fg0 + bg0 + sD
)

func printStyle(txt string, st ...string) {
	fmt.Print(strings.Join(st, ""), " ", txt, " ", dd)
}

func printlnStyle(txt string, st ...string) {
	fmt.Print(strings.Join(st, ""), " ", txt, " ", dd, "\n")
}

func main() {
	columns, _ := strconv.Atoi(os.Getenv("COLUMNS"))
	ncols := uint(columns)
	// f := filepath.Join(os.Getenv("HOME"), "hello_indent.json")
	lang := "en_US"
	args := os.Args
	var word string
	if l := len(os.Args); l == 1 {
		log.SetFlags(0)
		log.Fatal("No word given")
	} else if l == 3 {
		if strings.ToUpper(args[1]) == "FR" {
			lang = "fr"
		}
		word = args[2]
	} else if l == 2 {
		word = args[1]
	}

	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/%s/%s", lang, word)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	content, _ := io.ReadAll(resp.Body)

	fmt.Println(string(content))

	// jsonparser.ObjectEach(content, func(value []byte, dataType jsonparser.ValueType, offset int))
	jsonparser.ArrayEach(content, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		txt, _ := jsonparser.GetString(value, "word")
		// fmt.Println(txt)
		printStyle(txt, sB, fg15, bg8)
		// "word", "phonetics", ["text", "audio"]
		jsonparser.ArrayEach(value, func(val []byte, dataType jsonparser.ValueType, offset int, err error) {
			txt, _ := jsonparser.GetString(val, "text")
			// audio, _ := jsonparser.GetString(val, "audio")
			printStyle("  "+txt, fg2)
			// fmt.Println(audio)
		}, "phonetics")
		fmt.Println() // Add a new line
		// "meanings"
		jsonparser.ArrayEach(value, func(val []byte, dataType jsonparser.ValueType, offset int, err error) {
			// "partOfSpeech"
			txt, _ := jsonparser.GetString(val, "partOfSpeech")
			// fmt.Println(txt)
			printlnStyle(txt, sI)
			// "definitions"
			ind := 0
			jsonparser.ArrayEach(val, func(va []byte, dataType jsonparser.ValueType, offset int, err error) {
				ind += 1
				def, _ := jsonparser.GetString(va, "definition")
				eg, _ := jsonparser.GetString(va, "example")
				// fmt.Println("\t", def)
				// fmt.Printf("\t%v. %v\n", ind, def)
				def = fmt.Sprintf("  %v. %v", ind, def)
				fmt.Println(wordwrap.WrapString(def, ncols))
				if len(eg) > 0 {
					fmt.Println("   ⤷", "\""+eg+"\"")
				}
				// syn, _ := jsonparser.GetString(va, "synonyms")
				// fmt.Println("\t", syn)
				var synonyms []string
				jsonparser.ArrayEach(va, func(v []byte, dataType jsonparser.ValueType, offset int, err error) {
					// fmt.Println(string(v))
					synonyms = append(synonyms, string(v))
				}, "synonyms")
				// fmt.Printf("%v\n", len(synonyms))
				if len(synonyms) > 0 {
					for i, s := range synonyms {
						if i%4 == 0 {
							fmt.Print("   ")
						}
						// fmt.Printf("\"%v\" ", s)
						// fmt.Printf("%v ", s)
						printStyle(s+" ", fg7)
						if i%4 == 3 {
							fmt.Print("\n")
						}
					}
				}
				fmt.Println()
			}, "definitions")
		}, "meanings")

	})

	// if err != nil {
	// 	log.Fatal(err)
	// }

}
