package common

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ************* CustomFields **************

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
