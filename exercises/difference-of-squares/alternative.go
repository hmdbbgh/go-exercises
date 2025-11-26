package diffsquares

// SquareOfSumLoop computes the square of sum via an explicit loop.
func SquareOfSumLoop(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum * sum
}

// SumOfSquaresLoop computes the sum of squares via an explicit loop.
func SumOfSquaresLoop(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i * i
    }
    return sum
}

// DifferenceLoop uses the loop-based versions.
func DifferenceLoop(n int) int {
    return SquareOfSumLoop(n) - SumOfSquaresLoop(n)
}
