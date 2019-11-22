package proteinProvision

var haveFood map[string]float64 = map[string]float64{
	"upper" : 1.5,
	"floor" : 1.2,
}

var notFood map[string]float64 = map[string]float64{
	"upper" : 1.2,
	"floor" : 1,
}

func CalcProteinProvide(food bool, weight float64) map[string]float64 {
	var proteinProvision map[string]float64 = map[string]float64{
		"upper" : 0,
		"floor" : 0,
	}
	if food {
		proteinProvision["upper"] = weight * haveFood["upper"]
		proteinProvision["floor"] = weight * haveFood["floor"]
 	} else {
		proteinProvision["upper"] = weight * notFood["upper"]
		proteinProvision["floor"] = weight * notFood["floor"]
	}

	return proteinProvision
}