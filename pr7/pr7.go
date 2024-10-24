package main

import (
	"fmt"
	"sort"
	"strings"
)

// 1. Функция вычисления площади треугольника
func task1() {
	var base, height float64
	fmt.Println("Введите основание и высоту треугольника:")
	fmt.Scanf("%f %f\n", &base, &height)
	area := triangleArea(base, height)
	fmt.Printf("Площадь треугольника: %.2f\n", area)
}

func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

// 2. Сортировка массива по возрастанию
func task2() {
	arr := []int{5, 3, 8, 4, 1}
	fmt.Println("Исходный массив:", arr)
	sortedArr := sortArray(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}

func sortArray(arr []int) []int {
	sort.Ints(arr)
	return arr
}

// 3. Сумма квадратов чётных чисел
func task3() {
	var n int
	fmt.Println("Введите число n:")
	fmt.Scanf("%d\n", &n)
	result := sumOfSquares(n)
	fmt.Printf("Сумма квадратов чётных чисел от 1 до %d: %d\n", n, result)
}

func sumOfSquares(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i * i
	}
	return sum
}

// 4. Проверка палиндрома
func task4() {
	var s string
	fmt.Println("Введите строку для проверки на палиндром:")
	fmt.Scanf("%s\n", &s)
	if isPalindrome(s) {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}

func isPalindrome(s string) bool {
	// Приводим строку к единому регистру
	s = strings.ToLower(s)
	// Сравниваем строку с её перевёрнутой версией
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// 5. Проверка числа на простоту
func task5() {
	var n int
	fmt.Println("Введите число для проверки на простоту:")
	fmt.Scanf("%d\n", &n)
	if isPrime(n) {
		fmt.Println("Число простое.")
	} else {
		fmt.Println("Число не простое.")
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 6. Генерация последовательности простых чисел
func task6() {
	var limit int
	fmt.Println("Введите предел для генерации простых чисел:")
	fmt.Scanf("%d\n", &limit)
	primes := generatePrimes(limit)
	fmt.Println("Простые числа:", primes)
}

func generatePrimes(limit int) []int {
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// 7. Перевод числа в двоичную систему
func task7() {
	var n int
	fmt.Println("Введите число для перевода в двоичную систему:")
	fmt.Scanf("%d\n", &n)
	binary := toBinary(n)
	fmt.Println("Двоичная запись:", binary)
}

func toBinary(n int) string {
	return fmt.Sprintf("%b", n)
}

// 8. Нахождение максимального элемента в массиве
func task8() {
	arr := []int{5, 8, 2, 9, 3}
	fmt.Println("Массив:", arr)
	max := findMax(arr)
	fmt.Println("Максимальный элемент массива:", max)
}

func findMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// 9. Функция вычисления НОД
func task9() {
	var a, b int
	fmt.Println("Введите два числа для вычисления НОД:")
	fmt.Scanf("%d %d\n", &a, &b)
	result := gcd(a, b)
	fmt.Printf("НОД чисел %d и %d: %d\n", a, b, result)
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 10. Сумма элементов массива
func task10() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Массив:", arr)
	sum := sumArray(arr)
	fmt.Println("Сумма элементов массива:", sum)
}

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	var taskId int
	fmt.Println("Введите номер задачи")
	fmt.Scanf("%d\n", &taskId)
	switch taskId {
	case 1:
		task1()
	case 2:
		task2()
	case 3:
		task3()
	case 4:
		task4()
	case 5:
		task5()
	case 6:
		task6()
	case 7:
		task7()
	case 8:
		task8()
	case 9:
		task9()
	case 10:
		task10()
	default:
		fmt.Println("Задача не найдена")
	}
}
