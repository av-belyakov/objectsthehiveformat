package common

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewCommonEventAlertObject создает новый объект типа CommonEventAlertObject
func NewCommonEventAlertObject() *CommonEventAlertObject {
	return &CommonEventAlertObject{
		CreatedAt: "1970-01-01T00:00:00+00:00",
		UpdatedAt: "1970-01-01T00:00:00+00:00",
	}
}

func (o *CommonEventAlertObject) Get() *CommonEventAlertObject {
	return o
}

func (o *CommonEventAlertObject) GetTlp() uint64 {
	return o.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (o *CommonEventAlertObject) SetValueTlp(v uint64) {
	o.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (o *CommonEventAlertObject) SetAnyTlp(i any) {
	if v, ok := i.(float32); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		o.Tlp = v
	}
}

func (o *CommonEventAlertObject) GetUnderliningId() string {
	return o.UnderliningId
}

// SetValueUnderliningId значение поля UnderliningId
func (o *CommonEventAlertObject) SetValueUnderliningId(v string) {
	o.UnderliningId = v
}

// SetAnyUnderliningId значение поля UnderliningId
func (o *CommonEventAlertObject) SetAnyUnderliningId(i any) {
	o.UnderliningId = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetId() string {
	return o.Id
}

// SetValueId значение поля Id
func (o *CommonEventAlertObject) SetValueId(v string) {
	o.Id = v
}

// SetAnyId значение поля Id
func (o *CommonEventAlertObject) SetAnyId(i any) {
	o.Id = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetCreatedBy() string {
	return o.CreatedBy
}

// SetValueCreatedBy значение поля CreatedBy
func (o *CommonEventAlertObject) SetValueCreatedBy(v string) {
	o.CreatedBy = v
}

// SetAnyCreatedBy значение поля CreatedBy
func (o *CommonEventAlertObject) SetAnyCreatedBy(i any) {
	o.CreatedBy = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetUpdatedBy() string {
	return o.UpdatedBy
}

// SetValueUpdatedBy значение поля UpdatedBy
func (o *CommonEventAlertObject) SetValueUpdatedBy(v string) {
	o.UpdatedBy = v
}

// SetAnyUpdatedBy значение поля UpdatedBy
func (o *CommonEventAlertObject) SetAnyUpdatedBy(i any) {
	o.UpdatedBy = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetCreatedAt() string {
	return o.CreatedAt
}

// SetValueCreatedAt значение в формате RFC3339 поля CreatedAt
func (o *CommonEventAlertObject) SetValueCreatedAt(v string) {
	o.CreatedAt = v
}

// SetAnyCreatedAt значение поля CreatedAt
func (o *CommonEventAlertObject) SetAnyCreatedAt(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonEventAlertObject) GetUpdatedAt() string {
	return o.UpdatedAt
}

// SetValueUpdatedAt значение  в формате RFC3339 поля UpdatedAt
func (o *CommonEventAlertObject) SetValueUpdatedAt(v string) {
	o.UpdatedAt = v
}

// SetAnyUpdatedAtзначение поля UpdatedAt
func (o *CommonEventAlertObject) SetAnyUpdatedAt(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonEventAlertObject) GetUnderliningType() string {
	return o.UnderliningType
}

// SetValueUnderliningType значение поля UnderliningType
func (o *CommonEventAlertObject) SetValueUnderliningType(v string) {
	o.UnderliningType = v
}

// SetAnyUnderliningType значение поля UnderliningType
func (o *CommonEventAlertObject) SetAnyUnderliningType(i any) {
	o.UnderliningType = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetTitle() string {
	return o.Title
}

// SetValueTitle значение поля Title
func (o *CommonEventAlertObject) SetValueTitle(v string) {
	o.Title = v
}

// SetAnyTitle значение поля Title
func (o *CommonEventAlertObject) SetAnyTitle(i any) {
	o.Title = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetDescription() string {
	return o.Description
}

// SetValueDescription значение поля Description
func (o *CommonEventAlertObject) SetValueDescription(v string) {
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	o.Description = v
}

// SetAnyDescription значение поля Description
func (o *CommonEventAlertObject) SetAnyDescription(i any) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	o.Description = str
}

func (o *CommonEventAlertObject) GetStatus() string {
	return o.Status
}

// SetValueStatus значение поля Status
func (o *CommonEventAlertObject) SetValueStatus(v string) {
	o.Status = v
}

// SetAnyStatus значение поля Status
func (o *CommonEventAlertObject) SetAnyStatus(i any) {
	o.Status = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetDate() string {
	return o.Date
}

// SetValueDate значение в формате RFC3339 поля Date
func (o *CommonEventAlertObject) SetValueDate(v string) {
	o.Date = v
}

// SetAnyDate значение поля Date
func (o *CommonEventAlertObject) SetAnyDate(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonEventAlertObject) GetType() string {
	return o.Type
}

// SetValueType значение поля Type
func (o *CommonEventAlertObject) SetValueType(v string) {
	o.Type = v
}

// SetAnyType значение поля Type
func (o *CommonEventAlertObject) SetAnyType(i any) {
	o.Type = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetObjectType() string {
	return o.ObjectType
}

// SetValueObjectType значение поля ObjectType
func (o *CommonEventAlertObject) SetValueObjectType(v string) {
	o.ObjectType = v
}

// SetAnyObjectType значение поля ObjectType
func (o *CommonEventAlertObject) SetAnyObjectType(i any) {
	o.ObjectType = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetSource() string {
	return o.Source
}

// SetValueSource значение поля Source
func (o *CommonEventAlertObject) SetValueSource(v string) {
	o.Source = v
}

// SetAnySource значение поля Source
func (o *CommonEventAlertObject) SetAnySource(i any) {
	o.Source = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetSourceRef() string {
	return o.SourceRef
}

// SetValueSourceRef значение поля SourceRef
func (o *CommonEventAlertObject) SetValueSourceRef(v string) {
	o.SourceRef = v
}

// SetAnySourceRef значение поля SourceRef
func (o *CommonEventAlertObject) SetAnySourceRef(i any) {
	o.SourceRef = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetCase() string {
	return o.Case
}

// SetValueCase значение поля Case
func (o *CommonEventAlertObject) SetValueCase(v string) {
	o.Case = v
}

// SetAnyCase значение поля Case
func (o *CommonEventAlertObject) SetAnyCase(i any) {
	o.Case = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetCaseTemplate() string {
	return o.CaseTemplate
}

// SetValueCaseTemplate значение поля CaseTemplate
func (o *CommonEventAlertObject) SetValueCaseTemplate(v string) {
	o.CaseTemplate = v
}

// SetAnyCaseTemplate значение поля CaseTemplate
func (o *CommonEventAlertObject) SetAnyCaseTemplate(i any) {
	o.CaseTemplate = fmt.Sprint(i)
}

// ToStringBeautiful форматированный вывод
func (o *CommonEventAlertObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, o.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, o.Id))
	str.WriteString(fmt.Sprintf("%s'createdBy': '%s'\n", ws, o.CreatedBy))
	str.WriteString(fmt.Sprintf("%s'updatedBy': '%s'\n", ws, o.UpdatedBy))
	str.WriteString(fmt.Sprintf("%s'createdAt': '%s'\n", ws, o.CreatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedAt': '%s'\n", ws, o.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, o.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, o.Tlp))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, o.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, o.Description))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, o.Status))
	str.WriteString(fmt.Sprintf("%s'date': '%s'\n", ws, o.Date))
	str.WriteString(fmt.Sprintf("%s'type': '%s'\n", ws, o.Type))
	str.WriteString(fmt.Sprintf("%s'objectType': '%s'\n", ws, o.ObjectType))
	str.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, o.Source))
	str.WriteString(fmt.Sprintf("%s'sourceRef': '%s'\n", ws, o.SourceRef))
	str.WriteString(fmt.Sprintf("%s'case': '%s'\n", ws, o.Case))
	str.WriteString(fmt.Sprintf("%s'caseTemplate': '%s'\n", ws, o.CaseTemplate))

	return str.String()
}
