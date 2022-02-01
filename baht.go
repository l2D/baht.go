package bahtgo

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var EMPTY = ""
var ONE = "หนึ่ง"
var TWO = "สอง"
var THREE_TO_NINE = []string{"สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var ED = "เอ็ด"
var YEE = "ยี่"
var LAN = "ล้าน"

var DIGIT = []string{EMPTY, "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}

var ONES = []string{}        // To store replace "เอ็ด" to "หนึ่ง" in array
var TENS = []string{}        // To store replace "ยี่" to "สอง" in tens array
var SUB_HUNDRED = []string{} // To store sub-hunred strings
var SUB_TEN = []string{}     // To store sub-tens strings

func init() {
	ONES = append(ONES, "", ED, TWO)
	TENS = append(TENS, "", YEE)
	SUB_TEN = append(SUB_TEN, EMPTY, ONE, TWO)
	SUB_HUNDRED = append(SUB_HUNDRED, ONE+DIGIT[2], TWO+DIGIT[2])

	// Loop for append numbers
	for i := 1; i <= len(THREE_TO_NINE); i++ {
		ONES = append(ONES, THREE_TO_NINE[i-1])
		TENS = append(TENS, THREE_TO_NINE[i-1])
		SUB_TEN = append(SUB_TEN, THREE_TO_NINE[i-1])
	}

	// Loop for append "สิบ" to TENS
	for i := 1; i <= len(TENS); i++ {
		TENS[i-1] = TENS[i-1] + DIGIT[1]
	}

	// Loop for append "ร้อย" to SUB_HUNDRED
	for i := 1; i <= len(THREE_TO_NINE); i++ {
		SUB_HUNDRED = append(SUB_HUNDRED, THREE_TO_NINE[i-1]+DIGIT[2])
	}
}

func convert(s interface{}) string {
	baht := 0
	bahtStr := ""
	satang := 0
	satangStr := ""
	isNegative := false
	output := ""

	// Check type of s
	switch s := s.(type) {
	case float64: // handle cases for float64
		fmt.Println("s: ", s)
		isNegative = s < 0

		if isNegative {
			s = -s
		}

		// Seperate baht and satang
		// baht = int(s)
		// satang = int((s - float32(baht)) * 100)
		baht = int(math.Floor(s))
		satang = int(math.Floor((s - float64(baht)) * 100))
		/*
			f := 5.8
			ipart := int64(f)
			fmt.Println(ipart)
			decpart := fmt.Sprintf("%.7g", f-float64(ipart))[2:]
			fmt.Println(decpart)
		*/

		bahtStr = strconv.FormatInt(int64(s), 10)
		satangStr = strconv.FormatInt(int64(satang), 10)

	case int: // handle cases for int
		if s < 0 {
			isNegative = true
			fmt.Println("s: ", s)
		}

		baht = s
		satang = 0
		bahtStr = strconv.FormatInt(int64(baht), 10)
		satangStr = strconv.FormatInt(int64(satang), 10)

	case string: // handle cases for string

		// if s[0] == '-' {
		// 	isNegative = true
		// 	s = s[1:]
		// }

		leadingZeroPattern := regexp.MustCompile(`^0+`)
		negativeLeadingZeroPattern := regexp.MustCompile(`^-0+`)

		fmt.Println("s: ", s)
		if strings.HasPrefix(s, "-") {
			isNegative = true
			// s = negativeLeadingZeroPattern.ReplaceAllString(s, "0")
			if negativeLeadingZeroPattern.MatchString(s) {
				s = negativeLeadingZeroPattern.ReplaceAllString(s, "0")
			} else {
				s = s[1:]
			}
			// s = s[1:]
		} else {
			if leadingZeroPattern.MatchString(s) {
				// Check is s has leadingZeroPattern and decimalPoint
				if !strings.Contains(s, ".") {
					s = leadingZeroPattern.ReplaceAllString(s, "")
				}
			}
		}

		fmt.Println("s: ", s)

		// Find decimal point from string
		decimalPoint := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '.' {
				decimalPoint = i
				break
			}
		}

		// Convert string to int
		if decimalPoint > 0 {
			baht, _ = strconv.Atoi(s[:decimalPoint])
			satang, _ = strconv.Atoi(s[decimalPoint+1:])
			bahtStr = s[:decimalPoint]
			satangStr = s[decimalPoint+1:]
		} else {
			baht, _ = strconv.Atoi(s)
			bahtStr = s
			// bahtStr = leadingZeroPattern.ReplaceAllString(s, "")
		}

	default:
		return "Invalid type"
	}

	if (baht == 0) && (satang == 0) {
		return "ศูนย์บาทถ้วน"
	}

	if isNegative {
		output += "ลบ"
	}

	// Baht
	output += numberToWord(bahtStr)

	// Satang
	if satang > 0 {
		if baht > 0 {
			output += "บาท"
		}

		output += numberToWord(satangStr) + "สตางค์"
		// output += SUB_HUNDRED[satang] + "สตางค์"
	} else {
		output += "บาทถ้วน"
	}

	return output
}

func numberToWord(n string) string {
	var result string
	var length = len(n)

	for i := 0; i < length; i++ {
		d := n[i]
		di := length - i - 1
		diMod := di % 6
		isSib := diMod == 1

		if d == strconv.FormatInt(0, 10)[0] {
			// No-op
		} else if isSib && d == strconv.FormatInt(1, 10)[0] {
			result += DIGIT[diMod]
		} else if isSib && d == strconv.FormatInt(2, 10)[0] {
			result += YEE + DIGIT[diMod]
		} else if diMod == 0 && d == strconv.FormatInt(1, 10)[0] && i != 0 {
			result += ED
		} else {
			pInt, _ := strconv.ParseInt(string(d), 10, 64)
			result += SUB_TEN[pInt] + DIGIT[diMod]
		}

		if diMod == 0 && di != 0 {
			result += LAN
		}
	}

	return result
}

// func main() {
// 	timer := time.Now()
// 	// fmt.Println("ONES: ", ONES)
// 	// fmt.Println("TENS: ", TENS)
// 	// fmt.Println("SUB_HUNDRED: ", SUB_HUNDRED)
// 	// fmt.Println("SUB_TEN: ", SUB_TEN)
// 	fmt.Println(`convert(123): `, convert(123))
// 	fmt.Println(`convert(123.12): `, convert(123.12))
// 	fmt.Println(`convert(0.12): `, convert(0.12))
// 	fmt.Println(`convert(0): `, convert(0))
// 	fmt.Println(`convert("123"): `, convert("123"))
// 	fmt.Println(`convert("123.12"): `, convert("123.12"))
// 	fmt.Println(`convert("-123.12"): `, convert("-123.12"))
// 	output := func() string {
// 		if time.Since(timer).Milliseconds() > 0 {
// 			return strconv.FormatInt(time.Since(timer).Milliseconds(), 10) + " ms"
// 		} else {
// 			return strconv.FormatInt(time.Since(timer).Microseconds(), 10) + " μs"
// 		}
// 	}()
// 	fmt.Println(output)
// }
