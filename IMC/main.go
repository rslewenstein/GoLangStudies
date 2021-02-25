package main

import (
	"fmt"
)

func VerificaImc(peso float64, altura float64) string{
	var msg string
	var imcTotal = peso / (altura * altura)

	if imcTotal < 18.5{
		msg = fmt.Sprint(imcTotal) + " Está abaixo do peso"
	}else if imcTotal > 18.6 && imcTotal <= 24.9 {
		msg = fmt.Sprint(imcTotal) + " Peso ideal"
	}else if imcTotal >= 25.0 && imcTotal <= 29.9 {
		msg = fmt.Sprint(imcTotal) + " Levemente acima do peso"
	}else if imcTotal >= 30.0 && imcTotal <= 34.9 {
		msg = fmt.Sprint(imcTotal) + " Obesidade grau I"
	}else if imcTotal >= 35.0 && imcTotal <= 39.9 {
		msg = fmt.Sprint(imcTotal) + " Obesidade grau II"
	}else{
		msg = fmt.Sprint(imcTotal) + " Obesidade grau III"
	}	

	return msg
}

func main(){
	var peso float64
	var altura float64

	fmt.Print("Digite o peso: ")
	fmt.Scanln(&peso)

	fmt.Print("Digite a altura: ")
	fmt.Scanln(&altura)

	fmt.Println(VerificaImc(peso, altura))
}