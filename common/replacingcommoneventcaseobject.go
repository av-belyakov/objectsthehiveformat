package common

import "reflect"

// ReplacingOldValues заменяет старые значения структуры CommonEventCaseObject
// новыми значениями. Изменяемые поля:
// Flag - флаг
// CaseId - уникальный идентификатор дела
// Severity - строгость
// Tlp - tlp
// Pap - pap
// StartDate - начальная дата
// EndDate - конечная дата
// CreatedAt - дата создания
// UpdatedAt - дата обновления
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// CreatedBy - кем создан
// UpdatedBy - кем обновлен
// UnderliningType - тип
// Title - заголовок
// Description - описание
// ImpactStatus - краткое описание воздействия
// ResolutionStatus - статус разрешения
// Status - статус
// Summary - резюме
// Owner - владелец
func (e *CommonEventCaseObject) ReplacingOldValues(element CommonEventCaseObject) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(e).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
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

					//не обновлять значение если оно соответствует пустой дате
					if str == "1970-01-01T00:00:00+00:00" {
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
