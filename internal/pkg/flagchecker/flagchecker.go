package flagchecker

func CheckDualFlag(option1 string, option2 string) bool {
	var dualOption bool = false

	if len(option1) == 0 && len(option2) == 0 {
		dualOption = true
	}
	if len(option1) != 0 && len(option2) != 0 {
		dualOption = true
	}

	return dualOption
}

func CheckFlag(option1 string) bool {
	var flag bool = false

	if len(option1) != 0 {
		flag = true
	}

	return flag
}
