package main

import "fmt"

func fibonacci(n int) {
	if n <= 0 {
		fmt.Println("Liczba powinna być większa od 0")
		return
	} else {
		switch n {
		case 1:
			fmt.Println(0)
		case 2:
			fmt.Println(0)
			fmt.Println(1)
		default:
			fn := 0
			f1 := 0
			f2 := 1
			for i := 0; i < n; i++ {
				switch i {
				case 0:
					fn = 0
				case 1:
					fn = 1
				default:

					fn = f1 + f2
					f1 = f2
					f2 = fn
				}
				fmt.Println(fn)

			}

		}
	}

}

func fibonacci2(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fibonacci2(n-2) + fibonacci2(n-1)
	}
}

func fibonacci_print(n int) {
	for i := 1; i <= n; i += 1 {
		println(fibonacci2(i))
	}
}

func main() {

	fibonacci(8)
	fibonacci_print(8)
	var n int
	fmt.Print("Proszę podać liczbę: ")
	fmt.Scan(&n)
	fibonacci(n)
}
