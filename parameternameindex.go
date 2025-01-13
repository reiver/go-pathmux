package pathmux

func parameterNameIndex(name string, parameterNames ...string) (int, bool) {
	for index, parameterName := range parameterNames {
		if parameterName == name {
			return index, true
		}
	}

	var nada int
	return nada, false
}
