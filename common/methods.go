package common

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"maps"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

var list = map[string]string{
	"attack-type":         "string",
	"class-attack":        "string",
	"event-source":        "string",
	"external-letter":     "string",
	"geoip":               "string",
	"ncircc-class-attack": "string",
	"ncircc-bulletin-id":  "string",
	"inbox1":              "string",
	"inner-letter":        "string",
	"ir-name":             "string",
	"id-soa":              "string",
	"notification":        "string",
	"sphere":              "string",
	"state":               "string",
	"report":              "string",
	"first-time":          "date",
	"last-time":           "date",
	"b2mid":               "integer",
	"is-incident":         "boolen",
	"work-admin":          "boolen",
}

// ************* CustomFields **************

// Get поле CustomFields
func (fields *CustomFields) Get() CustomFields {
	return *fields
}

// Set поле CustomFields
func (fields *CustomFields) Set(v CustomFields) {
	if *fields == nil {
		*fields = make(CustomFields)
	}

	maps.Copy((*fields), v)
}

// UnmarshalJSON настраиваемый автоматический анмаршалинг
func (fields *CustomFields) UnmarshalJSON(data []byte) error {
	type tmpCustomFieldType map[string]*json.RawMessage

	tmp := tmpCustomFieldType{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	newResult := make(CustomFields, len(tmp))
	for k, v := range tmp {
		name, ok := list[k]
		if !ok {
			continue
		}

		var field CustomerFields
		switch name {
		case "string":
			field = &CustomFieldStringType{}

		case "date":
			field = &CustomFieldDateType{}

		case "integer":
			field = &CustomFieldIntegerType{}

		case "boolen":
			field = &CustomFieldBoolenType{}

		}

		if err := json.Unmarshal(*v, field); err != nil {
			return err
		}

		newResult[k] = field
	}

	fields.Set(newResult)

	return nil
}

// UnmarshalBSON настраиваемый автоматический анмаршалинг
func (fields *CustomFields) UnmarshalBSON(data []byte) error {
	type tmpCustomFieldType map[string]*bson.Raw

	tmp := tmpCustomFieldType{}
	if err := bson.Unmarshal(data, &tmp); err != nil {
		return err
	}

	//newResult := make(map[string]CustomerFields, len(tmp.CustomFields))
	newResult := make(CustomFields, len(tmp))
	for k, v := range tmp {
		name, ok := list[k]
		if !ok {
			continue
		}

		var field CustomerFields
		switch name {
		case "string":
			field = &CustomFieldStringType{}

		case "date":
			field = &CustomFieldDateType{}

		case "integer":
			field = &CustomFieldIntegerType{}

		case "boolen":
			field = &CustomFieldBoolenType{}

		}

		if err := json.Unmarshal(*v, field); err != nil {
			return err
		}

		newResult[k] = field
	}

	fields.Set(newResult)

	return nil
}

// Get возвращает значения CustomFieldStringType, где 1 и 3 значение это
// наименование поля
func (cf *CustomFieldStringType) Get() (string, int, string, string) {
	return "order", cf.Order, "string", cf.String
}

// Set устанавливает значения CustomFieldStringType
func (cf *CustomFieldStringType) Set(order, str interface{}) {
	cf.Order = supportingfunctions.ConversionAnyToInt(order)
	cf.String = fmt.Sprint(str)
}

// Get возвращает значения CustomFieldDateType, где 1 и 3 значение это
// наименование поля
func (cf *CustomFieldDateType) Get() (string, int, string, string) {
	return "order", cf.Order, "date", cf.Date
}

// Set устанавливает значения CustomFieldDateType
func (cf *CustomFieldDateType) Set(order, date interface{}) {
	cf.Order = supportingfunctions.ConversionAnyToInt(order)

	if str, ok := date.(string); ok {
		cf.Date = str

		return
	}

	tmp := supportingfunctions.ConversionAnyToInt(date)
	cf.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// Get возвращает значения CustomFieldIntegerType, где 1 и 3 значение это
// наименование поля
func (cf *CustomFieldIntegerType) Get() (string, int, string, string) {
	return "order", cf.Order, "integer", fmt.Sprint(cf.Integer)
}

// Set устанавливает значения CustomFieldIntegerType
func (cf *CustomFieldIntegerType) Set(order, integer interface{}) {
	cf.Order = supportingfunctions.ConversionAnyToInt(order)

	if i, ok := integer.(int); ok {
		cf.Integer = i
	}
}

// Get возвращает значения CustomFieldBoolenType, где 1 и 3 значение это
// наименование поля
func (cf *CustomFieldBoolenType) Get() (string, int, string, string) {
	return "order", cf.Order, "boolen", fmt.Sprint(cf.Boolean)
}

// Set устанавливает значения CustomFieldBoolenType
func (cf *CustomFieldBoolenType) Set(order, boolen interface{}) {
	cf.Order = supportingfunctions.ConversionAnyToInt(order)

	if i, ok := boolen.(bool); ok {
		cf.Boolean = i
	}
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
func (a *AttachmentData) SetAnyId(i any) {
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
func (a *AttachmentData) SetAnyName(i any) {
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
func (a *AttachmentData) SetAnyContentType(i any) {
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
func (a *AttachmentData) SetAnyHashes(i any) {
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

// ********************* ReportTaxonomys *******************
func (t *ReportTaxonomies) GetTaxonomys() []Taxonomy {
	return t.Taxonomies
}

func (t *ReportTaxonomies) GetReportTaxonomys() ReportTaxonomies {
	return *t
}

func (t *ReportTaxonomies) AddTaxonomy(taxonomy Taxonomy) {
	t.Taxonomies = append(t.Taxonomies, taxonomy)
}
