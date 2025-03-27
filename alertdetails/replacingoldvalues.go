package alertdetails

import (
	"reflect"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ReplacingOldValues заменяет старые значения структуры EventAlertDetails
// новыми значениями. Изменяемые поля:
// SourceRef - ссылка
// Title - заголовок
// Description - описание
// GeoIp - сюда помещаются теги относящиеся к географическому позиционированию
// SensorId - сюда помещаются теги относящиеся к номерам сигнатур
// Reasons - сюда помещаются теги относящиеся к причине возникновения события
func (d *EventAlertDetails) ReplacingOldValues(element EventAlertDetails) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(d).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
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
						if d.SetValueTags(key, v) {
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
