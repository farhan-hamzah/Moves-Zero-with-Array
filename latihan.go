package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var B[NMAX]int
	var n, i, p int
	fmt.Scan(&n)
	p = 0
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		if A[i] != 0{
			B[p] = A[i]
			p++
		}
	}
	for p < n{
		B[i] = 0
		p++
	}
	for i = 0; i < n; i++ {
		fmt.Print(B[i], " ")
	}
}