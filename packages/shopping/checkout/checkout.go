package checkout

func CalculateTotal(prices []*float64) float64 {
	var total float64
	for i := 0; i < len(prices); i++ {
		total += *prices[i]
	}
	return total
}
