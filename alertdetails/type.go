package alertdetails

// EventAlertDetails детальная информация о событии
type EventAlertDetails struct {
	Tags        map[string][]string `json:"tags" bson:"tags"`                         //теги после обработки
	TagsAll     []string            `json:"tagsAll" bson:"tagsAll"`                   //все теги
	SourceRef   string              `json:"sourceRef,omitempty" bson:"sourceRef"`     //ссылка на источник
	Title       string              `json:"title,omitempty" bson:"title"`             //заголовок
	Description string              `json:"description,omitempty" bson:"description"` //описание
}
