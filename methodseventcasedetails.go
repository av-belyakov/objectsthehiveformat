package objectsthehiveformat

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

func NewEventCaseDetails() *EventCaseDetails {
	return &EventCaseDetails{
		EndDate:      "1970-01-01T00:00:00+00:00",
		CustomFields: common.CustomFields{},
	}
}

func (e *EventCaseDetails) Get() *EventCaseDetails {
	return e
}

// GetEndDate дата завершения в формате RFC3339
func (e *EventCaseDetails) GetEndDate() string {
	return e.EndDate
}

// SetValueEndDate дата завершения в формате RFC3339
func (e *EventCaseDetails) SetValueEndDate(v string) {
	e.EndDate = v
}

// SetAnyEndDate дата завершения в формате RFC3339
func (e *EventCaseDetails) SetAnyEndDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.EndDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetResolutionStatus статус постановления
func (e *EventCaseDetails) GetResolutionStatus() string {
	return e.ResolutionStatus
}

// SetValueResolutionStatus статус постановления
func (e *EventCaseDetails) SetValueResolutionStatus(v string) {
	e.ResolutionStatus = v
}

// SetAnyResolutionStatus статус постановления
func (e *EventCaseDetails) SetAnyResolutionStatus(i interface{}) {
	e.ResolutionStatus = fmt.Sprint(i)
}

// GetSummary резюме
func (e *EventCaseDetails) GetSummary() string {
	return e.Summary
}

// SetValueSummary резюме
func (e *EventCaseDetails) SetValueSummary(v string) {
	e.Summary = v
}

// SetAnySummary резюме
func (e *EventCaseDetails) SetAnySummary(i interface{}) {
	e.Summary = fmt.Sprint(i)
}

// GetStatus статус
func (e *EventCaseDetails) GetStatus() string {
	return e.Status
}

// SetValueStatus статус
func (e *EventCaseDetails) SetValueStatus(v string) {
	e.Status = v
}

// SetAnyStatus статус
func (e *EventCaseDetails) SetAnyStatus(i interface{}) {
	e.Status = fmt.Sprint(i)
}

// GetImpactStatus краткое описание воздействия
func (e *EventCaseDetails) GetImpactStatus() string {
	return e.ImpactStatus
}

// SetValueImpactStatus краткое описание воздействия
func (e *EventCaseDetails) SetValueImpactStatus(v string) {
	e.ImpactStatus = v
}

// SetAnyImpactStatus краткое описание воздействия
func (e *EventCaseDetails) SetAnyImpactStatus(i interface{}) {
	e.ImpactStatus = fmt.Sprint(i)
}

// GetCustomFields настраиваемые поля
func (e *EventCaseDetails) GetCustomFields() common.CustomFields {
	return e.CustomFields
}

// SetValueCustomFields настраиваемые поля
func (e *EventCaseDetails) SetValueCustomFields(v common.CustomFields) {
	e.CustomFields = v
}

// ToStringBeautiful форматированный вывод
func (ed EventCaseDetails) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'endDate': '%s'\n", ws, ed.EndDate))
	str.WriteString(fmt.Sprintf("%s'resolutionStatus': '%s'\n", ws, ed.ResolutionStatus))
	str.WriteString(fmt.Sprintf("%s'summary': '%s'\n", ws, ed.Summary))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, ed.Status))
	str.WriteString(fmt.Sprintf("%s'impactStatus': '%s'\n", ws, ed.ImpactStatus))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, common.CustomFieldsToStringBeautiful(ed.CustomFields, num)))

	return str.String()
}
