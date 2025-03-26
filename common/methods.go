package common

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

//****************** CommonEventCase ******************

func NewCommonEventType() *CommonEventType {
	return &CommonEventType{
		StartDate: "1970-01-01T00:00:00+00:00",
	}
}

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

func CustomFieldsToStringBeautiful(l CustomFields, num int) string {
	strB := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num + 2)

	for k, v := range l {
		strB.WriteString(fmt.Sprintf("%s'%s':\n", supportingfunctions.GetWhitespace(num+1), k))

		nameOne, dataOne, nameTwo, dataTwo := v.Get()
		strB.WriteString(fmt.Sprintf("%s'%s': %d\n", ws, nameOne, dataOne))
		strB.WriteString(fmt.Sprintf("%s'%s': %s\n", ws, nameTwo, dataTwo))
	}

	return strB.String()
}

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

//************* CommonObservableType **************

func NewCommonObservableType() *CommonObservableType {
	return &CommonObservableType{
		UnderliningCreatedAt: "1970-01-01T00:00:00+00:00",
		UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
		StartDate:            "1970-01-01T00:00:00+00:00",
	}
}

func (o *CommonObservableType) Get() *CommonObservableType {
	return o
}

func (o *CommonObservableType) GetIoc() bool {
	return o.Ioc
}

// SetValueIoc устанавливает BOOL значение для поля Ioc
func (o *CommonObservableType) SetValueIoc(v bool) {
	o.Ioc = v
}

// SetAnyIoc устанавливает ЛЮБОЕ значение для поля Ioc
func (o *CommonObservableType) SetAnyIoc(i interface{}) {
	if v, ok := i.(bool); ok {
		o.Ioc = v
	}
}

func (o *CommonObservableType) GetSighted() bool {
	return o.Sighted
}

// SetValueSighted устанавливает BOOL значение для поля Sighted
func (o *CommonObservableType) SetValueSighted(v bool) {
	o.Sighted = v
}

// SetAnySighted устанавливает ЛЮБОЕ значение для поля Sighted
func (o *CommonObservableType) SetAnySighted(i interface{}) {
	if v, ok := i.(bool); ok {
		o.Sighted = v
	}
}

func (o *CommonObservableType) GetIgnoreSimilarity() bool {
	return o.IgnoreSimilarity
}

// SetValueIgnoreSimilarity устанавливает BOOL значение для поля IgnoreSimilarity
func (o *CommonObservableType) SetValueIgnoreSimilarity(v bool) {
	o.IgnoreSimilarity = v
}

// SetAnyIgnoreSimilarity устанавливает ЛЮБОЕ значение для поля IgnoreSimilarity
func (o *CommonObservableType) SetAnyIgnoreSimilarity(i interface{}) {
	if v, ok := i.(bool); ok {
		o.IgnoreSimilarity = v
	}
}

func (o *CommonObservableType) GetTlp() uint64 {
	return o.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (o *CommonObservableType) SetValueTlp(v uint64) {
	o.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (o *CommonObservableType) SetAnyTlp(i interface{}) {
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

func (o *CommonObservableType) GetUnderliningCreatedAt() string {
	return o.UnderliningCreatedAt
}

// SetValueUnderliningCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (o *CommonObservableType) SetValueUnderliningCreatedAt(v string) {
	o.UnderliningCreatedAt = v
}

// SetAnyUnderliningCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (o *CommonObservableType) SetAnyUnderliningCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UnderliningCreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetUnderliningUpdatedAt() string {
	return o.UnderliningUpdatedAt
}

// SetValueUnderliningUpdatedAt устанавливает значениев формате RFC3339 для поля UpdatedAt
func (o *CommonObservableType) SetValueUnderliningUpdatedAt(v string) {
	o.UnderliningUpdatedAt = v
}

// SetAnyUnderliningUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (o *CommonObservableType) SetAnyUnderliningUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UnderliningUpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetStartDate() string {
	return o.StartDate
}

// SetValueStartDate устанавливает значениев формате RFC3339 для поля StartDate
func (o *CommonObservableType) SetValueStartDate(v string) {
	o.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (o *CommonObservableType) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetUnderliningCreatedBy() string {
	return o.UnderliningCreatedBy
}

// SetValueUnderliningCreatedBy устанавливает STRING значение для поля CreatedBy
func (o *CommonObservableType) SetValueUnderliningCreatedBy(v string) {
	o.UnderliningCreatedBy = v
}

// SetAnyUnderliningCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (o *CommonObservableType) SetAnyUnderliningCreatedBy(i interface{}) {
	o.UnderliningCreatedBy = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningUpdatedBy() string {
	return o.UnderliningUpdatedBy
}

// SetValueUnderliningUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (o *CommonObservableType) SetValueUnderliningUpdatedBy(v string) {
	o.UnderliningUpdatedBy = v
}

// SetAnyUnderliningUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (o *CommonObservableType) SetAnyUnderliningUpdatedBy(i interface{}) {
	o.UnderliningUpdatedBy = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningId() string {
	return o.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (o *CommonObservableType) SetValueUnderliningId(v string) {
	o.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (o *CommonObservableType) SetAnyUnderliningId(i interface{}) {
	o.UnderliningId = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningType() string {
	return o.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (o *CommonObservableType) SetValueUnderliningType(v string) {
	o.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (o *CommonObservableType) SetAnyUnderliningType(i interface{}) {
	o.UnderliningType = fmt.Sprint(i)
}

func (o *CommonObservableType) GetData() string {
	return o.Data
}

// SetValueData устанавливает STRING значение для поля Data
func (o *CommonObservableType) SetValueData(v string) {
	o.Data = v
}

// SetAnyData устанавливает ЛЮБОЕ значение для поля Data
func (o *CommonObservableType) SetAnyData(i interface{}) {
	o.Data = fmt.Sprint(i)
}

func (o *CommonObservableType) GetDataType() string {
	return o.DataType
}

// SetValueDataType устанавливает STRING значение для поля DataType
func (o *CommonObservableType) SetValueDataType(v string) {
	o.DataType = v
}

// SetAnyDataType устанавливает ЛЮБОЕ значение для поля DataType
func (o *CommonObservableType) SetAnyDataType(i interface{}) {
	o.DataType = fmt.Sprint(i)
}

func (o *CommonObservableType) GetMessage() string {
	return o.Message
}

// SetValueMessage устанавливает STRING значение для поля Message
func (o *CommonObservableType) SetValueMessage(v string) {
	o.Message = v
}

// SetAnyMessage устанавливает ЛЮБОЕ значение для поля Message
func (o *CommonObservableType) SetAnyMessage(i interface{}) {
	o.Message = fmt.Sprint(i)
}

// ToStringBeautiful форматированный вывод
func (om CommonObservableType) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_createdAt': '%s'\n", ws, om.UnderliningCreatedAt))
	str.WriteString(fmt.Sprintf("%s'_createdBy': '%s'\n", ws, om.UnderliningCreatedBy))
	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, om.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, om.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'_updatedAt': '%s'\n", ws, om.UnderliningUpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_updatedBy': '%s'\n", ws, om.UnderliningUpdatedBy))
	str.WriteString(fmt.Sprintf("%s'data': '%s'\n", ws, om.Data))
	str.WriteString(fmt.Sprintf("%s'dataType': '%s'\n", ws, om.DataType))
	str.WriteString(fmt.Sprintf("%s'ignoreSimilarity': '%v'\n", ws, om.IgnoreSimilarity))
	str.WriteString(fmt.Sprintf("%s'ioc': '%v'\n", ws, om.Ioc))
	str.WriteString(fmt.Sprintf("%s'message': '%s'\n", ws, om.Message))
	str.WriteString(fmt.Sprintf("%s'sighted': '%v'\n", ws, om.Sighted))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, om.StartDate))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, om.Tlp))

	return str.String()
}

// ****************** AttachmentData ********************

func NewAttachmentData() *AttachmentData {
	return &AttachmentData{Hashes: []string(nil)}
}

// GetSize размер
func (a *AttachmentData) GetSize() uint64 {
	return a.Size
}

// SetValueSize размер
func (a *AttachmentData) SetValueSize(v uint64) {
	a.Size = v
}

// SetAnySize размер
func (a *AttachmentData) SetAnySize(i any) {
	if v, ok := i.(float32); ok {
		a.Size = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		a.Size = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Size = v
	}
}

// GetId идентификатор
func (a *AttachmentData) GetId() string {
	return a.Id
}

// SetValueId идентификатор
func (a *AttachmentData) SetValueId(v string) {
	a.Id = v
}

// SetAnyId идентификатор
func (a *AttachmentData) SetAnyId(i interface{}) {
	a.Id = fmt.Sprint(i)
}

// GetName имя
func (a *AttachmentData) GetName() string {
	return a.Name
}

// SetValueName имя
func (a *AttachmentData) SetValueName(v string) {
	a.Name = v
}

// SetAnyName имя
func (a *AttachmentData) SetAnyName(i interface{}) {
	a.Name = fmt.Sprint(i)
}

// GetContentType тип контента
func (a *AttachmentData) GetContentType() string {
	return a.ContentType
}

// SetValueContentType тип контента
func (a *AttachmentData) SetValueContentType(v string) {
	a.ContentType = v
}

// SetAnyContentType тип контента
func (a *AttachmentData) SetAnyContentType(i interface{}) {
	a.ContentType = fmt.Sprint(i)
}

// GetHashes хеши
func (a *AttachmentData) GetHashes() []string {
	return a.Hashes
}

// SetValueHashes хеши
func (a *AttachmentData) SetValueHashes(v string) {
	a.Hashes = append(a.Hashes, v)
}

// SetAnyHashes хеши
func (a *AttachmentData) SetAnyHashes(i interface{}) {
	a.Hashes = append(a.Hashes, fmt.Sprint(i))
}

// ToStringBeautiful форматированный вывод
func (a AttachmentData) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'size': '%d'\n", ws, a.Size))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, a.Id))
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, a.Name))
	str.WriteString(fmt.Sprintf("%s'contentType': '%s'\n", ws, a.ContentType))
	str.WriteString(fmt.Sprintf("%s'hashes': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, a.Hashes)))

	return str.String()
}
