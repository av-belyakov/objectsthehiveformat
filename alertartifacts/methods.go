package alertartifacts

import (
	"fmt"
	"strconv"
	"strings"

	"slices"

	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewArtifact создаёт новый объект типа Artifacts
func NewArtifact() *Artifacts {
	return &Artifacts{
		CommonArtifactType: common.CommonArtifactType{
			CreatedAt: "1970-01-01T00:00:00+00:00",
			StartDate: "1970-01-01T00:00:00+00:00",
		},
		SnortSid:       []string(nil),
		SnortSidNumber: []int(nil),
		Tags:           make(map[string][]string),
		TagsAll:        []string(nil),
	}
}

// Get возвращает объект типа Artifacts
func (a *Artifacts) Get() *Artifacts {
	return a
}

func (a *Artifacts) GetSensorId() string {
	return a.SensorId
}

// SetSensorId добавляет значение SensorId
func (a *Artifacts) SetValueSensorId(v string) {
	a.SensorId = v
}

// SetSensorId добавляет значение SensorId
func (a *Artifacts) SetAnySensorId(i any) {
	a.SensorId = fmt.Sprint(i)
}

func (a *Artifacts) GetSnortSid() []string {
	return a.SnortSid
}

// SetValueSnortSid добавляет значение в список поля SnortSid
func (a *Artifacts) SetValueSnortSid(v string) {
	if a.SnortSid == nil {
		a.SnortSid = []string(nil)
	}

	a.SnortSid = append(a.SnortSid, v)
}

// SetAnySnortSid добавляет значение в список поля SnortSid
func (a *Artifacts) SetAnySnortSid(i any) {
	a.SetValueSnortSid(fmt.Sprint(i))
}

func (a *Artifacts) GetSnortSidNumber() []int {
	return a.SnortSidNumber
}

// SetValueSnortSidNumber добавляет значение в список поля SnortSidNumber
func (a *Artifacts) SetValueSnortSidNumber(v int) {
	if a.SnortSidNumber == nil {
		a.SnortSidNumber = []int(nil)
	}

	a.SnortSidNumber = append(a.SnortSidNumber, v)
}

// SetAnySnortSidNumber добавляет значение в список поля SnortSidNumber
func (a *Artifacts) SetAnySnortSidNumber(i any) {
	str := fmt.Sprint(i)
	if num, err := strconv.ParseUint(str, 10, 32); err == nil {
		a.SetValueSnortSidNumber(int(num))
	}
}

func (a *Artifacts) GetTags() map[string][]string {
	return a.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (a *Artifacts) SetValueTags(k, v string) bool {
	if a.Tags == nil {
		a.Tags = make(map[string][]string)
	}

	if _, ok := a.Tags[k]; !ok {
		a.Tags[k] = []string(nil)
	}

	if slices.Contains(a.Tags[k], v) {
		return false
	}

	a.Tags[k] = append(a.Tags[k], v)

	return true
}

// SetAnyTags устанавливает значение для поля Tags
func (a *Artifacts) SetAnyTags(k string, i any) bool {
	return a.SetValueTags(k, fmt.Sprint(i))
}

func (a *Artifacts) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll добавляет значение в список поля TagsAll
func (a *Artifacts) SetValueTagsAll(v string) {
	if a.TagsAll == nil {
		a.TagsAll = []string(nil)
	}

	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll добавляет значение в список поля TagsAll
func (a *Artifacts) SetAnyTagsAll(i any) {
	a.SetValueTagsAll(fmt.Sprint(i))
}

func (a *Artifacts) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(a.CommonArtifactType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'sensorId': '%s'\n", ws, a.SensorId))
	str.WriteString(fmt.Sprintf("%s'snortSid': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, a.SnortSid)))
	str.WriteString(fmt.Sprintf("%s'snortSidNumber': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, a.SnortSidNumber)))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulMapSlice(num, a.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, a.TagsAll)))

	return str.String()
}
