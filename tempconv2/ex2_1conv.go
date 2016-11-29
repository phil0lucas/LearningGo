package tempconv2

/*
	F -> C -> K
	K -> C -> F
*/

// FtoC converts a Fahrenheit temparature to a Celcius one
func FtoC (f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}

// CtoK converts a Celsius temperature to Kelvin
func CtoK (c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FtoK convert Fahrenheit to Kelvin
func FtoK (f Fahrenheit) Kelvin {
	return Kelvin(CtoK(FtoC(f)))
}

// ---------------------------------------------------------

// KtoC converts a Kelvin temperature to Celsius
func KtoC (k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CtoF converts a Celcius temperature to Fahrenheit
func CtoF (c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// KtoF - Kelvin to Fahrenheit
func KtoF (k Kelvin) Fahrenheit {
	return Fahrenheit(CtoF(KtoC(k)))
}
