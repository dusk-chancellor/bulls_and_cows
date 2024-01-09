package implemented

func DeleteFunc(s []string, i int) []string { //deletes fixed element in slice
	//it is done for further implementations
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	return s
}
