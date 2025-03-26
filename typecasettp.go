package objectsthehiveformat

// Ttps список TTP сообщений
type Ttps struct {
	Ttp []TtpMessage `json:"ttp" bson:"ttp"`
}

// TtpMessage TTP сообщения
type TtpMessage struct {
	OccurDate            string              `json:"occurDate,omitempty" bson:"occurDate"`   //формат RFC3339 дата возникновения
	UnderliningCreatedAt string              `json:"_createdAt,omitempty" bson:"_createdAt"` //формат RFC3339 время создания
	UnderliningId        string              `json:"_id,omitempty" bson:"_id"`               //уникальный идентификатор
	UnderliningCreatedBy string              `json:"_createdBy,omitempty" bson:"_createdBy"` //кем создан
	PatternId            string              `json:"patternId,omitempty" bson:"patternId"`   //уникальный идентификатор шаблона
	Tactic               string              `json:"tactic,omitempty" bson:"tactic"`         //тактика
	ExtraData            ExtraDataTtpMessage `json:"extraData" bson:"extraData"`             //дополнительные данные
}

// ExtraDataTtpMessage дополнительные данные TTP сообщения
type ExtraDataTtpMessage struct {
	Pattern       PatternExtraData `json:"pattern,omitempty" bson:"pattern"`             //шаблон
	PatternParent PatternExtraData `json:"patternParent,omitempty" bson:"patternParent"` //родительский шаблон
}

// PatternExtraData шаблон дополнительных данных
type PatternExtraData struct {
	RemoteSupport        bool     `json:"remoteSupport,omitempty" bson:"remoteSupport"`   //удаленная поддержка
	Revoked              bool     `json:"revoked,omitempty" bson:"revoked"`               //аннулированный
	UnderliningCreatedAt string   `json:"_createdAt,omitempty" bson:"_createdAt"`         //формат RFC3339 время создания
	UnderliningCreatedBy string   `json:"_createdBy,omitempty" bson:"_createdBy"`         //кем создан
	UnderliningId        string   `json:"_id,omitempty" bson:"_id"`                       //уникальный идентификатор
	UnderliningType      string   `json:"_type,omitempty" bson:"_type"`                   //тип
	Description          string   `json:"description,omitempty" bson:"description"`       //описание
	Detection            string   `json:"detection,omitempty" bson:"detection"`           //обнаружен
	Name                 string   `json:"name,omitempty" bson:"name"`                     //наименование
	PatternId            string   `json:"patternId,omitempty" bson:"patternId"`           //уникальный идентификатор шаблона
	PatternType          string   `json:"patternType,omitempty" bson:"patternType"`       //тип шаблона
	URL                  string   `json:"url,omitempty" bson:"url"`                       //URL
	Version              string   `json:"version,omitempty" bson:"version"`               //версия
	Platforms            []string `json:"platforms" bson:"platforms"`                     //список платформ
	PermissionsRequired  []string `json:"permissionsRequired" bson:"permissionsRequired"` //требуемые разрешения
	DataSources          []string `json:"dataSources" bson:"dataSources"`                 //источники данных
	Tactics              []string `json:"tactics" bson:"tactics"`                         //тактики
	//данное поле редко используемое, думаю пока оно не требует реализации
	//DefenseBypassed     []string               `json:"defenseBypassed"` //надо проверить тип
	//данное поле редко используемое, думаю пока оно не требует реализации
	//SystemRequirements  []string               `json:"systemRequirements"` //надо проверить тип
	//данное поле редко используемое, думаю пока оно не требует реализации
	//ExtraData           map[string]interface{} `json:"extraData"`
}
