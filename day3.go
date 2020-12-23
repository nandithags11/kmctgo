package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Enter x and y:")
	fmt.Scanf("%d %d", &x, &y)
	fmt.Println("x>y =>", x > y)
	fmt.Println("x<y =>", x < y)
	fmt.Println("x>=y =>", x >= y)
	fmt.Println("x<=y =>", x <= y)
	fmt.Println("x==y =>", x == y)
	fmt.Println("x!=y =>", x != y)
	fmt.Println("x>y && x==y =>", (x > y && x == y))
	fmt.Println("(!(x<=y)) =>", (!(x <= y)))
	fmt.Println("(x==y || x<y) =>", (x == y || x < y))

}
