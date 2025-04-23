package common

import "github.com/av-belyakov/objectsthehiveformat/interfaces"

// CommonEventType общие поля для описания события
type CommonEventType struct {
	StartDate      string `json:"startDate,omitempty" bson:"startDate"`           //начальная дата в формате RFC3339
	RootId         string `json:"rootId,omitempty" bson:"rootId"`                 //основной уникальный идентификатор
	Organisation   string `json:"organisation,omitempty" bson:"organisation"`     //наименование организации
	OrganisationId string `json:"organisationId,omitempty" bson:"organisationId"` //уникальный идентификатор организации
	ObjectId       string `json:"objectId,omitempty" bson:"objectId"`             //уникальный идентификатор объекта
	ObjectType     string `json:"objectType,omitempty" bson:"objectType"`         //тип объекта
	Operation      string `json:"operation,omitempty" bson:"operation"`           //операция
	RequestId      string `json:"requestId,omitempty" bson:"requestId"`           //уникальный идентификатор запроса
	Base           bool   `json:"base,omitempty" bson:"base"`                     //основа
}

// CustomFields настраиваемые поля
type CustomFields map[string]interfaces.CustomerFields

type CustomFieldStringType struct {
	String string `json:"string" bson:"string"`
	Order  int    `json:"order" bson:"order"`
}

type CustomFieldDateType struct {
	Date  string `json:"date" bson:"date"`
	Order int    `json:"order" bson:"order"`
}

type CustomFieldIntegerType struct {
	Order   int `json:"order" bson:"order"`
	Integer int `json:"integer" bson:"integer"`
}

type CustomFieldBoolenType struct {
	Order   int  `json:"order" bson:"order"`
	Boolean bool `json:"boolen" bson:"boolen"`
}

// CommonEventCaseObject общие поля для описания объекта события
type CommonEventCaseObject struct {
	StartDate        string `json:"startDate,omitempty" bson:"startDate"`               //начальная дата в формате RFC3339
	EndDate          string `json:"endDate,omitempty" bson:"endDate"`                   //конечная дата в формате RFC3339
	CreatedAt        string `json:"createdAt,omitempty" bson:"createdAt"`               //дата создания в формате RFC3339
	UpdatedAt        string `json:"updatedAt,omitempty" bson:"updatedAt"`               //дата обновления в формате RFC3339
	UnderliningId    string `json:"_id,omitempty" bson:"_id"`                           //уникальный идентификатор
	Id               string `json:"id,omitempty" bson:"id"`                             //уникальный идентификатор
	CreatedBy        string `json:"createdBy,omitempty" bson:"createdBy"`               //кем создан
	UpdatedBy        string `json:"updatedBy,omitempty" bson:"updatedBy"`               //кем обновлен
	UnderliningType  string `json:"_type,omitempty" bson:"_type"`                       //тип
	Title            string `json:"title,omitempty" bson:"title"`                       //заголовок
	Description      string `json:"description,omitempty" bson:"description"`           //описание
	ImpactStatus     string `json:"impactStatus,omitempty" bson:"impactStatus"`         //краткое описание воздействия
	ResolutionStatus string `json:"resolutionStatus,omitempty" bson:"resolutionStatus"` //статус разрешения
	Status           string `json:"status,omitempty" bson:"status"`                     //статус
	Summary          string `json:"summary,omitempty" bson:"summary"`                   //резюме
	Owner            string `json:"owner,omitempty" bson:"owner"`                       //владелец
	CaseId           uint64 `json:"caseId,omitempty" bson:"caseId"`                     //уникальный идентификатор дела
	Severity         uint64 `json:"severity,omitempty" bson:"severity"`                 //строгость
	Tlp              uint64 `json:"tlp,omitempty" bson:"tlp"`                           //tlp
	Pap              uint64 `json:"pap,omitempty" bson:"pap"`                           //pap
	Flag             bool   `json:"flag,omitempty" bson:"flag"`                         //флаг
}

// CommonObservableType общие поля наблюдаемого сообщения
type CommonObservableType struct {
	UnderliningCreatedAt string `json:"_createdAt,omitempty" bson:"_createdAt"`             //время создания в формате RFC3339
	UnderliningUpdatedAt string `json:"_updatedAt,omitempty" bson:"_updatedAt"`             //время обновления в формате RFC3339
	StartDate            string `json:"startDate,omitempty" bson:"startDate"`               //дата начала в формате RFC3339
	UnderliningCreatedBy string `json:"_createdBy,omitempty" bson:"_createdBy"`             //кем создан
	UnderliningUpdatedBy string `json:"_updatedBy,omitempty" bson:"_updatedBy"`             //кем обновлен
	UnderliningId        string `json:"_id,omitempty" bson:"_id"`                           //уникальный идентификатор
	UnderliningType      string `json:"_type,omitempty" bson:"_type"`                       //тип
	Data                 string `json:"data,omitempty" bson:"data"`                         //данные
	DataType             string `json:"dataType,omitempty" bson:"dataType"`                 //тип данных
	Message              string `json:"message,omitempty" bson:"message"`                   //сообщение
	Tlp                  uint64 `json:"tlp,omitempty" bson:"tlp"`                           //tlp
	Ioc                  bool   `json:"ioc,omitempty" bson:"ioc"`                           //индикатор компрометации
	Sighted              bool   `json:"sighted,omitempty" bson:"sighted"`                   //видящий
	IgnoreSimilarity     bool   `json:"ignoreSimilarity,omitempty" bson:"ignoreSimilarity"` //игнорировать сходство
}

// CommonEventAlertObject объект события для Alert
type CommonEventAlertObject struct {
	Id              string `json:"id,omitempty" bson:"id"`                               //уникальный идентификатор
	Date            string `json:"date,omitempty" bson:"date"`                           //дата (формат RFC3339)
	Case            string `json:"case,omitempty" bson:"case,omitempty"`                 //кейс
	Type            string `json:"type,omitempty" bson:"type"`                           //тип
	Title           string `json:"title,omitempty" bson:"title"`                         //заголовок
	Status          string `json:"status,omitempty" bson:"status"`                       //список тегов
	Source          string `json:"source,omitempty" bson:"source"`                       //источник
	CreatedBy       string `json:"createdBy,omitempty" bson:"createdBy"`                 //кем создан
	UpdatedBy       string `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`       //кем обновлен
	CreatedAt       string `json:"createdAt,omitempty" bson:"createdAt"`                 //дата создания (формат RFC3339)
	UpdatedAt       string `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`       //дата обновления (формат RFC3339)
	SourceRef       string `json:"sourceRef,omitempty" bson:"sourceRef"`                 //ссылка на источник
	ObjectType      string `json:"objectType,omitempty" bson:"objectType"`               //тип объекта
	Description     string `json:"description,omitempty" bson:"description"`             //описание
	CaseTemplate    string `json:"caseTemplate,omitempty" bson:"caseTemplate,omitempty"` //шаблон обращения
	UnderliningId   string `json:"_id,omitempty" bson:"_id"`                             //уникальный идентификатор
	UnderliningType string `json:"_type,omitempty" bson:"_type"`                         //тип
	Tlp             uint64 `json:"tlp,omitempty" bson:"tlp"`                             //tlp
}

// AttachmentData прикрепленные данные
type AttachmentData struct {
	Hashes      []string `json:"hashes,omitempty" bson:"hashes"`           //список хешей
	Id          string   `json:"id,omitempty" bson:"id"`                   //идентификатор
	Name        string   `json:"name,omitempty" bson:"name"`               //наименование
	ContentType string   `json:"contentType,omitempty" bson:"contentType"` //тип контента
	Size        uint64   `json:"size,omitempty" bson:"size"`               //размер
}

// ReportTaxonomies
type ReportTaxonomies struct {
	Taxonomies []Taxonomy `json:"taxonomies,omitempty" bson:"taxonomies"`
}

// Taxonomy
type Taxonomy struct {
	Level     string `json:"level,omitempty" bson:"level"`
	Namespace string `json:"namespace,omitempty" bson:"namespace"`
	Predicate string `json:"predicate,omitempty" bson:"predicate"`
	Value     string `json:"value,omitempty" bson:"value"`
}

// CommonArtifactType содержит общие поля для описания объекта Artifact
type CommonArtifactType struct {
	Id              string `json:"id,omitempty" bson:"id"`               //уникальный идентификатор
	Data            string `json:"data,omitempty" bson:"data"`           //данные
	Message         string `json:"message,omitempty" bson:"message"`     //сообщение
	DataType        string `json:"dataType,omitempty" bson:"dataType"`   //тип данных
	CreatedAt       string `json:"createdAt,omitempty" bson:"createdAt"` //время создания в формате RFC3339
	CreatedBy       string `json:"createdBy,omitempty" bson:"createdBy"` //кем создан
	StartDate       string `json:"startDate,omitempty" bson:"startDate"` //дата начала в формате RFC3339
	UnderliningId   string `json:"_id,omitempty" bson:"_id"`             //уникальный идентификатор
	UnderliningType string `json:"_type,omitempty" bson:"_type"`         //тип
	Tlp             uint64 `json:"tlp,omitempty" bson:"tlp"`             //tlp
	Ioc             bool   `json:"ioc,omitempty" bson:"ioc"`             //индикатор компрометации
}
