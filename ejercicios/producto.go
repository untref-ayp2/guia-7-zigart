package ejercicios

// Escriba un método recursivo que tome un entero n
// devuelva su factorial
func Producto(n int, m int) int {
	if n == 0 || m == 0 {
		return 0
	}

	return n + Producto(n, m-1)

}
