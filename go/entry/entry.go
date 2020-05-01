package entry

type Entry struct {
	Parent interface{}
	Object interface{}
}

func (e *Entry) O() interface{} {
	return e.Object
}

type Entries []*Entry

type NewObjFunc func(path string) (interface{}, error)

func NewEntry(path string, parent interface{}, newObj NewObjFunc) (*Entry, error) {
	obj, err := newObj(path)
	if err != nil {
		return nil, err
	}
	return &Entry{
		Parent: parent,
		Object: obj,
	}, nil
}
