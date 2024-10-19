/*
	package main

import (

	"fmt"
	"math"

)

	type Circle struct {
		radius float64
	}

	type Rectangle struct {
		width  float64
		height float64
	}

	type Shape interface {
		Area() float64
	}

	func (c Circle) Area() float64 {
		a := math.Pi * c.radius * c.radius
		return a
	}

func (r Rectangle) Area() float64 {

		a := r.width * r.height
		return a
	}

	func main() {
		figure_1 := Circle{radius: 1.0}
		fmt.Println(figure_1.Area())
		figure_2 := Rectangle{width: 57.2, height: 10.2}
		fmt.Println(figure_2.Area())
	}

package main

import "fmt"

	type Animal interface {
		MakeSound()
	}

type Dog struct {
}
type Cat struct {
}

	func (d Dog) MakeSound() {
		fmt.Println("Гав!")
	}

	func (c Cat) MakeSound() {
		fmt.Println("Мяу!")
	}

	func main() {
		cat1 := Cat{}
		cat1.MakeSound()
		dog1 := Dog{}
		dog1.MakeSound()
	}

package main

import "fmt"

type LogLevel string

	type Logger interface {
		Log(message string)
	}

const (

	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"

)

	type Log struct {
		Level LogLevel
	}

func (l Log) Log(message string) {

	switch l.Level {
	case Error:
		fmt.Printf("%s: %s\n", Error, message)
	case Info:
		fmt.Printf("%s: %s\n", Info, message)
	}

}

	func main() {
		errorLog := &Log{Level: Error}
		errorLog.Log("This is an error message")
	}

package main

import "fmt"

func ConcatStringsAndInt(str1, str2 string, num int) string {

	return str1 + str2 + fmt.Sprint(num)

}

	func main() {
		fmt.Println(ConcatStringsAndInt("al", "himik", 239))
	}

package main

import "fmt"

	func DivideIntegers(a, b int) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		a1 := float64(a)
		b1 := float64(b)

		return a1 / b1, nil
	}

	func main() {
		fmt.Println(DivideIntegers(10, 2))
	}

package main

import "fmt"

func GetCharacterAtPosition(str string, position int) (rune, error) {

		runes := []rune(str)
		if len(runes)-1 < position {
			return 0, fmt.Errorf("position out of range")
		}

		recover()
		return []rune(str)[position], nil
	}

	func main() {
		fmt.Println(GetCharacterAtPosition("Оле ола вперёд Спартак Москва", 0))
	}

package main

import "fmt"

	func Factorial(n int) (int, error) {
		p := 1
		if n < 0 {
			return 0, fmt.Errorf("factorial is not defined for negative numbers")
		}
		for i := 1; i <= n; i++ {
			p *= i
		}
		return p, nil
	}

package main

import (

	"fmt"
	"strconv"

)

	func IntToBinary(num int) (string, error) {
		p := ""
		if num < 0 {
			return "", fmt.Errorf("negative numbers are not allowed")
		}
		for {
			if num <= 0 {
				break
			}
			f := strconv.Itoa(num % 2)
			p = f + p
			num = num / 2
		}
		return p, nil

}

	func main() {
		fmt.Println(IntToBinary(234))
	}

package main

import (

	"fmt"
	"strconv"
	"strings"

)

	func SumTwoIntegers(a, b string) (int, error) {
		if strings.Contains(a, ".")

}

	func main() {
		fmt.Println(SumTwoIntegers("1.0", "2"))
	}

package main

import (

	"fmt"
	"time"

)

	func main() {
		now := time.Now()
		fmt.Println("Current time:", now)
	}

package main

import (

	"fmt"
	"time"

)

	func main() {
		t1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
		t2 := time.Now()
		diff := t2.Sub(t1)
		fmt.Println("Difference between t1 and t2:", diff)
	}

package main

import (

	"fmt"
	"time"

)

	func main() {
		now := time.Now()

		fmt.Println("1. Время в формате RFC3339:", now.Format(time.RFC3339))
		fmt.Println("2. Полная дата и время:", now.Format("2006-01-02 15:04:05"))
		fmt.Println("3. Краткая дата:", now.Format("2006-01-02"))
		fmt.Println("4. Время в 24-часовом формате:", now.Format("15:04:05"))
		fmt.Println("5. Время в 12-часовом формате с AM/PM:", now.Format("03:04 PM"))
		fmt.Println("6. День недели:", now.Format("Monday"))
		fmt.Println("7. Сокращённый месяц:", now.Format("Jan"))
	}

package main

import (

	"fmt"
	"time"

)

	func main() {
		i := time.NewTimer(5 * time.Second) // Создаётся таймер  на 5 секунд
		duration := 5 * time.Second
		fmt.Println("Duration:", duration, i)

		// Выполняем действия в течение заданного промежутка времени
		fmt.Println("Start")
		time.Sleep(duration)
		fmt.Println("End")

		// Вычисляем разницу между моментами времени
		t1 := time.Now()
		time.Sleep(duration)
		t2 := time.Now()
		diff := t2.Sub(t1)
		fmt.Println("Difference:", diff)
	}

package main

import (

	"fmt"
	"time"

)

	func TimeDifference(start, end time.Time) time.Duration {
		return end.Sub(start)

}

	func main() {
		start := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
		end := time.Date(2023, 10, 23, 4, 41, 49, 0, time.UTC)
		diff := TimeDifference(start, end)
		fmt.Println(diff) // Output: 2h0m0s
	}

package main

import (

	"fmt"
	"time"

)

	func FormatTimeToString(timestamp time.Time, format string) string {
		return timestamp.Format(format)
	}

func main() {

		timestamp := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
		format := "2006-01-02 15:04:05"
		result := FormatTimeToString(timestamp, format)
		fmt.Println(result)
	} // Output: 2023-10-23 02:41:49

package main

import (

	"fmt"
	"time"

)

	func ParseStringToTime(dateString, format string) (time.Time, error) {
		pt, err := time.Parse(format, dateString)
		if err != nil {
			return pt, err
		}
		return pt, nil
	}

	func main() {
		dateString := "2023-10-23 02:41:49"
		format := "2006-01-02 15:04:05"
		result, err := ParseStringToTime(dateString, format)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	} // Output: 2023-10-23 02:41:49 +0000 UTC

package main

import (
	"fmt"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	t1 := time.Now()
	d := t1.Sub(pastTime)
	switch {
	case d < time.Minute:
		return fmt.Sprintf("%d seconds ago", int(d.Seconds()))
	case d < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(d.Minutes()))
	case d < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(d.Hours()))
	case d < 24*30*time.Hour:
		return fmt.Sprintf("%d days ago", int(d.Hours())/24)
	case d < 24*30*12*time.Hour:
		return fmt.Sprintf("%d month ago", int(d.Hours())/24/30)
	default:
		y := int(d.Hours()) / (24 * 365)
		return fmt.Sprintf("%d years ago", y)
	}

}

func main() {
	pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
	result := TimeAgo(pastTime)
	fmt.Println(result) // Output: 1 month ago
}
*/

package main

import (
	"time"
)

func NextWorkday(start time.Time) time.Time {
	nd := start.AddDate(0, 0, 1)
	if nd.Weekday() == time.Saturday {
		nd = start.AddDate(0, 0, 2)
	} else if nd.Weekday() == time.Sunday {
		nd = start.AddDate(0, 0, 1)

	} else if nd.Weekday() == time.Friday {
		nd = start.AddDate(0, 0, 3)

	}
	return nd
}
