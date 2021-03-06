package tcmp

// T just for generic mode
type T interface{}

type Slice struct{}

// cannot use generic T in normal mode
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

func (m *Map) Equal(t1, t2 map[T]T) bool {
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

func (m *Map) IsNil(t map[T]T) bool {
    if t == nil {
        return true
    }
    return false
}
