package caseobservables

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewObservables создаёт новый объект типа Observables
func NewObservables() *Observables {
	return &Observables{
		Observables: make(map[string][]Observable),
	}
}

func (o *Observables) GetObservables() map[string][]Observable {
	return o.Observables
}

func (o *Observables) GetKeyObservables(k string) ([]Observable, bool) {
	if value, ok := o.Observables[k]; ok {
		return value, true
	}

	return nil, false
}

func (o *Observables) SetKeyObservables(k string, observables []Observable) {
	o.Observables[k] = observables
}

// SetObservables значение для поля Observables
func (o *Observables) SetValueObservables(v map[string][]Observable) {
	o.Observables = v
}

// AddValueObservable значение для поля Observable
func (o *Observables) AddValueObservable(k string, v Observable) {
	if _, ok := o.Observables[k]; !ok {
		o.Observables[k] = []Observable(nil)
	}

	o.Observables[k] = append(o.Observables[k], v)
}

// ToStringBeautiful форматированный вывод
func (om Observables) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'observables': \n", supportingfunctions.GetWhitespace(num)))

	for key, value := range om.Observables {
		str.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), key))
		for k, v := range value {
			str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+2), k))
			str.WriteString(v.ToStringBeautiful(num + 3))
		}
	}

	return str.String()
}
