package alert

import (
	"maps"
	"reflect"

	"github.com/av-belyakov/objectsthehiveformat/alertartifacts"
	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ReplacingOldValues заменяет старые значения структуры TypeAlert
// новыми значениями. Изменяемые поля:
// Tags - теги
// TagsAll - все теги
// CustomFields - настраиваемые поля
// Artifacts - артефакты
func (a *TypeAlert) ReplacingOldValues(element TypeAlert) (int, error) {
	var (
		err                  error
		countReplacingFields int
	)

	currentStruct := reflect.ValueOf(a).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки общих полей
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "CommonAlertType" {
				countReplacingFields += a.CommonEventAlertObject.ReplacingOldValues(*element.CommonEventAlertObject.Get())

				continue
			}

			// для обработки поля "Tags"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				newTags, okNew := newStruct.Field(j).Interface().(map[string][]string)
				if !okNew {
					continue
				}

				for key, value := range newTags {
					for _, v := range value {
						if a.SetValueTags(key, v) {
							countReplacingFields++
						}
					}
				}

				continue
			}

			// для обработки поля "TagsAll"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "TagsAll" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "CustomFields"
			if typeOfCurrentStruct.Field(i).Name == "CustomFields" {
				currentCustomFields, okCurr := currentStruct.Field(i).Interface().(common.CustomFields)
				newCustomFields, okNew := newStruct.Field(j).Interface().(common.CustomFields)
				if !okCurr || !okNew {
					continue
				}

				if currentCustomFields == nil {
					currentCustomFields = make(common.CustomFields)
				}

				maps.Copy(currentCustomFields, newCustomFields)
				a.SetValueCustomFields(currentCustomFields)
				countReplacingFields++

				continue
			}

			// для обработки поля "Artifacts"
			if typeOfCurrentStruct.Field(i).Name == "Artifacts" {
				newArtifact, okNew := newStruct.Field(j).Interface().(map[string][]alertartifacts.Artifacts)
				if !okNew {
					continue
				}

				for key, value := range newArtifact {
					currentArtifacts, ok := a.GetKeyArtifacts(key)
					if !ok {
						a.SetKeyArtifacts(key, value)

						continue
					}

					modifiedArtifacts, num := comparisonListsArtifacts(currentArtifacts, value)
					countReplacingFields += num
					a.SetKeyArtifacts(key, modifiedArtifacts)
				}

				continue
			}

			//обработка полей содержащихся в AlertMessageForEsAlert
			//и не относящихся к вышеперечисленым значениям
			//*******************************************************
			if !currentStruct.Field(i).Equal(newStruct.Field(j)) {
				if !currentStruct.Field(i).CanSet() {
					continue
				}

				if str, ok := newStruct.Field(j).Interface().(string); ok {
					//не обновлять текущие значения новыми пустыми значениями
					if str == "" {
						continue
					}
				}

				currentStruct.Field(i).Set(newStruct.Field(j))
				countReplacingFields++
			}
		}
	}

	return countReplacingFields, err
}

// comparisonListsArtifacts объединяет два списка
func comparisonListsArtifacts(currentArtifacts, newArtifacts []alertartifacts.Artifacts) ([]alertartifacts.Artifacts, int) {
	var countReplacingFields int

	for _, value := range newArtifacts {
		var isExist bool

		for k, v := range currentArtifacts {
			if value.GetId() == v.GetId() || value.GetUnderliningId() == v.GetUnderliningId() {
				isExist = true
				countReplacingFields += currentArtifacts[k].ReplacingOldValues(value)

				break
			}
		}

		if !isExist {
			currentArtifacts = append(currentArtifacts, value)
		}
	}

	return currentArtifacts, countReplacingFields
}
