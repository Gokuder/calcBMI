package heat

// 计算公式的标量
const ScalarHeat float64 = 0.7

var heatRule map[string]map[string]map[string]float64 = map[string]map[string]map[string]float64{
	"veryMild": map[string]map[string]float64{
		"loseFlesh": map[string]float64{
			"upper": 30,
			"floor": 25,
		},
		"normal": map[string]float64{
			"upper": 25,
			"floor": 20,
		},
		"fat": map[string]float64{
			"upper": 20,
			"floor": 15,
		},
	},
	"mild": map[string]map[string]float64{
		"loseFlesh": map[string]float64{
			"upper": 35,
			"floor": 30,
		},
		"normal": map[string]float64{
			"upper": 30,
			"floor": 25,
		},
		"fat": map[string]float64{
			"upper": 25,
			"floor": 20,
		},
	},
	"moderate": map[string]map[string]float64{
		"loseFlesh": map[string]float64{
			"upper": 40,
			"floor": 35,
		},
		"normal": map[string]float64{
			"upper": 35,
			"floor": 30,
		},
		"fat": map[string]float64{
			"upper": 25,
			"floor": 20,
		},
	},
	"severe": map[string]map[string]float64{
		"loseFlesh": map[string]float64{
			"upper": 45,
			"floor": 40,
		},
		"normal": map[string]float64{
			"upper": 40,
			"floor": 35,
		},
		"fat": map[string]float64{
			"upper": 35,
			"floor": 30,
		},
	},
}
var ruleKeyForActiveFlag[4]string = [4]string{
	"veryMild",
	"mild",
	"moderate",
	"severe",
}

func CalcHeatProvideScope(weight float64,activeFlag uint8, bodilyForm string) map[string]float64 {
	ruleKey := ruleKeyForActiveFlag[activeFlag]
	upper := weight * ScalarHeat * heatRule[ruleKey][bodilyForm]["upper"]
	floor := weight * ScalarHeat * heatRule[ruleKey][bodilyForm]["floor"]

	heat := map[string]float64 {
		"upper" : upper,
		"floor" : floor,
	}

	return heat
}