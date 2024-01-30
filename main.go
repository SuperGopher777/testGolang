package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		fmt.Println(calculator(text))
	}
}

func calculator(s string) string {
	result := ""
	operand := ""
	operandMap := map[string]string{"*": "*", "/": "/", "+": "+", "-": "-"}

	for key, val := range operandMap {
		if strings.Contains(s, key) {
			operand = val
			delete(operandMap, key)
		}
	}

	if operand == "" {
		panic("Dont have operand")
	}

	arrayNumbers := strings.Split(s, operand)

	if len(arrayNumbers) == 2 && len(operandMap) > 2 { // Если будет больше одного операнда и чисел более 2-х - panic()

		isRomanOne := isRomanNumeral(arrayNumbers[0]) // Проверка Арабская цифра или Римская - Паника если они разные в массиве
		isRomanTwo := isRomanNumeral(arrayNumbers[1])

		if isRomanOne && isRomanTwo {
			result = calculateRoman(arrayNumbers, operand)
		} else if !isRomanOne && !isRomanTwo {
			result = calculateArabic(arrayNumbers, operand)
		} else {
			panic("Only Arabic or Roman numbers")
		}

	} else {
		panic("We need ONLY 2 numbers")
	}

	return result
}

func calculateRoman(roman []string, operand string) string {
	resultArabic := 0
	resultRoman := ""

	romanNumbers := map[string]int{ // Тут можно еще сделать логику для любых чисел переделать мапу под вид I:1 X:10 L:50...
		"I":        1,
		"II":       2,
		"III":      3,
		"IV":       4,
		"V":        5,
		"VI":       6,
		"VII":      7,
		"VIII":     8,
		"IX":       9,
		"X":        10,
		"XI":       11,
		"XII":      12,
		"XIII":     13,
		"XIV":      14,
		"XV":       15,
		"XVI":      16,
		"XVII":     17,
		"XVIII":    18,
		"XIX":      19,
		"XX":       20,
		"XXI":      21,
		"XXII":     22,
		"XXIII":    23,
		"XXIV":     24,
		"XXV":      25,
		"XXVI":     26,
		"XXVII":    27,
		"XXVIII":   28,
		"XXIX":     29,
		"XXX":      30,
		"XXXI":     31,
		"XXXII":    32,
		"XXXIII":   33,
		"XXXIV":    34,
		"XXXV":     35,
		"XXXVI":    36,
		"XXXVII":   37,
		"XXXVIII":  38,
		"XXXIX":    39,
		"XL":       40,
		"XLI":      41,
		"XLII":     42,
		"XLIII":    43,
		"XLIV":     44,
		"XLV":      45,
		"XLVI":     46,
		"XLVII":    47,
		"XLVIII":   48,
		"XLIX":     49,
		"L":        50,
		"LI":       51,
		"LII":      52,
		"LIII":     53,
		"LIV":      54,
		"LV":       55,
		"LVI":      56,
		"LVII":     57,
		"LVIII":    58,
		"LIX":      59,
		"LX":       60,
		"LXI":      61,
		"LXII":     62,
		"LXIII":    63,
		"LXIV":     64,
		"LXV":      65,
		"LXVI":     66,
		"LXVII":    67,
		"LXVIII":   68,
		"LXIX":     69,
		"LXX":      70,
		"LXXI":     71,
		"LXXII":    72,
		"LXXIII":   73,
		"LXXIV":    74,
		"LXXV":     75,
		"LXXVI":    76,
		"LXXVII":   77,
		"LXXVIII":  78,
		"LXXIX":    79,
		"LXXX":     80,
		"LXXXI":    81,
		"LXXXII":   82,
		"LXXXIII":  83,
		"LXXXIV":   84,
		"LXXXV":    85,
		"LXXXVI":   86,
		"LXXXVII":  87,
		"LXXXVIII": 88,
		"LXXXIX":   89,
		"XC":       90,
		"XCI":      91,
		"XCII":     92,
		"XCIII":    93,
		"XCIV":     94,
		"XCV":      95,
		"XCVI":     96,
		"XCVII":    97,
		"XCVIII":   98,
		"XCIX":     99,
		"C":        100,
	}

	romanOne := strings.ToUpper(roman[0])
	romanTwo := strings.ToUpper(roman[1])

	numOne := romanNumbers[romanOne]
	numTwo := romanNumbers[romanTwo]

	if numOne < numTwo && operand == "-" {
		panic("Roman result cant be negative")
	}

	switch operand {
	case "*":
		resultArabic = numOne * numTwo
	case "-":
		resultArabic = numOne - numTwo
	case "/":
		resultArabic = numOne / numTwo
	case "+":
		resultArabic = numOne + numTwo
	}

	for k, v := range romanNumbers {
		if resultArabic == v {
			resultRoman = k
		}
	}

	return resultRoman
}

func calculateArabic(arabic []string, operand string) string {
	result := 0
	numOne, errOne := strconv.Atoi(arabic[0])
	if errOne != nil {
		panic("Cant convert numOne string to int")
	}
	numTwo, errTwo := strconv.Atoi(arabic[1])
	if errTwo != nil {
		panic("Cant convert numTwo string to int")
	}

	switch operand {
	case "*":
		result = numOne * numTwo
	case "-":
		result = numOne - numTwo
	case "/":
		result = numOne / numTwo
	case "+":
		result = numOne + numTwo
	}

	return strconv.Itoa(result)
}

func isRomanNumeral(s string) bool {
	upperString := strings.ToUpper(s)
	romanRegex := regexp.MustCompile(`^(M{0,3})(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	return romanRegex.MatchString(upperString)
}
