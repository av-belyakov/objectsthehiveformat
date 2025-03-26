package common

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

//****************** CommonEventCase ******************

func (e *CommonEventType) Get() *CommonEventType {
	return e
}

// GetBase поле Base
func (e *CommonEventType) GetBase() bool {
	return e.Base
}

// SetValueBase поле Base
func (e *CommonEventType) SetValueBase(v bool) {
	e.Base = v
}

// SetAnyBase поле Base
func (e *CommonEventType) SetAnyBase(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Base = v
	}
}

// GetStartDate значение в формате RFC3339 для поля StartDate
func (e *CommonEventType) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate значение в формате RFC3339 для поля StartDate
func (e *CommonEventType) SetValueStartDate(v string) {
	e.StartDate = v
}

// SetAnyStartDate значение для поля StartDate
func (e *CommonEventType) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetRootId основной идентификатор
func (e *CommonEventType) GetRootId() string {
	return e.RootId
}

// SetValueRootId основной идентификатор
func (e *CommonEventType) SetValueRootId(v string) {
	e.RootId = v
}

// SetAnyRootId основной идентификатор
func (e *CommonEventType) SetAnyRootId(i interface{}) {
	e.RootId = fmt.Sprint(i)
}

// GetOrganisation наименование организации
func (e *CommonEventType) GetOrganisation() string {
	return e.Organisation
}

// SetValueOrganisation наименование организации
func (e *CommonEventType) SetValueOrganisation(v string) {
	e.Organisation = v
}

// SetAnyOrganisation наименование организации
func (e *CommonEventType) SetAnyOrganisation(i interface{}) {
	e.Organisation = fmt.Sprint(i)
}

// GetOrganisationId идентификатор организации
func (e *CommonEventType) GetOrganisationId() string {
	return e.OrganisationId
}

// SetValueOrganisationId идентификатор организации
func (e *CommonEventType) SetValueOrganisationId(v string) {
	e.OrganisationId = v
}

// SetAnyOrganisationId идентификатор организации
func (e *CommonEventType) SetAnyOrganisationId(i interface{}) {
	e.OrganisationId = fmt.Sprint(i)
}

// GetObjectId идентификатор объекта
func (e *CommonEventType) GetObjectId() string {
	return e.ObjectId
}

// SetValueObjectId идентификатор объекта
func (e *CommonEventType) SetValueObjectId(v string) {
	e.ObjectId = v
}

// SetAnyObjectId идентификатор объекта
func (e *CommonEventType) SetAnyObjectId(i interface{}) {
	e.ObjectId = fmt.Sprint(i)
}

// GetObjectType тип объекта
func (e *CommonEventType) GetObjectType() string {
	return e.ObjectType
}

// SetValueObjectType тип объекта
func (e *CommonEventType) SetValueObjectType(v string) {
	e.ObjectType = v
}

// SetAnyObjectType тип объекта
func (e *CommonEventType) SetAnyObjectType(i interface{}) {
	e.ObjectType = fmt.Sprint(i)
}

// GetOperation операция
func (e *CommonEventType) GetOperation() string {
	return e.Operation
}

// SetValueOperation операция
func (e *CommonEventType) SetValueOperation(v string) {
	e.Operation = v
}

// SetAnyOperation операция
func (e *CommonEventType) SetAnyOperation(i interface{}) {
	e.Operation = fmt.Sprint(i)
}

// GetRequestId идентификатор запроса
func (e *CommonEventType) GetRequestId() string {
	return e.RequestId
}

// SetValueRequestId идентификатор запроса
func (e *CommonEventType) SetValueRequestId(v string) {
	e.RequestId = v
}

// SetAnyRequestId идентификатор запроса
func (e *CommonEventType) SetAnyRequestId(i interface{}) {
	e.RequestId = fmt.Sprint(i)
}

// ToStringBeautiful форматированный вывод
func (em CommonEventType) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'operation': '%s'\n", ws, em.Operation))
	str.WriteString(fmt.Sprintf("%s'objectId': '%s'\n", ws, em.ObjectId))
	str.WriteString(fmt.Sprintf("%s'objectType': '%s'\n", ws, em.ObjectType))
	str.WriteString(fmt.Sprintf("%s'base': '%v'\n", ws, em.Base))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, em.StartDate))
	str.WriteString(fmt.Sprintf("%s'rootId': '%s'\n", ws, em.RootId))
	str.WriteString(fmt.Sprintf("%s'requestId': '%s'\n", ws, em.RequestId))
	str.WriteString(fmt.Sprintf("%s'organisationId': '%s'\n", ws, em.OrganisationId))
	str.WriteString(fmt.Sprintf("%s'organisation': '%s'\n", ws, em.Organisation))

	return str.String()
}
