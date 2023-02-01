package data_models

func convertJouleValuesToKilowattHour(jouleValues []float64) []float64 {
	var kilowattHourValues []float64
	for _, value := range jouleValues {
		value = value * 0.2_777_777 * 0.000_00_1
		kilowattHourValues = append(kilowattHourValues, value)
	}
	return kilowattHourValues
}
