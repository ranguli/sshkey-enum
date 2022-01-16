package flagchecker

import "log"

func CheckDualFlag(option1 string, option2 string) bool {
	var dualOption bool = false

	log.Printf("Checking for two incompatible flags:\n Option 1: %s\t Option 2: %s\n", option1, option2)

	if len(option1) == 0 && len(option2) == 0 {
		log.Printf("Both options are empty.\n")
		dualOption = true
	}
	if len(option1) != 0 && len(option2) != 0 {
		log.Printf("Two incompatible flags provided\n")
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
