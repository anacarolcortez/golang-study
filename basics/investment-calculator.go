package main

import (
	"fmt"
	"math"
)

func Calculate(investmentAmount, expectedReturnRate, years float64) {
	const inflation float64 = 6.5

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	var futureRealValue = futureValue / math.Pow((1+inflation/100), years)

	fmt.Printf("Seu investimento, em termos reais, somará: US$ %.2f em %.0f anos.\n", futureRealValue, years)
}
