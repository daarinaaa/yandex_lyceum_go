/*
package main

import (

	"fmt"

)

	func main() {
		var age int
		fmt.Scanln(&age)
		if age < 18 {
			fmt.Println("Вы несовершеннолетний")
		} else if age < 65 {
			fmt.Println("Вы взрослый")
		} else {
			fmt.Println("Вы очень взрослый")
		}
	}

	func main() {
		var x, y int
		fmt.Scanln(&x)
		fmt.Scanln(&y)
		if x > 0 && y > 0 {
			fmt.Println("больше 0")
		}
	}

	func main() {
		weather := "sunny"
		switch weather {
		case "rain":
			fmt.Println("Погода дождливая")
		case "sunny":
			fmt.Println("Погода солнечная")
		default:
			fmt.Println("Погода не определена")
		}
	}

	func main() {
		number := 5
		if square := number * number; square > 50 {
			fmt.Println("Квадрат числа больше 50")
		} else {
			fmt.Println("Квадрат числа меньше или равен 50")
		}
	}

package main

import (

	"fmt"

)

	func main() {
		var x int
		fmt.Scanln(&x)
		if x < 0 {
			fmt.Println("Некорректный ввод")
			return
		}
		switch {
		case (x < 10):
			fmt.Println("Число меньше 10")
		case (x >= 10) && (x < 100):
			fmt.Println("Число меньше 100")
		case (x <= 100) || (x < 1000):
			fmt.Println("Число меньше 1000")
		default:
			fmt.Println("Число больше или равно 1000")
		}
	}

package main

import (

	"fmt"

)

	func main() {
		for i := range 10 {
			fmt.Println(i)
		}

		for i, letter := range "Hello, world!" {
			fmt.Println(i, string(letter))
		}

		for _, letter := range "Hello, world!" {
			fmt.Println(string(letter))
		}

}

package main

import "fmt"

	func main() {
		for {
			var input string
			fmt.Scanln(&input)
			if input == "exit" {
				break
			}
		}
	}

package main

import "fmt"

	func main() {
		var x int
		fmt.Scanln(&x)
		a := 0
		if x < 1 {
			fmt.Println("Некорректный ввод")
			return
		}
		for i := 1; i <= x; i += 1 {
			if i%3 == 0 || i%5 == 0 {
				continue
			}
			a += i
		}
		fmt.Println(a)

}

package main

import (

	"fmt"
	"math"

)

// название         аргументы с типом

	func findDiscriminant(a, b, c float64) {
		// тело функции
		discriminant := math.Pow(b, 2) - 4.0*a*c
		fmt.Println(discriminant)
	}

	func main() {
		findDiscriminant(1, -3.0, 2) // a = 1  b = -3.0  c = 2
		// Вывод: 1

		findDiscriminant(1, 4.0, 2) // a = 1  b = 4.0   c = 2
		// Вывод: 8
	}

package main

import (

	"fmt"

)

	func factorial(n int) int {
		// крайний случай
		if n == 0 {
			return 1
		}

		// вызов функции с текущим значением, умноженным на результат (n - 1)
		return n * factorial(n-1)
	}

	func main() {
		result := factorial(5)
		fmt.Println(result)
	}

package main

import (

	"fmt"
	"math"

)

	func SqRoots() {
		var a, b, c float64
		fmt.Scanln(&a, &b, &c)

		d := math.Pow(b, 2) - 4*a*c
		//fmt.Println(d)

		switch {
		case d > 0:
			fmt.Println((-b-math.Pow(d, 0.5))/(2*a), (-b+math.Pow(d, 0.5))/(2*a))
		case d == 0:
			fmt.Println((-b) / (2 * a))
		case d < 0:
			fmt.Println(0, 0)
		}

}

	func main() {
		SqRoots()
	}

package main

import (

	"fmt"

)

	func Add(a, b float64) float64 {
		return a + b
	}

	func Multiply(a, b float64) float64 {
		return a * b
	}

	func PrintNumbersAscending(n int) {
		for i := 1; i <= n; i++ {
			fmt.Println(i)
		}
	}

	func main() {
		fmt.Println(Add(5, 10))
		fmt.Println(Multiply(5, 9))
		PrintNumbersAscending(5)
	}

package main

import (

	"fmt"

)

	func SumDigitsRecursive(n int) int {
		if n == 0 {
			return 0
		}

		return n%10 + SumDigitsRecursive(n/10)
	}

	func main() {
		fmt.Println(SumDigitsRecursive(123))
	}

package main

import (

	"fmt"

)

	func IsPowerOfTwoRecursive(N int) {
		a := 1
		for i := 1; i <= 20; i++ {
			if a == N {
				fmt.Println("YES")
				return
			}
			a *= 2
		}
		fmt.Println("NO")
	}

	func main() {
		IsPowerOfTwoRecursive(64)
	}

package main

import (

	"fmt"

)

	func CalculateSeriesSum(n int) float64 {
		if n == 0 {
			return 0
		}

		return 1/float64(n) + CalculateSeriesSum(n-1)
	}

	func main() {
		fmt.Println(CalculateSeriesSum(6))
	}

package main

import "fmt"

	func FiveSteps(array [5]int) [5]int {
		var a1 [5]int
		n := 0
		for i := len(array) - 1; i >= 0; i-- {
			a1[n] = array[i]
			n++
		}
		return a1
	}

	func main() {
		input := [5]int{1, 2, 3, 4, 5}
		output := FiveSteps(input)
		for i := 0; i < len(output); i++ {
			fmt.Printf("%d ", output[i])
		}
	}

package main

import "fmt"

	func ThirdElementInArray(array [6]int) int {
		return array[2]
	}

	func main() {
		a := [6]int{1, 2, 3, 4, 5, 6}
		fmt.Println(ThirdElementInArray(a))
	}

package main

import "fmt"

	func FindMinMaxInArray(array [10]int) (int, int) {
		m1 := -100000
		m2 := 100000000
		for i := 0; i < len(array); i++ {
			if array[i] > m1 {
				m1 = array[i]
			}
			if array[i] < m2 {
				m2 = array[i]
			}
		}
		return m1, m2
	}

	func main() {
		input := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		max, min := FindMinMaxInArray(input)

		fmt.Print(max, min)
	}

package main

import "fmt"

	func SumOfArray(array [6]int) int {
		s := 0
		for i := 0; i < len(array); i++ {
			s += array[i]
		}
		return s
	}

	func main() {
		a := [6]int{1, 2, 3, 4, 5, 6}
		fmt.Println(SumOfArray(a))
	}

package main

import "fmt"

	func PrettyArrayOutput(array [9]string) {
		n := 1
		for i := 0; i <= 6; i++ {
			fmt.Println(n, "я уже сделал:", array[i])
			n++
		}
		for i := 7; i < len(array); i++ {
			fmt.Println(n, "я не успел сделать:", array[i])
			n++
		}
	}

	func main() {
		input := [9]string{
			"проснуться",
			"позавтракать",
			"сходить в школу",
			"пообедать",
			"погулять с друзьями",
			"сделать домашнюю работу",
			"попрограммировать на Go",
			"поужинать",
			"лечь спать",
		}
		PrettyArrayOutput(input)
	}

package main

import "fmt"

	func main() {
		a := [5]int{0, 1, 1, 2, 3} // объявляем массив a из пяти элементов
		b := a[:]                  // в переменную b присваиваем слайс всего массива
		c := b[:]                  // в переменную c присваиваем слайс всего слайса

		fmt.Println(a) // [42 1 1 2 3]
		fmt.Println(b) // [42 1 1 2 3]
		c[0] = 42

		fmt.Println(a) // [42 1 1 2 3]
		fmt.Println(b) // [42 1 1 2 3]
		fmt.Println(c) // [42 1 1 2 3]
	}

package main

import (

	"fmt"

)

	func UnderLimit(nums []int, limit int, n int) ([]int, error) {
		if nums == nil {
			return nil, fmt.Errorf("input slice is nil")
		}
		if n < 0 {
			return nil, fmt.Errorf("n cannot be negative")
		}

		a := []int{}

		for i := 0; i < len(nums); i++ {
			if len(a) >= n {
				break
			}
			if nums[i] < limit {
				a = append(a, nums[i])
			}
		}
		return a, nil
	}

package main

import "fmt"

	func Clean(nums []int, x int) []int {
		w := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] != x {
				nums[w] = nums[i]
				w++
			}
		}
		return nums[:w]
	}

	func main() {
		a := []int{0, 1, 1, 2, 3, 1, 6}
		fmt.Println(Clean(a, 1))

}

package main

import "fmt"

	func SliceCopy(nums []int) []int {
		b := make([]int, len(nums), len(nums))
		for i := 0; i < len(b); i++ {
			b[i] = nums[i]
		}
		return b

}

	func main() {
		a := make([]int, 3, 6)
		r := (SliceCopy(a))
		fmt.Println(cap(r), len(r))
	}

package main

import "fmt"

	func Join(nums1, nums2 []int) []int {
		a := make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))
		for i := 0; i < len(nums1); i++ {
			a[i] = nums1[i]
		}
		k := len(nums1)
		for i := 0; i < len(nums2); i++ {
			a[k] = nums2[i]
			k++
		}
		return a
	}

	func main() {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6, 7}
		r := (Join(a, b))
		fmt.Println(cap(r), len(r), r)
	}

package main

import "fmt"

	func Mix(nums []int) []int {
		n := len(nums) / 2
		a := []int{}
		a1 := []int{}
		a2 := []int{}
		k1 := 0
		k2 := 0
		for i := 0; i < n; i++ {
			a1 = append(a1, nums[i])
		}
		for i := n; i < len(nums); i++ {
			a2 = append(a2, nums[i])
		}

		for i := 0; i < len(a1)+len(a2); i++ {
			if i%2 == 0 {
				a = append(a, a1[k1])
				k1++
			} else {
				a = append(a, a2[k2])
				k2++

			}

		}
		return a
	}

	func main() {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		r := (Mix(a))
		fmt.Println(r)
	}

package main

import "fmt"

	func StringLength(input string) int{
		return len(input)
	}

package main

	func ConcatenateStrings(str1, str2 string) string {
		return str1 + " " + str2

}

package main

import (

	"fmt"

)

	func isLatin(input string) bool {
		f := 1
		for _, i := range input {
			if i >= 65 && i <= 60 || i >= 97 && i <= 122 {
				f = 1
			} else {
				f = 0
				break
			}

		}
		if f == 1 {
			return true
		} else {
			return false
		}

}

	func main() {
		fmt.Println(isLatin("asd"))
	}

package main

import "fmt"

	func IsPalindrome(input string) bool {
		a := ""
		im := ""
		for i := len(input) - 1; i >= 0; i-- {
			if string(input[i]) != " " {
				a += string(input[i])
			}

		}
		for i := 0; i < len(input); i++ {
			if string(input[i]) != " " {
				im += string(input[i])
			}
		}

		return a == im
	}

	func main() {
		fmt.Println(IsPalindrome(" a b c       cba  "))

}

package main

import (

	"fmt"

)

	func countOccurrences(nums []int) map[int]int {
		occurrences := make(map[int]int)

		for _, num := range nums {
			occurrences[num]++
		}

		return occurrences
	}

	func main() {
		numbers := []int{1, 2, 3, 2, 1, 3, 4, 5, 4, 1}
		result := countOccurrences(numbers)

		fmt.Println("Результат подсчёта повторений:")
		for num, count := range result {
			fmt.Printf("%d встречается %d раз(а)\n", num, count)
		}
	}

package main

import (

	"fmt"

)

	func findStudentByID(id int, students map[int]string) (string, error) {
	    name, found := students[id]
	    if !found {
	        return "", fmt.Errorf("студент с ID %d не найден", id)
	    }

	    return name, nil
	}

	func main() {
	    students := map[int]string{
	        1: "Иванов",
	        2: "Петров",
	        3: "Сидоров",
	    }

	    id := 2
	    name, err := findStudentByID(id, students)
	    if err != nil {
	        fmt.Println(err)
	    } else {
	        fmt.Printf("Студент с ID %d: %s\n", id, name)
	    }
	}

package main

import "fmt"

	func FindMaxKey(m map[int]int) int {
		mx := -100000
		mxi := 0
		for q, i := range m {
			if i > mx {
				mx = i
				mxi = q

			}
		}
		return mxi
	}

	func main() {
		m := map[int]int{
			19: 37,
			20: 38,
		}
		fmt.Println(FindMaxKey(m))
	}

package main

import "fmt"

	func SumOfValuesInMap(m map[int]int) int {
		s := 0
		for _, j := range m {
			s += j
		}
		return s
	}

	func main() {
		m := map[int]int{
			19: 37,
			20: 38,
		}
		fmt.Println(SumOfValuesInMap(m))
	}

package main

import "fmt"

	func SwapKeysAndValues(m map[string]string) map[string]string {
		m1 := make(map[string]string)
		for i, j := range m {
			fmt.Println(i, j)
			m1[j] = i
		}
		return m1
	}

	func main() {
		m := map[string]string{
			"ключ1": "37",
			"ключ2": "38",
		}
		fmt.Println(SwapKeysAndValues(m))
	}

package main

import (

	"fmt"

)

	func CountingSort(contacts []string) map[string]int {
		m := make(map[string]int)
		for _, i := range contacts {
			m[i]++
		}
		return m
	}

	func main() {
		a := []string{"1", "2", "1"}
		fmt.Println(CountingSort(a))
	}

package main

import (

	"fmt"

)

	func DeleteLongKeys(m map[string]int) map[string]int {
		m1 := make((map[string]int))
		for i, j := range m {
			fmt.Println(i, j, len(i))
			if len(i) >= 6 {
				m1[i] = j
			}
		}
		return m1
	}

	func main() {
		m := map[string]int{
			"1111":   37,
			"11":     318,
			"111111": 3,
		}
		fmt.Println(DeleteLongKeys(m))
	}

package main

import (

	"fmt"
	"lesson/students"

)

	func main() {
		student1 := students.Student{}
		fmt.Println(student1)
	}

package main

import "fmt"

	type Person struct {
		name    string
		age     int
		address string
	}

	func (p Person) Print() {
		fmt.Printf("Name: %s\n", p.name)
		fmt.Printf("Age: %d\n", p.age)
		fmt.Printf("Addres: %s", p.address)

}

	func main() {
		man := Person{name: "Anton", age: 31, address: "Krasnogorsk"}
		man.Print()
	}

package main

import "fmt"

	type Employee struct {
		name     string
		position string
		salary   float64
		bonus    float64
	}

	func (e Employee) CalculateTotalSalary() {
		ts := e.bonus + e.salary
		fmt.Printf("Employee: %s\n", e.name)
		fmt.Printf("Position: %s\n", e.position)
		fmt.Printf("Total Salary: %.2f\n", ts)

}

	func main() {
		man := Employee{name: "Anton", position: "product manager", salary: 100.0, bonus: 100.0}
		man.CalculateTotalSalary()
	}

package main

import "fmt"

	type Student struct {
		name            string
		solvedProblems  int
		scoreForOneTask float64
		passingScore    float64
	}

	func (s Student) IsExcellentStudent() bool {
		a := s.solvedProblems * int(s.scoreForOneTask)
		return a >= int(s.passingScore)
	}

	func main() {
		student := Student{name: "Gosha", solvedProblems: 30, scoreForOneTask: 10.0, passingScore: 290.0}
		fmt.Print(student.IsExcellentStudent())
	}
*/
package main

import (
	"time"
)

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t Task) IsOverdue() bool {
	return t.deadline.Before(time.Now())

}

func (t Task) IsTopPriority() bool {
	return t.priority == 4 || t.priority == 5
}

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (td ToDoList) TasksCount() int {
	return len(td.tasks)

}

func (td ToDoList) NotesCount() int {
	return len(td.notes)

}

func (td ToDoList) CountTopPrioritiesTasks() int {
	k := 0
	for _, i := range td.tasks {
		if i.IsTopPriority() {
			k += 1
		}
	}
	return k
}

func (td ToDoList) CountOverdueTasks() int {
	k := 0
	for _, i := range td.tasks {
		if i.IsOverdue() {
			k += 1
		}
	}
	return k
}
