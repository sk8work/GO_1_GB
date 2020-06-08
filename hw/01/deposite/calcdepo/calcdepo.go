package calcdepo

// GetDepo func
func GetDepo(startValue, percentage, years float64) float64 {
	getMoneyBack := startValue
	for i := years; i > 0; i-- {
		getMoneyBack += getMoneyBack * percentage
	}
	return getMoneyBack
}
