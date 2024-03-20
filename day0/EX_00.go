package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readNumbers() (numbers []float32, mean float32, mode float32) {
	scanner := bufio.NewScanner(os.Stdin)
	frequency := make(map[float32]int)
	minFreq := math.MinInt
	maxKey := float32(math.MaxFloat32)
	next := true

	for next && scanner.Scan() {
		input := scanner.Text()

		parts := strings.Fields(input)

		for _, part := range parts {
			tmp, err := strconv.Atoi(part)
			num := float32(tmp)
			if num < -100000.0 || num > 100000.0 || err != nil {
				next = false
				break
			}
			numbers = append(numbers, num)
			frequency[num]++                               // Считаем частотность числа
			mean += num                                    // Суммируем все числа для рассчета среднего значения
			if frequency[num] > minFreq || num <= maxKey { // Обновляем Mode если его частотность больше минимального или оно меньше предыдущего числа такой же частоты
				minFreq = frequency[num]
				maxKey = num
			}
		}
	}
	if len(numbers) == 0 {
		os.Exit(1)
	}
	mean /= float32(len(numbers))
	mode = maxKey
	return numbers, mean, mode
}

func main() {
	var flags int = 15                                      // Все поля включены по умолчанию
	headerTable := []string{"Mean", "Median", "Mode", "SD"} // Заголовки таблицы
	valueHeaderTable := make([]float32, 4)                  // Переменные для таблицы
	if len(os.Args) == 2 {
		flags, _ = strconv.Atoi(os.Args[1])
	}
	var numbers []float32 // Последовательность считанных чисел
	numbers, valueHeaderTable[0], valueHeaderTable[2] = readNumbers()
	slices.Sort(numbers)

	length := len(numbers)
	if length%2 == 1 { // Если нечетный то медиана это элемент по середине
		valueHeaderTable[1] = numbers[length/2]
	} else { // иначе высчитываем среднее из середину и его соседа
		valueHeaderTable[1] = numbers[length/2] + numbers[length/2-1]/2.0
	}

	for i := range numbers { // Считаем дисперсию
		tmp := numbers[i] - valueHeaderTable[0]
		valueHeaderTable[3] += tmp * tmp
	}
	valueHeaderTable[3] /= float32(length)

	valueHeaderTable[3] = float32(math.Sqrt(float64(valueHeaderTable[3]))) // Среднеквадратичное отклонение

	mask := 1
	for i := 0; i < 4; i++ {
		if (flags & mask) > 0 {
			if i != 2 {
				fmt.Printf("%v: %.2f\n", headerTable[i], valueHeaderTable[i])
			} else {
				fmt.Printf("%v: %.0f\n", headerTable[i], valueHeaderTable[i])
			}
		}
		mask <<= 1
	}
}
