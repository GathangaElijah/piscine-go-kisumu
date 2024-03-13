package piscine

func ReverseMenuIndex(menu []string) []string {
	lengthOfMenu := len(menu)
	sortedMenu := make([]string, lengthOfMenu)
	for i := lengthOfMenu; i >= 1; i-- {
		sortedMenu = make([]string, i)
	}
	return sortedMenu
}
