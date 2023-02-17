package set

type set struct {
	storage map[any]struct{}
}

func NewSet() *set {

	return &set{storage: map[any]struct{}{}}
}

func (s *set) Add(values ...any) {
	for _, v := range values {
		if _, ok := s.storage[v]; ok {
			continue
		}
		s.storage[v] = struct{}{}
	}
}

func (s *set) Delete(values ...any) {
	for _, v := range values {
		if _, ok := s.storage[v]; !ok {
			continue
		}
		delete(s.storage, v)
	}
}

func (s *set) Get() []any {
	res := make([]any, 0, len(s.storage))
	for k, _ := range s.storage {
		res = append(res, k)
	}
	return res
}
