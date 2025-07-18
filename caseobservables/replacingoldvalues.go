package caseobservables

import (
	"reflect"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ReplacingOldValues заменяет старые значения структуры Observables
// новыми значениями. Изменяемые поля:
// Observables
func (o *Observables) ReplacingOldValues(element Observables) int {
	var countReplacingFields int

	for key, value := range element.Observables {
		currentObservables, ok := o.GetKeyObservables(key)
		if !ok {
			o.SetKeyObservables(key, value)

			continue
		}

		modifiedObservables, num := comparisonListsObservables(currentObservables, value)
		countReplacingFields += num
		o.SetKeyObservables(key, modifiedObservables)
	}

	return countReplacingFields
}

// ReplacingOldValues заменяет старые значения структуры Observable
// новыми значениями. Изменяемые поля:
// SensorId - идентификатор сенсора
// SnortSid - список идентификаторов сигнатур
// Tags - список тегов
// TagsAll - список всех тегов
// Attachment - приложенные данные
// Reports - список отчетов
func (o *Observable) ReplacingOldValues(element Observable) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(o).Elem()
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
			if typeOfCurrentStruct.Field(i).Name == "CommonObservableType" {
				countReplacingFields += o.CommonObservableType.Get().ReplacingOldValues(*element.CommonObservableType.Get())
			}

			// для обработки поля "SnortSid"
			//******************************
			if typeOfCurrentStruct.Field(i).Name == "SnortSid" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			//для обработки поля "SnortSidNumber"
			//***********************************
			if typeOfCurrentStruct.Field(i).Name == "SnortSidNumber" {
				if list, ok := supportingfunctions.ReplacingSlice[int](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

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
						if o.SetValueTags(key, v) {
							countReplacingFields++
						}
					}
				}

				continue
			}

			// для обработки поля "TagsAll"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "TagsAll" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "Attachment"
			//********************************
			if typeOfCurrentStruct.Field(i).Name == "Attachment" {
				countReplacingFields += o.Attachment.ReplacingOldValues(*element.GetAttachment())

				continue
			}

			//обработка полей содержащихся в ObservableMessageEs
			//и не относящихся к вышеперечисленым значениям
			//***************************************************
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

	return countReplacingFields
}
