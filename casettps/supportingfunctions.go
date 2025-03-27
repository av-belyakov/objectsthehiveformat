package casettps

// comparisonListsTtp объединяет два списка
func comparisonListsTtp(currentTtp, newTtp []Ttp) ([]Ttp, int) {
	var countReplacingFields int

	for _, value := range newTtp {
		var isExist bool

		for k, v := range currentTtp {
			if value.GetUnderliningId() == v.GetUnderliningId() {
				isExist = true
				countReplacingFields += currentTtp[k].ReplacingOldValues(value)

				break
			}
		}

		if !isExist {
			currentTtp = append(currentTtp, value)
		}
	}

	return currentTtp, countReplacingFields
}
