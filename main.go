package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Интерфейс с функциями для структуры
type Printable interface {
	AddStudent()
	LoadFromFile()
	RemoveStudent()
	PrintStudents()
	SaveToFile()
	SortByAverageGrade()
}

// Структура студента
type Student struct {
	Name   string
	Age    int
	Grades float64
}

// Добавить студента
func AddStudent(students []Student) []Student {
	var name string
	var age int
	var grades float64
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	fmt.Print("Введите средний балл: ")
	fmt.Scan(&grades)
	NewStudents := Student{
		Name:   name,
		Age:    age,
		Grades: grades,
	}

	students = append(students, NewStudents)
	return students

}
func RemoveStudent(students []Student, name string) []Student {
	NewList := []Student{}
	for _, s := range students {
		if s.Name != name {
			NewList = append(NewList, s)
		}
	}
	return NewList
}

// Ввывести список студентов
func (students Student) PrintStudents() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d, Средний бал:%.1f\n", students.Name, students.Age, students.Grades)
}

// Сортировка по возрастанию
func SortByAverageGrade(students []Student) []Student {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Grades > students[j].Grades
	})
	return students
}

// Сохранить данные в файле
func SaveToFile(students []Student) {

	filename, err := os.Create("Student.txt")
	if err != nil {
		fmt.Println("Ошибка при открытие:", err)
		return
	}
	defer filename.Close()

	for _, s := range students {
		_, err := filename.WriteString(s.PrintStudents() + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи:", err)
			return
		}
	}
	return
}
func LoadFromFile() []Student {
	var students []Student
	data, err := os.ReadFile("Student.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return nil
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Пример строки: "Имя: Илья, Возраст: 23, Средний бал:4.5"
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			fmt.Println("Некорректная строка:", line)
			continue
		}

		// Убираем префиксы и пробелы
		name := strings.TrimSpace(strings.Replace(parts[0], "Имя:", "", 1))
		ageStr := strings.TrimSpace(strings.Replace(parts[1], "Возраст:", "", 1))
		gradeStr := strings.TrimSpace(strings.Replace(parts[2], "Средний бал:", "", 1))

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Println("Ошибка при парсинге возраста:", err)
			continue
		}

		grades, err := strconv.ParseFloat(gradeStr, 64)
		if err != nil {
			fmt.Println("Ошибка при парсинге оценки:", err)
			continue
		}

		students = append(students, Student{
			Name:   name,
			Age:    age,
			Grades: grades,
		})
	}
	return students

}

func main() {
	var Action, name string
	students := []Student{
		{Name: "Илья", Age: 23, Grades: 4.5},
		{Name: "Павел", Age: 23, Grades: 3.7},
	}

	for Action != "7" {
		fmt.Println("Выберите номер действия со списком студентов:(пример: 1 (добавить нового студента), если хотите выйти напишите: exit")
		fmt.Println("1) Добавить нового студента")
		fmt.Println("2) Удалить студента")
		fmt.Println("3) Вывести список студентов")
		fmt.Println("4) Сортировать по убыванию оценок")
		fmt.Println("5) Сохранить список в файл")
		fmt.Println("6) Загрузить из файла")
		fmt.Println("7) Выход")
		fmt.Scan(&Action)

		switch Action {
		case "1":
			students = AddStudent(students)
		case "2":
			fmt.Println("Укажите имя студента которого надо удалить: ")
			fmt.Scan(&name)
			students = RemoveStudent(students, name)
			fmt.Println("Студент удален")
		case "3":
			for _, s := range students {
				fmt.Println(s.PrintStudents())
			}
		case "4":
			students = SortByAverageGrade(students)
		case "5":
			SaveToFile(students)
			fmt.Println("Данные сохранины")
		case "6":
			students = LoadFromFile()
		case "7":
			fmt.Println("Программа окончила работу ")
		default:
			fmt.Println("Вы не правильно выбрали действие")
		}
	}
}
