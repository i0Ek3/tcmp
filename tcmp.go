package tcmp

type T interface {
    //[]T
    int, int8, int16, int32, int64, string
}

type TK interface {
    T
}

type TV interface {
    T
}

type Slice struct {}

func (s *Slice) Equal(t1, t2 []T) bool {
    if len(t1) != len(t2) {
        return false
    }
    for i := range t1 {
        if t1[i] != t2[i] {
            return false
        }
    }
    return true
}

func (s *Slice) IsNil(t []T) bool {
    if t == nil {
        return true
    }
    return false
}

type Map struct{}

func (m *Map) Equal(t1, t2 map[TK]TV) bool {
    if len(t1) != len(t2) {
        return false
    }
    for k, t1v := range t1 {
        if t2v, ok := t2[k]; !ok || t1v != t2v {
            return false
        }
    }
    return true
}

func (m *Map) IsNil(t map[TK]TV) bool {
    if t == nil {
        return true
    }
    return false
}
