package main

var secret = ""
func setSecret(sec string){
	secret = sec
}

func guess(number string) string {
	result := ""
	for i := 0; i < len(number); i++ {
		for j := 0; j < len(secret); j++ {
			if (i == j) && (number[i] == secret[j]) {
				result += "x"
				break
			} else if (i != j) && (number[i] == secret[j]) {
				result += "-"
				break
			}
		}
	}
	return result
}
