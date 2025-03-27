package common

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

//**************** CommonEventCaseObject ***************

func NewCommonEventCaseObject() *CommonEventCaseObject {
	return &CommonEventCaseObject{
		StartDate: "1970-01-01T00:00:00+00:00",
		EndDate:   "1970-01-01T00:00:00+00:00",
		CreatedAt: "1970-01-01T00:00:00+00:00",
		UpdatedAt: "1970-01-01T00:00:00+00:00",
	}
}

func (e *CommonEventCaseObject) Get() *CommonEventCaseObject {
	return e
}

// GetFlag значение поля Flag
func (e *CommonEventCaseObject) GetFlag() bool {
	return e.Flag
}

// SetValueFlag значение поля Flag
func (e *CommonEventCaseObject) SetValueFlag(v bool) {
	e.Flag = v
}

// SetAnyFlag значение поля Flag
func (e *CommonEventCaseObject) SetAnyFlag(i any) {
	if v, ok := i.(bool); ok {
		e.Flag = v
	}
}

// GetCaseId идентификатор кейса
func (e *CommonEventCaseObject) GetCaseId() uint64 {
	return e.CaseId
}

// SetValueCaseId идентификатор кейса
func (e *CommonEventCaseObject) SetValueCaseId(v uint64) {
	e.CaseId = v
}

// SetAnyCaseId идентификатор кейса
func (e *CommonEventCaseObject) SetAnyCaseId(i any) {
	if v, ok := i.(float32); ok {
		e.CaseId = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		e.CaseId = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.CaseId = v
	}
}

// GetSeverity значение поля Severity
func (e *CommonEventCaseObject) GetSeverity() uint64 {
	return e.Severity
}

// SetValueSeverity значение поля Severity
func (e *CommonEventCaseObject) SetValueSeverity(v uint64) {
	e.Severity = v
}

// SetAnySeverity значение поля Severity
func (e *CommonEventCaseObject) SetAnySeverity(i any) {
	if v, ok := i.(float32); ok {
		e.Severity = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		e.Severity = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Severity = v
	}
}

// GetTlp значение поля Tlp
func (e *CommonEventCaseObject) GetTlp() uint64 {
	return e.Tlp
}

// SetValueTlp значение поля Tlp
func (e *CommonEventCaseObject) SetValueTlp(v uint64) {
	e.Tlp = v
}

// SetAnyTl pзначение поля Tlp
func (e *CommonEventCaseObject) SetAnyTlp(i any) {
	if v, ok := i.(float32); ok {
		e.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		e.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Tlp = v
	}
}

// GetPap значение поля Pap
func (e *CommonEventCaseObject) GetPap() uint64 {
	return e.Pap
}

// SetValuePap значение поля Pap
func (e *CommonEventCaseObject) SetValuePap(v uint64) {
	e.Pap = v
}

// SetAnyPap значение поля Pap
func (e *CommonEventCaseObject) SetAnyPap(i any) {
	if v, ok := i.(float32); ok {
		e.Pap = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		e.Pap = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Pap = v
	}
}

// GetStartDate начальная дата в формате RFC3339
func (e *CommonEventCaseObject) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate начальная дата в формате RFC3339
func (e *CommonEventCaseObject) SetValueStartDate(v string) {
	e.StartDate = v
}

// SetAnyStartDate начальная дата
func (e *CommonEventCaseObject) SetAnyStartDate(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetEndDate конечная дата в формате RFC3339
func (e *CommonEventCaseObject) GetEndDate() string {
	return e.EndDate
}

// SetValueEndDate конечная дата в формате RFC3339
func (e *CommonEventCaseObject) SetValueEndDate(v string) {
	e.EndDate = v
}

// SetAnyEndDate конечная дата
func (e *CommonEventCaseObject) SetAnyEndDate(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.EndDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetCreatedAt время создания в формате RFC3339
func (e *CommonEventCaseObject) GetCreatedAt() string {
	return e.CreatedAt
}

// SetValueCreatedAt время создания в формате RFC3339
func (e *CommonEventCaseObject) SetValueCreatedAt(v string) {
	e.CreatedAt = v
}

// SetAnyCreatedAt время создания
func (e *CommonEventCaseObject) SetAnyCreatedAt(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetUpdatedAt время обновления в формате RFC3339
func (e *CommonEventCaseObject) GetUpdatedAt() string {
	return e.UpdatedAt
}

// SetValueUpdatedAt время обновления в формате RFC3339
func (e *CommonEventCaseObject) SetValueUpdatedAt(v string) {
	e.UpdatedAt = v
}

// SetAnyUpdatedAt время обновления
func (e *CommonEventCaseObject) SetAnyUpdatedAt(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetUnderliningId идентификатор
func (e *CommonEventCaseObject) GetUnderliningId() string {
	return e.UnderliningId
}

// SetValueUnderliningId идентификатор
func (e *CommonEventCaseObject) SetValueUnderliningId(v string) {
	e.UnderliningId = v
}

// SetAnyUnderliningId идентификатор
func (e *CommonEventCaseObject) SetAnyUnderliningId(i any) {
	e.UnderliningId = fmt.Sprint(i)
}

// GetId идентификатор
func (e *CommonEventCaseObject) GetId() string {
	return e.Id
}

// SetValueId идентификатор
func (e *CommonEventCaseObject) SetValueId(v string) {
	e.Id = v
}

// SetAnyId идентификатор
func (e *CommonEventCaseObject) SetAnyId(i any) {
	e.Id = fmt.Sprint(i)
}

// GetCreatedBy создатель
func (e *CommonEventCaseObject) GetCreatedBy() string {
	return e.CreatedBy
}

// SetValueCreatedBy создатель
func (e *CommonEventCaseObject) SetValueCreatedBy(v string) {
	e.CreatedBy = v
}

// SetAnyCreatedBy создатель
func (e *CommonEventCaseObject) SetAnyCreatedBy(i any) {
	e.CreatedBy = fmt.Sprint(i)
}

// GetUpdatedBy лицо выполнившее обновление
func (e *CommonEventCaseObject) GetUpdatedBy() string {
	return e.UpdatedBy
}

// SetValueUpdatedBy лицо выполнившее обновление
func (e *CommonEventCaseObject) SetValueUpdatedBy(v string) {
	e.UpdatedBy = v
}

// SetAnyUpdatedBy лицо выполнившее обновление
func (e *CommonEventCaseObject) SetAnyUpdatedBy(i any) {
	e.UpdatedBy = fmt.Sprint(i)
}

// GetUnderliningType значение поля UnderliningType
func (e *CommonEventCaseObject) GetUnderliningType() string {
	return e.UnderliningType
}

// SetValueUnderliningType значение поля UnderliningType
func (e *CommonEventCaseObject) SetValueUnderliningType(v string) {
	e.UnderliningType = v
}

// SetAnyUnderliningType значение поля UnderliningType
func (e *CommonEventCaseObject) SetAnyUnderliningType(i any) {
	e.UnderliningType = fmt.Sprint(i)
}

// GetTitle заголовок
func (e *CommonEventCaseObject) GetTitle() string {
	return e.Title
}

// SetValueTitle заголовок
func (e *CommonEventCaseObject) SetValueTitle(v string) {
	e.Title = v
}

// SetAnyTitle заголовок
func (e *CommonEventCaseObject) SetAnyTitle(i any) {
	e.Title = fmt.Sprint(i)
}

// GetDescription описание
func (e *CommonEventCaseObject) GetDescription() string {
	return e.Description
}

// SetValueDescription описание
func (e *CommonEventCaseObject) SetValueDescription(v string) {
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	e.Description = v
}

// SetAnyDescription описание
func (e *CommonEventCaseObject) SetAnyDescription(i interface{}) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	e.Description = str
}

// GetImpactStatus значение поля ImpactStatus
func (e *CommonEventCaseObject) GetImpactStatus() string {
	return e.ImpactStatus
}

// SetValueImpactStatus значение поля ImpactStatus
func (e *CommonEventCaseObject) SetValueImpactStatus(v string) {
	e.ImpactStatus = v
}

// SetAnyImpactStatus значение поля ImpactStatus
func (e *CommonEventCaseObject) SetAnyImpactStatus(i any) {
	e.ImpactStatus = fmt.Sprint(i)
}

// GetResolutionStatus значение поля ResolutionStatus
func (e *CommonEventCaseObject) GetResolutionStatus() string {
	return e.ResolutionStatus
}

// SetValueResolutionStatus значение поля ResolutionStatus
func (e *CommonEventCaseObject) SetValueResolutionStatus(v string) {
	e.ResolutionStatus = v
}

// SetAnyResolutionStatus значение поля ResolutionStatus
func (e *CommonEventCaseObject) SetAnyResolutionStatus(i any) {
	e.ResolutionStatus = fmt.Sprint(i)
}

// GetStatus значение поля Status
func (e *CommonEventCaseObject) GetStatus() string {
	return e.Status
}

// SetValueStatus значение поля Status
func (e *CommonEventCaseObject) SetValueStatus(v string) {
	e.Status = v
}

// SetAnyStatus значение поля Status
func (e *CommonEventCaseObject) SetAnyStatus(i any) {
	e.Status = fmt.Sprint(i)
}

// GetSummary значение поля Summary
func (e *CommonEventCaseObject) GetSummary() string {
	return e.Summary
}

// SetValueSummary значение поля Summary
func (e *CommonEventCaseObject) SetValueSummary(v string) {
	e.Summary = v
}

// SetAnySummary значение поля Summary
func (e *CommonEventCaseObject) SetAnySummary(i any) {
	e.Summary = fmt.Sprint(i)
}

// GetOwner значение поля Owner
func (e *CommonEventCaseObject) GetOwner() string {
	return e.Owner
}

// SetValueOwner значение поля Owner
func (e *CommonEventCaseObject) SetValueOwner(v string) {
	e.Owner = v
}

// SetAnyOwner значение поля Owner
func (e *CommonEventCaseObject) SetAnyOwner(i any) {
	e.Owner = fmt.Sprint(i)
}

// ToStringBeautiful форматированный вывод
func (eo CommonEventCaseObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, eo.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, eo.Id))
	str.WriteString(fmt.Sprintf("%s'createdBy': '%s'\n", ws, eo.CreatedBy))
	str.WriteString(fmt.Sprintf("%s'updatedBy': '%s'\n", ws, eo.UpdatedBy))
	str.WriteString(fmt.Sprintf("%s'createdAt': '%s'\n", ws, eo.CreatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedAt': '%s'\n", ws, eo.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, eo.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'caseId': '%d'\n", ws, eo.CaseId))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, eo.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, eo.Description))
	str.WriteString(fmt.Sprintf("%s'severity': '%d'\n", ws, eo.Severity))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, eo.StartDate))
	str.WriteString(fmt.Sprintf("%s'endDate': '%s'\n", ws, eo.EndDate))
	str.WriteString(fmt.Sprintf("%s'impactStatus': '%s'\n", ws, eo.ImpactStatus))
	str.WriteString(fmt.Sprintf("%s'resolutionStatus': '%s'\n", ws, eo.ResolutionStatus))
	str.WriteString(fmt.Sprintf("%s'flag': '%v'\n", ws, eo.Flag))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, eo.Tlp))
	str.WriteString(fmt.Sprintf("%s'pap': '%d'\n", ws, eo.Pap))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, eo.Status))
	str.WriteString(fmt.Sprintf("%s'summary': '%s'\n", ws, eo.Summary))
	str.WriteString(fmt.Sprintf("%s'owner': '%s'\n", ws, eo.Owner))

	return str.String()
}
