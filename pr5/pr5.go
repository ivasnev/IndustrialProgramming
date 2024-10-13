package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Линейное программирование
func task_1_1() {
	var number string
	var baseFrom, baseTo int

	fmt.Println("Введите число, исходную систему счисления и конечную систему счисления:")
	fmt.Scan(&number, &baseFrom, &baseTo)

	num, err := strconv.ParseInt(number, baseFrom, 64)
	if err != nil {
		fmt.Println("Ошибка при переводе числа:", err)
		return
	}

	result := strconv.FormatInt(num, baseTo)
	fmt.Printf("Результат: %s\n", result)
}

// 2. Решение квадратного уравнения
func task_1_2() {
	var a, b, c float64
	fmt.Println("Введите коэффициенты a, b, c:")
	fmt.Scan(&a, &b, &c)

	d := b*b - 4*a*c
	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("Корни уравнения: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if d == 0 {
		x := -b / (2 * a)
		fmt.Printf("Один корень: x = %.2f\n", x)
	} else {
		real := -b / (2 * a)
		imaginary := math.Sqrt(-d) / (2 * a)
		fmt.Printf("Комплексные корни: x1 = %.2f + %.2fi, x2 = %.2f - %.2fi\n", real, imaginary, real, imaginary)
	}
}

// 3. Сортировка чисел по модулю
func task_1_3() {
	var n int
	fmt.Println("Введите количество чисел:")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if math.Abs(float64(arr[j])) > math.Abs(float64(arr[j+1])) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Отсортированный массив:", arr)
}

// 4. Слияние двух отсортированных массивов
func task_1_4() {
	var n1, n2 int
	fmt.Println("Введите размер первого массива:")
	fmt.Scan(&n1)
	fmt.Println("Введите первый массив:")
	arr1 := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Println("Введите размер второго массива:")
	fmt.Scan(&n2)
	fmt.Println("Введите второй массив:")
	arr2 := make([]int, n2)
	for i := 0; i < n2; i++ {
		fmt.Scan(&arr2[i])
	}

	result := mergeSortedArrays(arr1, arr2)
	fmt.Println("Объединенный массив:", result)
}

func mergeSortedArrays(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}
	return result
}

// 5. Нахождение подстроки в строке без встроенных функций
func task_1_5() {
	var str, substr string
	fmt.Println("Введите строку:")
	fmt.Scan(&str)
	fmt.Println("Введите подстроку:")
	fmt.Scan(&substr)

	index := findSubstring(str, substr)
	fmt.Printf("Индекс первого вхождения подстроки: %d\n", index)
}

func findSubstring(str, substr string) int {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// Условия
func task_2_1() {
	// Калькулятор с расширенными операциями
	var a, b float64
	var operator string
	fmt.Println("Введите два числа и оператор (+, -, *, /, ^, %):")
	fmt.Scan(&a, &b, &operator)

	switch operator {
	case "+":
		fmt.Printf("Результат: %.2f\n", a+b)
	case "-":
		fmt.Printf("Результат: %.2f\n", a-b)
	case "*":
		fmt.Printf("Результат: %.2f\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Результат: %.2f\n", a/b)
		} else {
			fmt.Println("Ошибка: деление на ноль!")
		}
	case "^":
		fmt.Printf("Результат: %.2f\n", math.Pow(a, b))
	case "%":
		if b != 0 {
			fmt.Printf("Результат: %.2f\n", math.Mod(a, b))
		} else {
			fmt.Println("Ошибка: деление на ноль!")
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}

func task_2_2() {
	// Проверка палиндрома
	var str string
	fmt.Println("Введите строку:")
	fmt.Scan(&str)

	isPalindrome := checkPalindrome(str)
	fmt.Printf("Является палиндромом: %v\n", isPalindrome)
}

func checkPalindrome(str string) bool {
	// Приводим строку к нижнему регистру и удаляем пробелы, знаки препинания
	cleaned := strings.ToLower(str)
	cleaned = strings.ReplaceAll(cleaned, " ", "")
	cleaned = strings.ReplaceAll(cleaned, ".", "")
	cleaned = strings.ReplaceAll(cleaned, ",", "")
	cleaned = strings.ReplaceAll(cleaned, "!", "")
	cleaned = strings.ReplaceAll(cleaned, "?", "")

	// Проверяем, является ли строка палиндромом
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-i-1] {
			return false
		}
	}
	return true
}

func task_2_3() {
	// Нахождение пересечения трёх отрезков
	var x1, x2, y1, y2, z1, z2 int
	fmt.Println("Введите начальную и конечную точки первого отрезка:")
	fmt.Scan(&x1, &x2)
	fmt.Println("Введите начальную и конечную точки второго отрезка:")
	fmt.Scan(&y1, &y2)
	fmt.Println("Введите начальную и конечную точки третьего отрезка:")
	fmt.Scan(&z1, &z2)

	// Находим пересечение всех трёх отрезков
	start := max(x1, y1, z1)
	end := min(x2, y2, z2)

	if start <= end {
		fmt.Println("Отрезки пересекаются")
	} else {
		fmt.Println("Отрезки не пересекаются")
	}
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func task_2_4() {
	// Выбор самого длинного слова в предложении
	var sentence string
	fmt.Println("Введите предложение:")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ = reader.ReadString('\n')
	longestWord := findLongestWord(sentence)
	fmt.Printf("Самое длинное слово: %s\n", longestWord)
}

func findLongestWord(sentence string) string {
	words := strings.FieldsFunc(sentence, func(r rune) bool {
		return r == ' ' || r == '.' || r == ',' || r == '!' || r == '?'
	})

	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func task_2_5() {
	// Проверка високосного года
	var year int
	fmt.Println("Введите год:")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Printf("%d является високосным годом\n", year)
	} else {
		fmt.Printf("%d не является високосным годом\n", year)
	}
}

func isLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

// Циклы
func task_3_1() {
	// Числа Фибоначчи до определённого числа
	var n int
	fmt.Println("Введите число n:")
	fmt.Scan(&n)

	fibSequence := fibonacci(n)
	fmt.Println("Числа Фибоначчи:", fibSequence)
}

func fibonacci(n int) []int {
	seq := []int{0, 1}
	for {
		next := seq[len(seq)-1] + seq[len(seq)-2]
		if next > n {
			break
		}
		seq = append(seq, next)
	}
	return seq
}

func task_3_2() {
	// Определение простых чисел в диапазоне
	var start, end int
	fmt.Println("Введите начало и конец диапазона:")
	fmt.Scan(&start, &end)

	primes := findPrimesInRange(start, end)
	fmt.Println("Простые числа в диапазоне:", primes)
}

func findPrimesInRange(start, end int) []int {
	primes := []int{}
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func task_3_3() {
	// Числа Армстронга в заданном диапазоне
	var start, end int
	fmt.Println("Введите начало и конец диапазона:")
	fmt.Scan(&start, &end)

	armstrongNumbers := findArmstrongNumbers(start, end)
	fmt.Println("Числа Армстронга в диапазоне:", armstrongNumbers)
}

func findArmstrongNumbers(start, end int) []int {
	armstrongs := []int{}
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			armstrongs = append(armstrongs, i)
		}
	}
	return armstrongs
}

func isArmstrong(num int) bool {
	sum := 0
	temp := num
	digits := len(strconv.Itoa(num))
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == num
}

func task_3_4() {
	// Реверс строки
	var str string
	fmt.Println("Введите строку:")
	fmt.Scan(&str)

	reversed := reverseString(str)
	fmt.Println("Реверс строки:", reversed)
}

func reverseString(str string) string {
	chars := []rune(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func task_3_5() {
	// Нахождение наибольшего общего делителя (НОД)
	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)

	gcd := findGCD(a, b)
	fmt.Printf("НОД: %d\n", gcd)
}

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var taskId, moduleId int
	fmt.Println("Введите номер модуля")
	fmt.Scanf("%d\n", &moduleId)
	fmt.Println("Введите номер задачи")
	fmt.Scanf("%d\n", &taskId)
	xy := fmt.Sprintf("%d_%d", moduleId, taskId)
	switch xy {
	case "1_1":
		task_1_1()
	case "1_2":
		task_1_2()
	case "1_3":
		task_1_3()
	case "1_4":
		task_1_4()
	case "1_5":
		task_1_5()
	case "2_1":
		task_2_1()
	case "2_2":
		task_2_2()
	case "2_3":
		task_2_3()
	case "2_4":
		task_2_4()
	case "2_5":
		task_2_5()
	case "3_1":
		task_3_1()
	case "3_2":
		task_3_2()
	case "3_3":
		task_3_3()
	case "3_4":
		task_3_4()
	case "3_5":
		task_3_5()
	default:
		fmt.Println("Задача не найдена")
	}
}
