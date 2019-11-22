package main

import (
	"calc/bodilyForm"
	"calc/BMI"
	"calc/heat"
	"calc/proteinProvision"
	"fmt"
)

// 身高 cm
const HEIGHT uint8 = 176
// 体重 kg
const WEIGHT float64 = 87.20
// 是否断主食
const Food bool = false
// 活动强度标示
const activeFlag uint8 = 1


func main() {
	BMI := BMI.CalcBMI(WEIGHT, HEIGHT)

	bodilyForm := bodilyForm.CalcBodilyFormByBMI(BMI)

	heat := heat.CalcHeatProvideScope(WEIGHT, activeFlag, bodilyForm)
	proteinProvide := proteinProvision.CalcProteinProvide(Food, WEIGHT)

	heatUpper := heat["upper"]
	heatFloor := heat["floor"]
	proteinProvideUpper := proteinProvide["upper"]
	proteinProvideFloor := proteinProvide["floor"]


	fmt.Printf("热量供给上限为%.2f\n", heatUpper)
	fmt.Printf("热量供给下限为%.2f\n", heatFloor)
	fmt.Printf("蛋白质供给上限为%.2f\n", proteinProvideUpper)
	fmt.Printf("蛋白质供给下限为%.2f\n", proteinProvideFloor)
}

