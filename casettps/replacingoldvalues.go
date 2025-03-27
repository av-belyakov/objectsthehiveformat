package casettps

import (
	"reflect"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

//***************** Ttps ********************

// ReplacingOldValues заменяет старые значения структуры Ttps
// новыми значениями. Изменяемые поля:
// Ttp
func (ttp *Ttps) ReplacingOldValues(element Ttps) int {
	modifiedTtp, num := comparisonListsTtp(ttp.Ttp, element.Ttp)
	ttp.SetTtps(modifiedTtp)

	return num
}

//***************** Ttp ********************

// ReplacingOldValues заменяет старые значения структуры Ttp
// новыми значениями. Изменяемые поля:
// OccurDate - дата возникновения
// UnderliningCreatedAt - время создания
// UnderliningId - уникальный идентификатор
// UnderliningCreatedBy - кем создан
// PatternId - уникальный идентификатор шаблона
// Tactic - тактика
// ExtraData - дополнительные данные
func (a *Ttp) ReplacingOldValues(element Ttp) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(a).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "ExtraData"
			if typeOfCurrentStruct.Field(i).Name == "ExtraData" {
				countReplacingFields += a.ExtraData.ReplacingOldValues(element.GetExtraData())
				continue
			}

			//обработка полей содержащихся в TtpMessage
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

//***************** ExtraDataTtpMessage ********************

// ReplacingOldValues заменяет старые значения структуры ExtraDataTtpMessage
// новыми значениями. Изменяемые поля:
// Pattern - шаблон
// PatternParent - родительский шаблон
func (ed *ExtraDataTtp) ReplacingOldValues(element ExtraDataTtp) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(ed).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "Pattern"
			if typeOfCurrentStruct.Field(i).Name == "Pattern" {
				countReplacingFields += ed.Pattern.ReplacingOldValues(element.GetPattern())
				continue
			}

			// для обработки поля "PatternParent"
			if typeOfCurrentStruct.Field(i).Name == "PatternParent" {
				countReplacingFields += ed.PatternParent.ReplacingOldValues(element.GetPatternParent())
				continue
			}
		}
	}

	return countReplacingFields
}

//***************** PatternExtraData ********************

// ReplacingOldValues заменяет старые значения структуры PatternExtraData
// новыми значениями. Изменяемые поля:
// RemoteSupport - удаленная поддержка
// Revoked - аннулированный
// UnderliningCreatedAt - время создания
// UnderliningCreatedBy - кем создан
// UnderliningId - уникальный идентификатор
// UnderliningType - тип
// DataSources - источники данных
// DefenseBypassed - чем выполнен обход защиты
// Description - описание
// ExtraData - дополнительные данные
// Name - наименование
// PatternId - уникальный идентификатор шаблона
// PatternType - тип шаблона
// PermissionsRequired - требуемые разрешения
// Platforms - список платформ
// SystemRequirements - системные требования
// Tactics - список тактик
// URL - URL
// Version - версия
func (ped *PatternExtraData) ReplacingOldValues(element PatternExtraData) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(ped).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "Platforms"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "Platforms" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "PermissionsRequired"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "PermissionsRequired" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "DataSources"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "DataSources" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "Tactics"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "Tactics" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			//обработка полей содержащихся в PatternExtraData
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
