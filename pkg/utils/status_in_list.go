package utils

// StatusInList checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		return i == status
	}
	return false
}
