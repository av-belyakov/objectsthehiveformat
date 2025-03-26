package objectsthehiveformat

// MainTypeCase это объект содержащий информацию о кейсе TheHive
type MainTypeCase struct {
	ID              string                `json:"@id" bson:"@id"`
	ElasticsearchID string                `json:"@es_id" bson:"@es_id"`
	CreateTimestamp string                `json:"@timestamp" bson:"@timestamp"`
	Source          string                `json:"source" bson:"source"`
	Event           EventMessageForEsCase `json:"event" bson:"event"`
	ObservablesMessageEs
	//TtpsMessageEs
	//поменял тип так как тип TtpsMessageEs способствует росту черезмерно большого
	//количества полей которое влечет за собой превышения лимита маппинга в
	//Elsticsearch), что выражается в ошибке от СУБД типа "Limit of total
	//fields [2000] has been exceeded while adding new fields"
	TtpsMessageTheHive
	AdditionalInformation
}
