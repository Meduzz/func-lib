package annotation

import "fmt"

type (
	Annotation string
)

func Annotations(annotations ...Annotation) []Annotation {
	return annotations
}

func NoAnnotations() []Annotation {
	return make([]Annotation, 0)
}

func FromString(annotation string) Annotation {
	return Annotation(annotation)
}

func KeyValue(key, value string) Annotation {
	return Annotation(fmt.Sprintf("%s=%s", key, value))
}

func Name(name string) Annotation {
	return KeyValue("name", name)
}

func KeyIntValue(key string, value int) Annotation {
	return Annotation(fmt.Sprintf("%s=%d", key, value))
}

func KeyBooleanValue(key string, value bool) Annotation {
	return Annotation(fmt.Sprintf("%s=%v", key, value))
}

func Type(value string) Annotation {
	return KeyValue("type", value)
}
