package main

import "fmt"

func nextclass(from, next *map[string]bool) {

	for student, _ := range *from {
		(*next)[student] = true
		delete(*from, student)
	}

}

func listClass(class map[string]bool) {
	if len(class) == 0 {
		fmt.Println("Bu sınıfta öğrenci yok")
	} else {
		for i, _ := range class {
			fmt.Println(i)
		}
	}
}

func controlAllClasses(classes ...*map[string]bool) {
	for i := 0; i < len(classes); i++ {
		fmt.Printf("%d . sınıf için kontrol \n ", i+1)
		listClass(*classes[i])
	}
}

func main() {
	class1 := map[string]bool{"Ali": true, "Ahmet": true, "Ayşe": true, "Fatma": true}
	class2 := map[string]bool{}
	// class3 := map[string]bool{}
	// class4 := map[string]bool{}
	fmt.Println("Başlangıçta 1.sınıf öğrecileri")
	listClass(class1)

	nextclass(&class1, &class2)
	controlAllClasses(&class1, &class2)

}
