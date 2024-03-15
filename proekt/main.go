package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("введённые данные по элиментам")
	scanner := bufio.NewScanner(os.Stdin) //создние строку
	scanner.Scan()
	line := scanner.Text()

	var dann [10]string

	pods := strings.Split(line, " ")

	var i int

	for _, substring := range pods { // цапуск цикла для размищения подстрок по элиментам массива

		dann[i] = substring
		i++
	}

	if dann[3] != "" || dann[2] == "" || dann[1] == "" || dann[0] == "" {
		panic("Некорректный формат ввода, ожидалось: 'аргумент1 оператор аргумент2'") //  оприделение ввода оргументов количеством >2
	}

	operand1, operator, operand2 := dann[0], dann[1], dann[2] // Определение элименты массива в переменные для удобства

	//сравнение переменных на моответствие эдентичности системы счисления и на точтобы они были <=10 и >0

	znak1 := 0
	znak2 := 0
	sum := 0
	numr1 := 0 // обьявление переменных по которым будем сортировать римские цифры от орабских
	numr2 := 0
	switch operand1 {
	case "1":
		znak1 = 1
		numr1 = 1
	case "2":
		znak1 = 2
		numr1 = 1
	case "3":
		numr1 = 1
		znak1 = 3
	case "4":
		znak1 = 4
		numr1 = 1
	case "5":
		numr1 = 1
		znak1 = 5
	case "6":
		numr1 = 1
		znak1 = 6
	case "7":
		numr1 = 1
		znak1 = 7
	case "8":
		numr1 = 1
		znak1 = 8
	case "9":
		numr1 = 1
		znak1 = 9
	case "10":
		numr1 = 1
		znak1 = 10
	}

	switch operand2 {
	case "1":
		numr2 = 1
		znak2 = 1
	case "2":
		numr2 = 1
		znak2 = 2
	case "3":
		numr2 = 1
		znak2 = 3
	case "4":
		numr2 = 1
		znak2 = 4
	case "5":
		numr2 = 1
		znak2 = 5
	case "6":
		numr2 = 1
		znak2 = 6
	case "7":
		numr2 = 1
		znak2 = 7
	case "8":
		numr2 = 1
		znak2 = 8
	case "9":
		numr2 = 1
		znak2 = 9
	case "10":
		numr2 = 1
		znak2 = 10
	}

	switch operand1 {
	case "I":
		numr1 = 2
		znak1 = 1
	case "II":
		numr1 = 2
		znak1 = 2
	case "III":
		numr1 = 2
		znak1 = 3
	case "IV":
		numr1 = 2
		znak1 = 4
	case "V":
		numr1 = 2
		znak1 = 5
	case "VI":
		numr1 = 2
		znak1 = 6
	case "VII":
		numr1 = 2
		znak1 = 7
	case "VIII":
		numr1 = 2
		znak1 = 8
	case "IX":
		numr1 = 2
		znak1 = 9
	case "X":
		numr1 = 2
		znak1 = 10
	}

	switch operand2 {
	case "I":
		numr2 = 2
		znak2 = 1
	case "II":
		numr2 = 2
		znak2 = 2
	case "III":
		numr2 = 2
		znak2 = 3
	case "IV":
		numr2 = 2
		znak2 = 4
	case "V":
		numr2 = 2
		znak2 = 5
	case "VI":
		numr2 = 2
		znak2 = 6
	case "VII":
		numr2 = 2
		znak2 = 7
	case "VIII":
		numr2 = 2
		znak2 = 8
	case "IX":
		numr2 = 2
		znak2 = 9
	case "X":
		numr2 = 2
		znak2 = 10
	}

	if numr1 != 0 && numr2 != 0 {

		if (numr1 == 1 && numr2 == 1) || (numr1 == 2 && numr2 == 2) { // проверка на соответствие системы счислениея операндов

			if numr1 == 1 && numr2 == 1 { //  расчет арабских цифр
				switch operator {
				case "+":
					sum = znak1 + znak2
					fmt.Println(sum)
				case "-":
					sum = znak1 - znak2
					fmt.Println(sum)
				case "*":
					sum = znak1 * znak2
					fmt.Println(sum)
				case "/":
					sum = znak1 / znak2
					fmt.Println(sum)
				default:
					panic(" нелопустимый знак операции")
				}

				if numr1 == 2 && numr2 == 2 { //  расчет римских цифр

					switch operator {
					case "+":
						sum = znak1 + znak2

					case "-":
						sum = znak1 - znak2

					case "*":
						sum = znak1 * znak2

					case "/":
						sum = znak1 / znak2

					default:
						panic(" нелопустимый знак операции")
					}

					naborRim := [100]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

					if sum > 0 {
						fmt.Println(naborRim[sum-1])

					} else {
						panic("недопустимая операция")
					}

				}

			}
		} else {
			panic("не соответствие cистемы счисления 1-го операнта со 2-ым")
		}
	} else {
		panic("недопустимое цифры")
	}
}

