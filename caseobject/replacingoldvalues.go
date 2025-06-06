package caseobject

import (
	"maps"
	"reflect"

	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ReplacingOldValues заменяет старые значения структуры EventForEsCaseObject
// новыми значениями. Изменяемые поля:
// Tags - список тегов
// TagsAll - список всех тегов
// CustomFields - настраиваемые поля
func (o *EventCaseObject) ReplacingOldValues(element EventCaseObject) int {
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
			if typeOfCurrentStruct.Field(i).Name == "CommonEventCaseObject" {
				countReplacingFields += o.CommonEventCaseObject.Get().ReplacingOldValues(*element.CommonEventCaseObject.Get())
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

			//для обработки поля "CustomFields"
			//*********************************
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
				o.SetValueCustomFields(currentCustomFields)
				countReplacingFields++

				continue
			}
		}
	}

	return countReplacingFields
}
