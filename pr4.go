package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func task_1_1() {
	var num int
	fmt.Println("Введи число?")
	fmt.Scanf("%d\n", &num)
	// для 4 символов
	res := 0
	res += num % 10
	num = num / 10
	res += num % 10
	num = num / 10
	res += num % 10
	num = num / 10
	res += num % 10
	num = num / 10
	fmt.Println(res)

	// через цикл
	//res = 0
	//for num2 != 0 {
	//	res += num2 % 10
	//	num2 = num2 / 10
	//}
	fmt.Println(res)
}

func task_1_2() {
	var temp float64
	var unit string

	fmt.Print("Введите температуру и единицу измерения (C для Цельсия, F для Фаренгейта): ")
	fmt.Scan(&temp, &unit)

	switch unit {
	case "C", "c":
		// Преобразование из Цельсия в Фаренгейты
		fahrenheit := temp*9/5 + 32
		fmt.Printf("%.2f (Celsius) = %.2f (Fahrenheit)\n", temp, fahrenheit)
	case "F", "f":
		// Преобразование из Фаренгейтов в Цельсий
		celsius := (temp - 32) * 5 / 9
		fmt.Printf("%.2f (Fahrenheit) = %.2f (Celsius)\n", temp, celsius)
	default:
		fmt.Println("Некорректная единица измерения. Используйте 'C' для Цельсия или 'F' для Фаренгейта.")
	}
}

func doubleArray(arr []int, index int) []int {
	if index >= len(arr) {
		return arr
	}
	arr[index] *= 2
	return doubleArray(arr, index+1)
}

func task_1_3() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	doubledArr := doubleArray(arr, 0)
	fmt.Println("Массив с удвоенными элементами:", doubledArr)
}

func task_1_4() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите строки через пробел:")
	input, _ := reader.ReadString('\n')

	// Убираем символ новой строки в конце
	input = strings.TrimSpace(input)

	joinedString := strings.Join(strings.Fields(input), " ")

	fmt.Println("Объединенная строка:", joinedString)
}

type Point struct {
	x float64
	y float64
}

func (p Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func task_1_5() {

	p1 := Point{x: 0, y: 0}
	p2 := Point{x: 0, y: 0}
	fmt.Println("Введите x и y в формате: 10 10")
	fmt.Scanf("%f %f", &p1.x, &p1.y)
	fmt.Println("Введите x и y в формате: 10 10")
	fmt.Scanf("%f %f", &p2.x, &p2.y)

	distance := p1.DistanceTo(&p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

func task_2_1() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

func task_2_2() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}
}

func task_2_3() {
	var a, b, c int
	fmt.Print("Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	fmt.Println("Наибольшее число:", max)
}

func task_2_4() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	if age >= 0 && age <= 12 {
		fmt.Println("Ребенок")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Подросток")
	} else if age >= 18 && age <= 64 {
		fmt.Println("Взрослый")
	} else if age >= 65 {
		fmt.Println("Пожилой")
	} else {
		fmt.Println("Некорректный возраст")
	}
}

func task_2_5() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}

func task_3_1() {
	var n int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&n)

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("Факториал числа:", factorial)
}

func task_3_2() {
	var n int
	fmt.Print("Введите количество чисел Фибоначчи: ")
	fmt.Scan(&n)

	fib := make([]int, n)
	if n > 0 {
		fib[0] = 0
	}
	if n > 1 {
		fib[1] = 1
	}

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println("Числа Фибоначчи:", fib)
}

func read_array() *[]int {
	var arr []int
	var elem, n int
	fmt.Print("Введите количество чисел в масиве: ")
	fmt.Scan(&n)
	fmt.Println("Введите исходный массив:")
	for i := 0; i < n; i++ {
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}
	return &arr
}

func task_3_3() {
	arr := *read_array()
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println("Перевернутый массив:", arr)
}

func task_3_4() {
	var n int
	fmt.Print("Введите число для поиска простых чисел до n: ")
	fmt.Scan(&n)

	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	fmt.Println("Простые числа до", n, ":", primes)
}

func task_3_5() {
	arr := *read_array()
	sum := 0
	for _, num := range arr {
		sum += num
	}

	fmt.Println("Сумма чисел в массиве:", sum)
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
		fmt.Println("%d задача не найдена", taskId)

	}
}
