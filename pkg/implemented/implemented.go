package implemented

func DeleteFunc(slice []string, i int) []string { //deletes fixed element in slice
	//it is done for further implementations
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}
