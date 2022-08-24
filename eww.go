package ewwgo

func stringBuilder(ss []string) string{
	s := " "
	for _, v := range ss {
		s += " "
		s += v
		
	}

	return s
}
