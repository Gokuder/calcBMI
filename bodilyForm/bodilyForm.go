package bodilyForm

var bodilyFormRule map[string]map[string]float64 = map[string]map[string]float64{
	"loseFlesh" : map[string]float64{
		"upper" : 18.4,
		"floor" : 0,
	},
	"normal" : map[string]float64{
		"upper" : 23.9,
		"floor" : 18.5,
	},
	"fat" : map[string]float64{
		"floor" : 24,
	},
}

func CalcBodilyFormByBMI(BMI float64) string  {
	for bodilyForm,bodilyFormValue := range bodilyFormRule {
		if bodilyFormValue["upper"] !=0 {
			if BMI >= bodilyFormValue["floor"] && BMI <= bodilyFormValue["upper"] {
				return bodilyForm
			}
		}else{
			if BMI >= bodilyFormValue["floor"]  {
				return bodilyForm
			}
		}
	}
	var bodilyForm string
	return bodilyForm
}