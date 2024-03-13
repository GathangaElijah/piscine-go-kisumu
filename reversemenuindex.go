package piscine

func ReverseMenuIndex(menu []string) []string {
	lengthOfMenu := len(menu)
	sortedMenu := make([]string, lengthOfMenu)
	for i, menuItem := range menu {
		reverseIndex := lengthOfMenu - 1 - i
		sortedMenu[reverseIndex] = menuItem
	}
	return sortedMenu
}
