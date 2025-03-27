package objectsthehiveformat

import "reflect"

// ReplacingOldValues заменяет старые значения структуры TypeEventForCase
// новыми значениями. Изменяемые поля:
// CommonEventType
// Details - детальная информация о событии
// Object - объект события
func (e *TypeEventForCase) ReplacingOldValues(element TypeEventForCase) (int, error) {
	var (
		err                  error
		countReplacingFields int
	)

	currentStruct := reflect.ValueOf(e).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

DONE:
	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки общих полей
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "CommonEventType" {
				num, errC := e.CommonEventType.Get().ReplacingOldValues(*element.CommonEventType.Get())
				if errC != nil {
					err = errC

					break DONE
				}

				countReplacingFields += num
			}

			//для обработки поля "Details"
			if typeOfCurrentStruct.Field(i).Name == "Details" {
				countReplacingFields += e.Details.ReplacingOldValues(element.GetDetails())

				continue
			}

			//для обработки поля "Object"
			if typeOfCurrentStruct.Field(i).Name == "Object" {
				countReplacingFields += e.Object.ReplacingOldValues(element.GetObject())

				continue
			}
		}
	}

	return countReplacingFields, err
}
