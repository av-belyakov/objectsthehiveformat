package casedetails

import (
	"reflect"

	"github.com/av-belyakov/objectsthehiveformat/common"
)

// ReplacingOldValues заменяет старые значения структуры EventCaseDetails
// новыми значениями. Изменяемые поля:
// SourceRef - ссылка
// Title - заголовок
// Description - описание
// GeoIp - сюда помещаются теги относящиеся к географическому позиционированию
// SensorId - сюда помещаются теги относящиеся к номерам сигнатур
// Reasons - сюда помещаются теги относящиеся к причине возникновения события
func (d *EventCaseDetails) ReplacingOldValues(element EventCaseDetails) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(d).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			//для обработки поля "CustomFields"
			//*********************************
			if typeOfCurrentStruct.Field(i).Name == "CustomFields" {
				currentCustomFields, okCurr := currentStruct.Field(i).Interface().(common.CustomFields)
				newCustomFields, okNew := newStruct.Field(j).Interface().(common.CustomFields)
				if !okCurr || !okNew {
					continue
				}

				for k, v := range newCustomFields {
					currentCustomFields[k] = v
				}

				d.SetValueCustomFields(currentCustomFields)
				countReplacingFields++

				continue
			}

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
