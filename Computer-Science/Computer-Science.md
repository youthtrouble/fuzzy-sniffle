# Why it is a bad idea to use recursion to find the fibonacci of a number.

```go
...

func fibo(nthNumber int) int {
	if nthNumber <= 1 {
		return nthNumber
	} 
	return fibo(nthNumber - 1) + fibo(nthNumber - 2)
}
```
Taking the function above as an example, this is the natural way of writing the Fibonacci series because according to the definition a Fibonacci number is 1 if 'nthNumber' is 0 or 1. Else it is f(nthNumber-1) + f(nthNumber-2) and the above code represents this in the most natural way. But the major problem with this is that the same value will be calculated again and again for calculation of a Fibonacci term and this will take more time and more importantly, more stack space. 