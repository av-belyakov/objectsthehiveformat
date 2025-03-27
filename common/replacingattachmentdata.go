package common

import (
	"reflect"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ReplacingOldValues заменяет старые значения структуры AttachmentData
// новыми значениями. Изменяемые поля:
// Size - размер
// Id - идентификатор
// Name - наименование
// ContentType - тип контента
// Hashes - список хешей
func (a *AttachmentData) ReplacingOldValues(element AttachmentData) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(a).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "Hashes"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Hashes" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			//обработка полей содержащихся в ObservableMessageEs
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

	return countReplacingFields
}
