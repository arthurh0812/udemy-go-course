package main

import "fmt"

func main() {
	foo()

	bar("Hi, Grian.")

	s1 := woo("Monneypenny")

	fmt.Println(s1)

	if s2, ok := mouse("Ian", "Fleming"); ok {
		fmt.Println(s2)
	}
}

// syntax
// "func" (receiver) identifier(paramters) (returns) { code... }

func foo() {
	fmt.Println("Hello from foo!")
}

func bar(s string) {
	fmt.Printf("This was the string: %q\n", s)
}

func woo(s string) string {
	return fmt.Sprintf("Hello from woo, %s!", s)
}

func mouse(first, last string) (string, bool) {
	s := fmt.Sprintf(`%s %s says, "Hello!".`, first, last)
	b := true
	return s, b
}
