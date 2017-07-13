package avito

import (
	"fmt"
	"strconv"
	"strings"
)

// GetIDFromURL returns id by given url
func GetIDFromURL(u string) (int, error) {
	arr := strings.Split(u, "_")
	if len(arr) < 2 {
		return 0, fmt.Errorf("Bad url")
	}
	return strconv.Atoi(arr[len(arr)-1])
}
