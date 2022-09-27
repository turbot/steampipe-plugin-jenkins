package jenkins

import (
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

// function which returns an ErrorPredicate for Jenkins API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		if err != nil {
			for _, item := range notFoundErrors {
				if strings.Contains(err.Error(), item) {
					return true
				}
			}
		}
		return false
	}
}
