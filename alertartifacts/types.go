package alertartifacts

import "github.com/av-belyakov/objectsthehiveformat/common"

// Artifacts содержит артефакт к алерту
type Artifacts struct {
	Tags           map[string][]string `json:"tags" bson:"tags"`                               //теги после обработки
	SnortSid       []string            `json:"snortSid,omitempty" bson:"snortSid"`             //список snort сигнатур (строка)
	TagsAll        []string            `json:"tagsAll" bson:"tagsAll"`                         //все теги
	SnortSidNumber []int               `json:"SnortSidNumber,omitempty" bson:"SnortSidNumber"` //список snort сигнатур (число)
	SensorId       string              `json:"sensorId,omitempty" bson:"sensorId"`             //сенсор id
	common.CommonArtifactType
}
