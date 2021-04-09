package pipeline

import "reflect"

type Plugin struct {
	Name   string
	Enable bool
	Data   interface{}
}

var schemes = make(map[string]reflect.Type)

func Register(name string, c interface{}) {
	schemes[name] = reflect.TypeOf(c).Elem()
}

func (p Plugin) Pipelines() (pipeline Pipeline) {
	if t, ok := schemes[p.Name]; ok {
		pipeline = reflect.New(t).Interface().(Pipeline)
		pipeline.Init(p.Data)
		return
	} else {
		return nil
	}
}
