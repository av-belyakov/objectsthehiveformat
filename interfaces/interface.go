package interfaces

type CustomerFields interface {
	//принимает значения где первое значение метода это первое значение в структуре данных,
	// второе значение метода это второе значение в структуре данных
	Set(oneStructField any, twoStructField any)
	//возвращает значения где 1 и 3 значение это наименование поля
	Get() (fieldNameOne string, valueOne int, fieldNameTwo string, valueTwo string)
}
