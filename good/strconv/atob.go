//line :1
package zBESu2SrRjP




func IoVSFa90tanT(jREhFXLW3 string) (bool, error) {
	switch jREhFXLW3 {
	case "1", "t", "T", "true", "TRUE", "True":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False":
		return false, nil
	}
	return false,  /*line :1*/ nF25LgNzv("ParseBool", jREhFXLW3)
}


func D4Oadm(o7OjhYc5wQns bool) string {
	if o7OjhYc5wQns {
		return "true"
	}
	return "false"
}



func AbnOjTnABoo(l_A2Ytb []byte, o7OjhYc5wQns bool) []byte {
	if o7OjhYc5wQns {
		return  /*line :1*/ append(l_A2Ytb, "true"...)
	}
	return  /*line :1*/ append(l_A2Ytb, "false"...)
}
