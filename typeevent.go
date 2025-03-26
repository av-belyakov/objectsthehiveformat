package objectsthehiveformat

// TypeEvent информация о событии
type TypeEvent struct {
	commonevent.CommonEventType
	Details EventCaseDetails     `json:"details,omitempty" bson:"details"` //детальная информация о событии
	Object  EventForEsCaseObject `json:"object,omitempty" bson:"object"`   //объект события
}
