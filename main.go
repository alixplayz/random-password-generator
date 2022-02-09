package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"randstr"
	"strconv"
	"strings"
	"time"
)

// define some vars
var version string = "0.1"
var author string = "alixplayz"
var project string = "random-password-generator"

var selectedOptions int = 0
var totalOptionsThatAddCharacters int = 2

func main() {

	// welcome message
	fmt.Println("Welcome to the Random Password Generator [v" + version + "] (type \"continue\" to continue)")
	scannerContinue := bufio.NewScanner(os.Stdin)
	scannerContinue.Scan()
	sContinueText := scannerContinue.Text()

	if sContinueText != "continue" {
		fmt.Println("Canceled. Couldn't detect \"continue\"")
		return
	}

	// option 1
	// include numbers?
	fmt.Println("[Option 1] Include numbers? (use true or false)")
	scannerNumbers := bufio.NewScanner(os.Stdin)
	scannerNumbers.Scan()
	sNumbersText, _ := strconv.ParseBool(scannerNumbers.Text())

	if sNumbersText {
		selectedOptions++
	}

	// option 2
	// include normal lowercase characters?
	fmt.Println("[Option 2] Include normal lowercase characters?")
	scannerChars := bufio.NewScanner(os.Stdin)
	scannerChars.Scan()
	sCharsText, _ := strconv.ParseBool(scannerChars.Text())

	if sCharsText {
		selectedOptions++
	}

	// option 3
	// password total length?
	fmt.Println("[Option 3] Total password length?")
	scannerLength := bufio.NewScanner(os.Stdin)
	scannerLength.Scan()
	sLengthText, _ := strconv.ParseInt(scannerLength.Text(), 10, 64)

	// final password
	fmt.Println("")
	fmt.Println(generatePass(sNumbersText, sCharsText, int(sLengthText)))
	fmt.Println("")

	// wait for user input until the app closes itself
	fmt.Println("Done. Thanks for using! (type \"exit\" to exit)")
	scannerExit := bufio.NewScanner(os.Stdin)
	scannerExit.Scan()
	scannerExitText := scannerExit.Text()

	if scannerExitText == "exit" {
		return
	}
}

func generatePass(numbers bool, normal bool, length int) string {
	var finalPass []string
	rand.Seed(time.Now().Unix())

	// add all the option "ids" to the to an array for the switch statement to read
	var allEnabledOptions []string
	if numbers {
		allEnabledOptions = append(allEnabledOptions, "1")
	}
	if normal {
		allEnabledOptions = append(allEnabledOptions, "2")
	}
	allEnabledOptionsString := strings.Join(allEnabledOptions, "")

	// append all the selected options into the array
	for i := 0; i < length; {

		// randomly select on of the "number ids" of the selected options
		randSelect := randstr.String(1, allEnabledOptionsString)

		switch randSelect {
		case "1":
			if numbers {
				finalPass = append(finalPass, strconv.Itoa(+rand.Intn(9-0)))
			}
		case "2":
			if normal {
				finalPass = append(finalPass, randstr.String(1, "abcdefghijklmnopqrstuvwxyz"))
			}
		default:
			println("Something went horribly wrong.")
			os.Exit(69)
		}

		i++
	}

	// make the array into a normal string
	finalPassString := strings.Join(finalPass, "")
	return finalPassString
}
