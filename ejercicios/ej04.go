package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
// func Invertir(cadena string) string {
// 	if len(cadena) == 1 {
// 		return cadena
// 	}
// 	if len(cadena) == 0 {
// 		return ""
// 	}

// 	firstLetter := string(cadena[0])
// 	lastLetter := string(cadena[len(cadena)-1])
// 	midLetters := string(cadena[1 : len(cadena)-1])
// 	return lastLetter + Invertir(midLetters) + firstLetter
// }

func Invertir(cadena string) string {
	if len(cadena) <= 1 {
		return cadena
	}

	return Invertir(cadena[1:]) + cadena[:1]
}
