package biograph

type MetaData map[string]string

func (md MetaData) merge(md2 *MetaData) {
	for k, v := range *md2 {
		md[k] = v
	}
}

func (md MetaData) Get(key string) string {
	result, ok := md[key]
	if ok {
		return result
	}
	return ""
}

// Keys returns array with all meta data keys
func (md MetaData) Keys() []string {
	keys := make([]string, len(md))
	i := 0
	for k := range md {
		keys[i] = k
		i++
	}
	return keys
}
