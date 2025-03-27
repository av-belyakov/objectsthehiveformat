package caseobservables

// comparisonListsObservables объединяет два списка
func comparisonListsObservables(currentObs, newObs []Observable) ([]Observable, int) {
	var countReplacingFields int

	for _, value := range newObs {
		var isExist bool

		for k, v := range currentObs {
			if value.GetUnderliningId() == v.GetUnderliningId() {
				isExist = true
				countReplacingFields += currentObs[k].ReplacingOldValues(value)

				break
			}
		}

		if !isExist {
			currentObs = append(currentObs, value)
		}
	}

	return currentObs, countReplacingFields
}
