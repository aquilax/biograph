package biograph

type MetaData map[string]string

func (md MetaData) merge(md2 *MetaData) {
	for k, v := range *md2 {
		md[k] = v
	}
}

func (md MetaData) get(key string) string {
	result, ok := md[key]
	if ok {
		return result
	}
	return ""
}
