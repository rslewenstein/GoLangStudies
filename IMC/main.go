package main

import (
	"fmt"
	"math"
)

func VerificaImc(peso float64, altura float64) string{
	var msg string
	var imcTotal = peso / (altura * altura)

	if imcTotal < 18.5{
		msg = fmt.Sprint(math.Floor(imcTotal * 100)/100) + " Está abaixo do peso"
	}else if imcTotal > 18.6 && imcTotal <= 24.9 {
		msg = fmt.Sprint(math.Floor(imcTotal * 100)/100) + " Peso ideal"
	}else if imcTotal >= 25.0 && imcTotal <= 29.9 {
		msg = fmt.Sprint(math.Floor(imcTotal * 100)/100) + " Levemente acima do peso"
	}else if imcTotal >= 30.0 && imcTotal <= 34.9 {
		msg = fmt.Sprint(math.Floor(imcTotal * 100)/100) + " Obesidade grau I"
	}else if imcTotal >= 35.0 && imcTotal <= 39.9 {
		msg = fmt.Sprint(math.Floor(imcTotal * 100)/100) + " Obesidade grau II"
	}else{
		msg = fmt.Sprint(math.Floor(imcTotal * 100)/100) + " Obesidade grau III"
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