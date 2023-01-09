package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cel int

func main() {

	latin := [100]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII",
		"IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI",
		"XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII",
		"XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX",
		"XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII",
		"XLVIII", "IL", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII",
		"LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LIXX",
		"LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LIXXX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII",
		"LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

	fmt.Print("Input: ")
	var mainstr string

	fmt.Scanln(&mainstr)

	strcheck := mainstr
	symbols := []rune(strcheck)
	q := 0
	w := 0
	e := 0
	r := 0
	for i := 0; i < len(symbols); i++ {
		if symbols[i] == '+' {
			q++
		}
		if symbols[i] == '-' {
			w++
		}
		if symbols[i] == '*' {
			e++
		}
		if symbols[i] == '/' {
			r++
		}
	}
	if q+w+e+r > 1 {
		println("Вывод ошибки, так как формат математической операции не " +
			"удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(1)
	}

	if strings.Contains(mainstr, "+") == false && strings.Contains(mainstr, "-") == false &&
		strings.Contains(mainstr, "/") == false && strings.Contains(mainstr, "*") == false {
		println("Output: ОШИБКА: не корректный ввод!")
		os.Exit(1)
	}

	if strings.Contains(mainstr, "+") == true {
		words := strings.Split(mainstr, "+")
		if latinbool(words[0], words[1]) == 1 {
			x := latintoarab(words[0])
			y := latintoarab(words[1])
			print("Output: ")
			println(latin[x+y-1])
			os.Exit(1)
		}
		if latinbool(words[0], words[1]) == 2 {
			println("Output: ОШИБКА: используются одновременно разные системы счисления!")
			os.Exit(1)
		}
		a, _ := strconv.Atoi(words[0])
		b, _ := strconv.Atoi(words[1])

		if a < 0 || a > 10 || b < 0 || b > 10 {
			println("Output: Ошибка: вы ввели не корректные значения операнда")
			os.Exit(1)
		} else if a >= 0 && a <= 10 && b >= 0 && b <= 10 {
			c := a + b
			print("Output: ")
			println(c)
			os.Exit(1)
		}
	}
	if strings.Contains(mainstr, "-") == true {
		words := strings.Split(mainstr, "-")
		if latinbool(words[0], words[1]) == 1 {
			x := latintoarab(words[0])
			y := latintoarab(words[1])
			z := x - y
			if z < 1 {
				println("Output: ОШИБКА: результат меньше 1!")
				os.Exit(1)
			}
			print("Output: ")
			println(latin[z-1])
			os.Exit(1)
		}
		if latinbool(words[0], words[1]) == 2 {
			println("Output: ОШИБКА: используются одновременно разные системы счисления!")
			os.Exit(1)
		}

		a1, _ := strconv.Atoi(words[0])
		b1, _ := strconv.Atoi(words[1])
		if a1 < 0 || a1 > 10 || b1 < 0 || b1 > 10 {
			println("Output: Ошибка: вы ввели не корректные значения операнда")
			os.Exit(1)
		} else if a1 >= 0 && a1 <= 10 && b1 >= 0 && b1 <= 10 {
			c := a1 - b1
			print("Output: ")
			println(c)
			os.Exit(1)
		}

	}

	if strings.Contains(mainstr, "*") == true {
		words := strings.Split(mainstr, "*")
		if latinbool(words[0], words[1]) == 1 {
			x := latintoarab(words[0])
			y := latintoarab(words[1])
			print("Output: ")
			println(latin[x*y-1])
			os.Exit(1)
		}
		if latinbool(words[0], words[1]) == 2 {
			println("Output: ОШИБКА: используются одновременно разные системы счисления")
			os.Exit(1)
		}
		a2, _ := strconv.Atoi(words[0])
		b2, _ := strconv.Atoi(words[1])
		if a2 < 0 || a2 > 10 || b2 < 0 || b2 > 10 {
			println("Output: Ошибка: вы ввели не корректные значения операнда")
			os.Exit(1)
		} else if a2 >= 0 && a2 <= 10 && b2 >= 0 && b2 <= 10 {
			c := a2 * b2
			print("Output: ")
			println(c)
			os.Exit(1)
		}

	}

	if strings.Contains(mainstr, "/") == true {
		words := strings.Split(mainstr, "/")
		if latinbool(words[0], words[1]) == 1 {
			x := latintoarab(words[0])
			y := latintoarab(words[1])
			cel = x / y
			println(latin[cel-1])
			os.Exit(1)
		}
		if latinbool(words[0], words[1]) == 2 {
			println("Output: ОШИБКА: используются одновременно разные системы счисления!")
			os.Exit(1)
		}
		a2, _ := strconv.Atoi(words[0])
		b2, _ := strconv.Atoi(words[1])
		if a2 < 0 || a2 > 10 || b2 < 0 || b2 > 10 {
			println("Output: Ошибка: вы ввели не корректные значения операнда")
			os.Exit(1)
		} else if a2 >= 0 && a2 <= 10 && b2 >= 0 && b2 <= 10 {
			if b2 == 0 {
				println("Output: на ноль делить нельзя!")
				os.Exit(1)
			}
			if a2/b2 < 1 {
				println("Output: ИСКЛЮЧЕНИЕ: результат меньше 1!")
				os.Exit(1)
			} else {
				cel := a2 / b2
				print("Output: ")
				println(cel)
			}
		}

	}

}

func latintoarab(cifralatin string) int {
	switch cifralatin {
	case "I":
		return 1

	case "II":
		return 2

	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		return 0
	}

}

func latinbool(str1, str2 string) int {
	if (str1 == "I" || str1 == "II" || str1 == "III" || str1 == "IV" || str1 == "V" || str1 == "VI" ||
		str1 == "VII" || str1 == "VIII" || str1 == "IX" || str1 == "X") && (str2 == "I" || str2 == "II" ||
		str2 == "III" || str2 == "IV" || str2 == "V" || str2 == "VI" ||
		str2 == "VII" || str2 == "VIII" || str2 == "IX" || str2 == "X") {
		return 1
	}
	if ((str1 == "I" || str1 == "II" || str1 == "III" || str1 == "IV" || str1 == "V" || str1 == "VI" ||
		str1 == "VII" || str1 == "VIII" || str1 == "IX" || str1 == "X") && (str2 != "I" || str2 != "II" ||
		str2 != "III" || str2 != "IV" || str2 != "V" || str2 != "VI" ||
		str2 != "VII" || str2 != "VIII" || str2 != "IX" || str2 != "X")) || ((str1 != "I" || str1 != "II" ||
		str1 == "III" || str1 == "IV" || str1 == "V" || str1 == "VI" ||
		str1 != "VII" || str1 != "VIII" || str1 != "IX" || str1 != "X") && (str2 == "I" || str2 == "II" ||
		str2 == "III" || str2 == "IV" || str2 == "V" || str2 == "VI" ||
		str2 == "VII" || str2 == "VIII" || str2 == "IX" || str2 == "X")) {
		return 2
	}
	return 3

}
