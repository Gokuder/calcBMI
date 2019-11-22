package BMI

import "math"

func CalcBMI(weight float64, height uint8) float64 {
	var BMI float64 = weight / math.Pow(float64(height), 2)
	return BMI*1000
}
