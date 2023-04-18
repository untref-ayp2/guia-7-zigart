package ejercicios

// Escriba un método recursivo que tome un entero n
// devuelva su factorial
func Division(a, b int) int {
	if a < b {
		return 0
	}

	return 1 + Division(a-b, b)
}
