package decrypt

func Dec(ch string) string {
	//init string
	strDecryp :=""
	//loop string
	for _,c := range ch {
		//codeAscii of each caractere of my string
		codeAscii := int(c)
		// convert ascii to string
		cToSting := string(codeAscii-3)
		//build encrypt str
		strDecryp += cToSting	
	}
	return strDecryp
}