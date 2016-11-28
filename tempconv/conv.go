package tempconv

// CtoF converts a Celcius temperature to a Fahrenheit one
func CtoF (c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
// FtoC converts a Fahrenheit temparature to a Celcius one
func FtoC (f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}
