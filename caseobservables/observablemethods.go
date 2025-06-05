package caseobservables

import (
	"fmt"
	"strconv"
	"strings"

	"slices"

	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewObservable создаёт новый объект типа Observable
func NewObservable() *Observable {
	return &Observable{
		CommonObservableType: *common.NewCommonObservableType(),
		Tags:                 make(map[string][]string),
		TagsAll:              []string(nil),
		Attachment:           *common.NewAttachmentData(),
		//Reports:    make(map[string]ReportTaxonomies),
	}
}

func (o *Observable) Get() *Observable {
	return o
}

// GetSensorId идентификатор сенсора
func (o *Observable) GetSensorId() string {
	return o.SensorId
}

// SetSensorId идентификатор сенсора
func (o *Observable) SetValueSensorId(v string) {
	o.SensorId = v
}

// SetSensorId идентификатор сенсора
func (o *Observable) SetAnySensorId(i any) {
	o.SensorId = fmt.Sprint(i)
}

// GetSnortSid идентификатор snort
func (o *Observable) GetSnortSid() []string {
	return o.SnortSid
}

// SetValueSnortSid идентификатор snort
func (o *Observable) SetValueSnortSid(v string) {
	if o.SnortSid == nil {
		o.SnortSid = []string(nil)
	}

	o.SnortSid = append(o.SnortSid, v)
}

// SetValueSnortSid идентификатор snort
func (o *Observable) SetAnySnortSid(i any) {
	o.SetValueSnortSid(fmt.Sprint(i))
}

// GetSnortSidNumber список идентификаторов snort
func (o *Observable) GetSnortSidNumber() []int {
	return o.SnortSidNumber
}

// SetValueSnortSidNumber список идентификаторов snort
func (o *Observable) SetValueSnortSidNumber(v int) {
	if o.SnortSidNumber == nil {
		o.SnortSidNumber = []int(nil)
	}

	o.SnortSidNumber = append(o.SnortSidNumber, v)
}

// SetAnySnortSidNumber список идентификаторов snort
func (o *Observable) SetAnySnortSidNumber(i any) {
	str := fmt.Sprint(i)
	if num, err := strconv.ParseUint(str, 10, 32); err == nil {
		o.SetValueSnortSidNumber(int(num))
	}
}

// GetTags список тегов
func (o *Observable) GetTags() map[string][]string {
	return o.Tags
}

// SetValueTags список тегов
func (o *Observable) SetValueTags(k, v string) bool {
	if o.Tags == nil {
		o.Tags = make(map[string][]string)
	}

	if _, ok := o.Tags[k]; !ok {
		o.Tags[k] = []string(nil)
	}

	if slices.Contains(o.Tags[k], v) {
		return false
	}

	o.Tags[k] = append(o.Tags[k], v)

	return true
}

// SetAnyTags список тегов
func (o *Observable) SetAnyTags(k string, i any) bool {
	return o.SetValueTags(k, fmt.Sprint(i))
}

// GetTagsAll список тегов
func (o *Observable) GetTagsAll() []string {
	return o.TagsAll
}

// SetValueTags список тегов
func (o *Observable) SetValueTagsAll(v string) {
	if o.TagsAll == nil {
		o.TagsAll = []string(nil)
	}

	o.TagsAll = append(o.TagsAll, v)
}

// SetAnyTagsAll список тегов
func (o *Observable) SetAnyTagsAll(i any) {
	o.SetValueTagsAll(fmt.Sprint(i))
}

func (o *Observable) GetAttachment() *common.AttachmentData {
	return &o.Attachment
}

/*
ИСКЛЮЧИЛ из-за черезмерно большого количества
 полей которое влечет за собой превышения лимита маппинга в Elsticsearch
 что выражается в ошибке от СУБД типа "Limit of total fields [2000] has been exceeded while adding new fields"
func (o *Observable) GetReports() map[string]ReportTaxonomies {
	return o.Reports
}

// GetTaxonomies возвращает список Taxonomy по заданному ключу
func (o *Observable) GetTaxonomies(key string) (*ReportTaxonomies, bool) {
	if res, ok := o.Reports[key]; ok {
		return &res, true
	}

	return &ReportTaxonomies{}, false
}

// SetValueReports устанавливает значение для поля Reports
func (o *Observable) SetValueReports(v map[string]ReportTaxonomies) {
	o.Reports = v
}

// AddValueReports добавляет список значений ReportTaxonomies
func (o *Observable) AddValueReports(key string, rt ReportTaxonomies) {
	o.Reports[key] = rt
}*/

// ToStringBeautiful форматированный вывод
func (om Observable) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(om.CommonObservableType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'sensorId': '%s'\n", ws, om.SensorId))
	str.WriteString(fmt.Sprintf("%s'snortSid': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, om.SnortSid)))
	str.WriteString(fmt.Sprintf("%s'snortSidNumber': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, om.SnortSidNumber)))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulMapSlice(num, om.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, om.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'attachment': \n%s", ws, om.Attachment.ToStringBeautiful(num+1)))
	/*str.WriteString(fmt.Sprintf("%s'reports': \n%s", ws, func(l map[string]ReportTaxonomies) string {
		var str strings.Builder = strings.Builder{}
		for key, value := range l {
			str.WriteString(fmt.Sprintf("%s'%s':\n", supportingfunctions.GetWhitespace(num+1), key))
			str.WriteString(fmt.Sprintf("%s'taxonomys':\n", supportingfunctions.GetWhitespace(num+2)))
			for k, v := range value.Taxonomies {
				str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+3), k+1))
				str.WriteString(fmt.Sprintf("%s'level': %v\n", supportingfunctions.GetWhitespace(num+4), v.Level))
				str.WriteString(fmt.Sprintf("%s'namespace': %v\n", supportingfunctions.GetWhitespace(num+4), v.Namespace))
				str.WriteString(fmt.Sprintf("%s'predicate': %v\n", supportingfunctions.GetWhitespace(num+4), v.Predicate))
				str.WriteString(fmt.Sprintf("%s'value': %v\n", supportingfunctions.GetWhitespace(num+4), v.Value))
			}
		}
		return str.String()
	}(om.Reports)))*/

	return str.String()
}
