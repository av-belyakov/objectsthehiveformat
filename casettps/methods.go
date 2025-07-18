package casettps

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

//************* ExtraDataTtp ***************

func (tm *Ttp) GetExtraData() ExtraDataTtp {
	return tm.ExtraData
}

// GetPattern шаблон
func (ed *ExtraDataTtp) GetPattern() PatternExtraData {
	return ed.Pattern
}

// SetValueDetails шаблон
func (ed *ExtraDataTtp) SetValuePattern(v PatternExtraData) {
	ed.Pattern = v
}

// GetPatternParent родительский шаблон
func (ed *ExtraDataTtp) GetPatternParent() PatternExtraData {
	return ed.PatternParent
}

// SetValuePatternParent родительский шаблон
func (ed *ExtraDataTtp) SetValuePatternParent(v PatternExtraData) {
	ed.PatternParent = v
}

// ToStringBeautiful форматированный вывод
func (edtm ExtraDataTtp) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'pattern':\n", ws))
	str.WriteString(edtm.Pattern.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'patternParent':\n", ws))
	str.WriteString(edtm.PatternParent.ToStringBeautiful(num + 1))

	return str.String()
}

//*********** PatternExtraData **************

func NewPatternExtraData() *PatternExtraData {
	return &PatternExtraData{
		UnderliningCreatedAt: "1970-01-01T00:00:00+00:00",
		Platforms:            []string(nil),
		PermissionsRequired:  []string(nil),
		DataSources:          []string(nil),
		Tactics:              []string(nil),
	}
}

func (ped *PatternExtraData) GetRemoteSupport() bool {
	return ped.RemoteSupport
}

// SetValueRemoteSupport устанавливает BOOL значение для поля RemoteSupport
func (ped *PatternExtraData) SetValueRemoteSupport(v bool) {
	ped.RemoteSupport = v
}

// SetAnyRemoteSupport устанавливает ЛЮБОЕ значение для поля RemoteSupport
func (ped *PatternExtraData) SetAnyRemoteSupport(i any) {
	if v, ok := i.(bool); ok {
		ped.RemoteSupport = v
	}
}

func (ped *PatternExtraData) GetRevoked() bool {
	return ped.Revoked
}

// SetValueRevoked устанавливает BOOL значение для поля Revoked
func (ped *PatternExtraData) SetValueRevoked(v bool) {
	ped.Revoked = v
}

// SetAnyRemoteSupport устанавливает ЛЮБОЕ значение для поля Revoked
func (ped *PatternExtraData) SetAnyRevoked(i any) {
	if v, ok := i.(bool); ok {
		ped.Revoked = v
	}
}

func (ped *PatternExtraData) GetUnderliningCreatedAt() string {
	return ped.UnderliningCreatedAt
}

// SetValueUnderliningCreatedAt устанавливает дату в формате RFC3339
func (ped *PatternExtraData) SetValueUnderliningCreatedAt(v string) {
	ped.UnderliningCreatedAt = v
}

// SetAnyUnderliningCreatedAt устанавливает ЛЮБОЕ значение для поля UnderliningCreatedAt
func (ped *PatternExtraData) SetAnyUnderliningCreatedAt(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	ped.UnderliningCreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (ped *PatternExtraData) GetUnderliningCreatedBy() string {
	return ped.UnderliningCreatedBy
}

// SetValueUnderliningCreatedBy устанавливает значение поля UnderliningCreatedBy
func (ped *PatternExtraData) SetValueUnderliningCreatedBy(v string) {
	ped.UnderliningCreatedBy = v
}

// SetAnyUnderliningCreatedBy устанавливает ЛЮБОЕ значение для поля UnderliningCreatedBy
func (ped *PatternExtraData) SetAnyUnderliningCreatedBy(i any) {
	ped.UnderliningCreatedBy = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetUnderliningId() string {
	return ped.UnderliningId
}

// SetValueUnderliningId устанавливает значение поля UnderliningId
func (ped *PatternExtraData) SetValueUnderliningId(v string) {
	ped.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (ped *PatternExtraData) SetAnyUnderliningId(i any) {
	ped.UnderliningId = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetUnderliningType() string {
	return ped.UnderliningType
}

// SetValueUnderliningType устанавливает значение поля UnderliningType
func (ped *PatternExtraData) SetValueUnderliningType(v string) {
	ped.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (ped *PatternExtraData) SetAnyUnderliningType(i any) {
	ped.UnderliningType = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetDetection() string {
	return ped.Detection
}

// SetAnyDetection устанавливает значение поля Detection
func (ped *PatternExtraData) SetValueDetection(v string) {
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	ped.Detection = v
}

// SetAnyDetection устанавливает ЛЮБОЕ значение для поля Detection
func (ped *PatternExtraData) SetAnyDetection(i any) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	ped.Detection = str
}

func (ped *PatternExtraData) GetDescription() string {
	return ped.Description
}

// SetValueDescription устанавливает значение поля Description
func (ped *PatternExtraData) SetValueDescription(v string) {
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	ped.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (ped *PatternExtraData) SetAnyDescription(i any) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	ped.Description = str
}

func (ped *PatternExtraData) GetName() string {
	return ped.Name
}

// SetValueName устанавливает значение поля Name
func (ped *PatternExtraData) SetValueName(v string) {
	ped.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (ped *PatternExtraData) SetAnyName(i any) {
	ped.Name = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetPatternId() string {
	return ped.PatternId
}

// SetValuePatternId устанавливает значение поля PatternId
func (ped *PatternExtraData) SetValuePatternId(v string) {
	ped.PatternId = v
}

// SetAnyPatternId устанавливает ЛЮБОЕ значение для поля PatternId
func (ped *PatternExtraData) SetAnyPatternId(i any) {
	ped.PatternId = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetPatternType() string {
	return ped.PatternType
}

// SetValuePatternType устанавливает значение поля PatternType
func (ped *PatternExtraData) SetValuePatternType(v string) {
	ped.PatternType = v
}

// SetAnyPatternType устанавливает ЛЮБОЕ значение для поля PatternType
func (ped *PatternExtraData) SetAnyPatternType(i any) {
	ped.PatternType = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetURL() string {
	return ped.URL
}

// SetValueURL устанавливает значение поля URL
func (ped *PatternExtraData) SetValueURL(v string) {
	ped.URL = v
}

// SetAnyURL устанавливает ЛЮБОЕ значение для поля URL
func (ped *PatternExtraData) SetAnyURL(i any) {
	ped.URL = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetVersion() string {
	return ped.Version
}

// SetValueVersion устанавливает значение поля Version
func (ped *PatternExtraData) SetValueVersion(v string) {
	ped.Version = v
}

// SetAnyVersion устанавливает ЛЮБОЕ значение для поля Version
func (ped *PatternExtraData) SetAnyVersion(i any) {
	ped.Version = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetPlatforms() []string {
	return ped.Platforms
}

// SetValuePlatforms устанавливает STRING значение для поля Platforms
func (ped *PatternExtraData) SetValuePlatforms(v string) {
	if ped.Platforms == nil {
		ped.Platforms = []string(nil)
	}

	ped.Platforms = append(ped.Platforms, v)
}

// SetAnyPlatforms устанавливает ЛЮБОЕ значение для поля Platforms
func (ped *PatternExtraData) SetAnyPlatforms(i any) {
	ped.SetValuePlatforms(fmt.Sprint(i))
}

func (ped *PatternExtraData) GetPermissionsRequired() []string {
	return ped.PermissionsRequired
}

// SetValuePermissionsRequired устанавливает STRING значение для поля PermissionsRequired
func (ped *PatternExtraData) SetValuePermissionsRequired(v string) {
	if ped.PermissionsRequired == nil {
		ped.PermissionsRequired = []string(nil)
	}

	ped.PermissionsRequired = append(ped.PermissionsRequired, v)
}

// SetAnyPermissionsRequired устанавливает ЛЮБОЕ значение для поля PermissionsRequired
func (ped *PatternExtraData) SetAnyPermissionsRequired(i any) {
	ped.SetValuePermissionsRequired(fmt.Sprint(i))
}

func (ped *PatternExtraData) GetDataSources() []string {
	return ped.DataSources
}

// SetValueDataSources устанавливает STRING значение для поля DataSources
func (ped *PatternExtraData) SetValueDataSources(v string) {
	if ped.DataSources == nil {
		ped.DataSources = []string(nil)
	}

	ped.DataSources = append(ped.DataSources, v)
}

// SetAnyDataSources устанавливает ЛЮБОЕ значение для поля DataSources
func (ped *PatternExtraData) SetAnyDataSources(i any) {
	ped.SetValueDataSources(fmt.Sprint(i))
}

func (ped *PatternExtraData) GetTactics() []string {
	return ped.Tactics
}

// SetValueTactics устанавливает STRING значение для поля Tactics
func (ped *PatternExtraData) SetValueTactics(v string) {
	if ped.Tactics == nil {
		ped.Tactics = []string(nil)
	}

	ped.Tactics = append(ped.Tactics, v)
}

// SetAnyTactics устанавливает ЛЮБОЕ значение для поля Tactics
func (ped *PatternExtraData) SetAnyTactics(i any) {
	ped.SetValueTactics(fmt.Sprint(i))
}

func (ped PatternExtraData) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_createdAt': '%s'\n", ws, ped.UnderliningCreatedAt))
	str.WriteString(fmt.Sprintf("%s'_createdBy': '%s'\n", ws, ped.UnderliningCreatedBy))
	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, ped.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, ped.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'dataSources': \n%v", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}

		return str.String()
	}(ped.DataSources)))
	/*str.WriteString(fmt.Sprintf("%sdefenseBypassed: \n%v", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}
		return str.String()
	}(ped.DefenseBypassed)))*/
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, ped.Description))
	/*str.WriteString(fmt.Sprintf("%sextraData: \n%s", ws, func(l map[string]interface{}) string {
		var str strings.Builder = string.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%s: '%v'\n", supportingfunctions.GetWhitespace(num+1), k, v))
		}
		return str
	}(ped.ExtraData)))*/
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, ped.Name))
	str.WriteString(fmt.Sprintf("%s'patternId': '%s'\n", ws, ped.PatternId))
	str.WriteString(fmt.Sprintf("%s'patternType': '%s'\n", ws, ped.PatternType))
	str.WriteString(fmt.Sprintf("%s'permissionsRequired': \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}

		return str.String()
	}(ped.PermissionsRequired)))
	str.WriteString(fmt.Sprintf("%s'platforms': \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}

		return str.String()
	}(ped.Platforms)))
	str.WriteString(fmt.Sprintf("%s'remoteSupport': '%v'\n", ws, ped.RemoteSupport))
	str.WriteString(fmt.Sprintf("%s'revoked': '%v'\n", ws, ped.Revoked))
	/*str.WriteString(fmt.Sprintf("%ssystemRequirements: \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder()
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}
		return str.String()
	}(ped.SystemRequirements)))*/
	str.WriteString(fmt.Sprintf("%s'tactics': \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}

		return str.String()
	}(ped.Tactics)))
	str.WriteString(fmt.Sprintf("%s'url': '%s'\n", ws, ped.URL))
	str.WriteString(fmt.Sprintf("%s'version': '%s'\n", ws, ped.Version))

	return str.String()
}
