package cerror

// CheckErr output message of errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}