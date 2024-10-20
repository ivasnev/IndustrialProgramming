package main

import (
	bufio "bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Чтение массива
func read_array() *[]int {
	var arr []int
	var elem, n int
	fmt.Print("Введите количество чисел в массиве: ")
	fmt.Scan(&n)
	fmt.Println("Введите исходный массив:")
	for i := 0; i < n; i++ {
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}
	return &arr
}

// 1. Проверка на простоту
func task_1() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanf("%d", &num)
	if num < 2 {
		fmt.Println("Число не является простым")
		return
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("Простое")
}

// 2. НОД (алгоритм Евклида)
func task_2() {
	var a, b int
	fmt.Println("Введите два числа через пробел")
	fmt.Scanf("%d %d", &a, &b)
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println("НОД:", a)
}

// 3. Сортировка пузырьком
func task_3() {
	arr := *read_array()
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println("Шаг", i+1, ":", arr)
	}
	fmt.Println("Отсортированный массив:", arr)
}

// 4. Таблица умножения 10x10
func task_4() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

// 5. Фибоначчи с мемоизацией
var memo = map[int]int{}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func task_5() {
	var n int
	fmt.Println("Введите число для вычисления Фибоначчи")
	fmt.Scan(&n)
	fmt.Println("Число Фибоначчи:", fibonacci(n))
}

// 6. Обратные числа
func task_6() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	fmt.Println("Обратное число:", reversed)
}

// 7. Треугольник Паскаля
func task_7() {
	var levels int
	fmt.Println("Введите количество уровней")
	fmt.Scan(&levels)
	triangle := make([][]int, levels)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		fmt.Println(triangle[i])
	}
}

// 8. Число палиндром
func task_8() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)
	original, reversed := num, 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	if original == reversed {
		fmt.Println("Число является палиндромом")
	} else {
		fmt.Println("Число не является палиндромом")
	}
}

// 9. Минимум и максимум в массиве
func task_9() {
	arr := *read_array()
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println("Минимум:", min, "Максимум:", max)
}

// 10. Игра "Угадай число"
func task_10() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 0
	var guess int

	for attempts < 10 {
		fmt.Print("Введите число: ")
		fmt.Scan(&guess)
		attempts++
		if guess < target {
			fmt.Println("Больше!")
		} else if guess > target {
			fmt.Println("Меньше!")
		} else {
			fmt.Println("Вы угадали!")
			return
		}
	}
	fmt.Println("Вы не угадали, было загадано число:", target)
}

// 11. Число Армстронга
func task_11() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)
	original, sum, n := num, 0, 0

	for temp := num; temp > 0; temp /= 10 {
		n++
	}
	for temp := num; temp > 0; temp /= 10 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(n)))
	}
	if sum == original {
		fmt.Println("Число является числом Армстронга")
	} else {
		fmt.Println("Число не является числом Армстронга")
	}
}

// 12. Подсчёт слов в строке
func task_12() {
	var input string
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	// Приводим строку к нижнему регистру и разбиваем на слова
	words := strings.Fields(strings.ToLower(input))

	// Создаем map для хранения слов и их количества
	wordCount := make(map[string]int)

	// Подсчитываем каждое слово
	for _, word := range words {
		wordCount[word]++
	}

	// Выводим количество уникальных слов и их частоту
	fmt.Println("Количество уникальных слов и их частота:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// Размер поля
const rows, cols = 10, 10

// Инициализация случайного поля
func initializeGrid() [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2) // 0 или 1 (мертвая или живая клетка)
		}
	}
	return grid
}

// Вывод поля
func printGrid(grid [][]int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Подсчет живых соседей
func countNeighbors(grid [][]int, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			ni, nj := x+i, y+j
			if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
				count += grid[ni][nj]
			}
		}
	}
	return count
}

// Следующее поколение
func nextGeneration(grid [][]int) [][]int {
	newGrid := make([][]int, rows)
	for i := range newGrid {
		newGrid[i] = make([]int, cols)
		for j := range newGrid[i] {
			liveNeighbors := countNeighbors(grid, i, j)
			if grid[i][j] == 1 && (liveNeighbors == 2 || liveNeighbors == 3) {
				newGrid[i][j] = 1 // Остается живой
			} else if grid[i][j] == 0 && liveNeighbors == 3 {
				newGrid[i][j] = 1 // Оживает
			} else {
				newGrid[i][j] = 0 // Умирает или остается мертвой
			}
		}
	}
	return newGrid
}

func task_13() {
	var num int
	fmt.Println("Введите количество генераций:")
	fmt.Scan(&num)
	rand.Seed(time.Now().UnixNano())
	grid := initializeGrid()

	fmt.Println("Начальное состояние:")
	printGrid(grid)

	for generation := 1; generation <= num; generation++ { // 5 поколений
		time.Sleep(1 * time.Second) // Задержка для наглядности
		fmt.Printf("Поколение %d:\n", generation)
		grid = nextGeneration(grid)
		printGrid(grid)
	}
}

// 14. Цифровой корень числа
func task_14() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)
	for num >= 10 {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
	}
	fmt.Println("Цифровой корень:", num)
}

// 15. Римские цифры
func task_15() {
	var num int
	fmt.Println("Введите арабское число")
	fmt.Scan(&num)
	roman := []struct {
		value  int
		symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	result := ""
	for _, r := range roman {
		for num >= r.value {
			result += r.symbol
			num -= r.value
		}
	}
	fmt.Println("Римская запись:", result)
}

func main() {
	var taskId int
	fmt.Println("Введите номер задачи")
	fmt.Scanf("%d\n", &taskId)
	switch taskId {
	case 1:
		task_1()
	case 2:
		task_2()
	case 3:
		task_3()
	case 4:
		task_4()
	case 5:
		task_5()
	case 6:
		task_6()
	case 7:
		task_7()
	case 8:
		task_8()
	case 9:
		task_9()
	case 10:
		task_10()
	case 11:
		task_11()
	case 12:
		task_12()
	case 13:
		task_13()
	case 14:
		task_14()
	case 15:
		task_15()
	default:
		fmt.Println("Задача не найдена")
	}
}
