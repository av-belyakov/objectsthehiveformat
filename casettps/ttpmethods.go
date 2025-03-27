package casettps

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewTtp создаёт объект типа Ttp
func NewTtp() *Ttp {
	return &Ttp{
		OccurDate:            "1970-01-01T00:00:00+00:00",
		UnderliningCreatedAt: "1970-01-01T00:00:00+00:00",
		ExtraData: ExtraDataTtp{
			Pattern:       *NewPatternExtraData(),
			PatternParent: *NewPatternExtraData(),
		},
	}
}

func (ttpm *Ttp) Get() *Ttp {
	return ttpm
}

// GetOccurDate дата происшествия в формате RFC3339
func (ttpm *Ttp) GetOccurDate() string {
	return ttpm.OccurDate
}

// SetValueOccurDate дата происшествия в формате RFC3339
func (ttpm *Ttp) SetValueOccurDate(v string) {
	ttpm.OccurDate = v
}

// SetAnyOccurDate дата происшествия
func (ttpm *Ttp) SetAnyOccurDate(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	ttpm.OccurDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetUnderliningCreatedAt дата создания в формате RFC3339
func (ttpm *Ttp) GetUnderliningCreatedAt() string {
	return ttpm.UnderliningCreatedAt
}

// SetValueUnderliningCreatedAt дата создания в формате RFC3339
func (ttpm *Ttp) SetValueUnderliningCreatedAt(v string) {
	ttpm.UnderliningCreatedAt = v
}

// SetAnyUnderliningCreatedAt дата создания
func (ttpm *Ttp) SetAnyUnderliningCreatedAt(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	ttpm.UnderliningCreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetUnderliningId идентификатор
func (ttpm *Ttp) GetUnderliningId() string {
	return ttpm.UnderliningId
}

// SetValueUnderliningId идентификатор
func (ttpm *Ttp) SetValueUnderliningId(v string) {
	ttpm.UnderliningId = v
}

// SetAnyUnderliningId идентификатор
func (ttpm *Ttp) SetAnyUnderliningId(i any) {
	ttpm.UnderliningId = fmt.Sprint(i)
}

// GetUnderliningCreatedBy создатель
func (ttpm *Ttp) GetUnderliningCreatedBy() string {
	return ttpm.UnderliningCreatedBy
}

// SetValueUnderliningCreatedBy создатель
func (ttpm *Ttp) SetValueUnderliningCreatedBy(v string) {
	ttpm.UnderliningCreatedBy = v
}

// SetAnyUnderliningCreatedBy создатель
func (ttpm *Ttp) SetAnyUnderliningCreatedBy(i any) {
	ttpm.UnderliningCreatedBy = fmt.Sprint(i)
}

// GetPatternId идентификатор шаблона
func (ttpm *Ttp) GetPatternId() string {
	return ttpm.PatternId
}

// SetValuePatternId идентификатор шаблона
func (ttpm *Ttp) SetValuePatternId(v string) {
	ttpm.PatternId = v
}

// SetAnyPatternId идентификатор шаблона
func (ttpm *Ttp) SetAnyPatternId(i any) {
	ttpm.PatternId = fmt.Sprint(i)
}

// GetTactic тактика
func (ttpm *Ttp) GetTactic() string {
	return ttpm.Tactic
}

// SetValueTactic тактика
func (ttpm *Ttp) SetValueTactic(v string) {
	ttpm.Tactic = v
}

// SetAnyTactic тактика
func (ttpm *Ttp) SetAnyTactic(i any) {
	ttpm.Tactic = fmt.Sprint(i)
}

// GetPattern шаблон
func (tm *Ttp) GetPattern() *PatternExtraData {
	return &tm.ExtraData.Pattern
}

// GetPatternParent родительский шаблон
func (tm *Ttp) GetPatternParent() *PatternExtraData {
	return &tm.ExtraData.PatternParent
}

// ToStringBeautiful форматированный вывод
func (tm Ttp) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_createdAt': '%s'\n", ws, tm.UnderliningCreatedAt))
	str.WriteString(fmt.Sprintf("%s'_createdBy': '%s'\n", ws, tm.UnderliningCreatedBy))
	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, tm.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'occurDate': '%s'\n", ws, tm.OccurDate))
	str.WriteString(fmt.Sprintf("%s'patternId': '%s'\n", ws, tm.PatternId))
	str.WriteString(fmt.Sprintf("%s'tactic': '%s'\n", ws, tm.Tactic))
	str.WriteString(fmt.Sprintf("%s'extraData':\n", ws))
	str.WriteString(tm.ExtraData.ToStringBeautiful(num + 1))

	return str.String()
}
