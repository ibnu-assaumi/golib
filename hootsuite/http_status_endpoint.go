package hootsuite

import (
	"fmt"
	"strings"

	"github.com/hootsuite/healthchecks"
	"github.com/hootsuite/healthchecks/checks/httpsc"
)

const (
	httpURLType  string = "http"
	httpsURLType string = "http"
)

// GetHTTPStatusEndpoint : get hootsuite healtchecks status endpoint
// required hootsuite healthchecks implementation to use this, see : https://github.com/hootsuite/healthchecks
func GetHTTPStatusEndpoint(name, slug, url, urlType string, isTraversable bool) (statusEndpoint healthchecks.StatusEndpoint, err error) {
	if strings.TrimSpace(name) == "" {
		return statusEndpoint, fmt.Errorf("parameter 'name' is required")
	}
	if strings.ContainsAny(slug, " ") {
		return statusEndpoint, fmt.Errorf("parameter 'slug' should not contain whitespaces")
	}
	if slug == "" {
		return statusEndpoint, fmt.Errorf("parameter 'slug' is required")
	}
	if strings.TrimSpace(url) == "" {
		return statusEndpoint, fmt.Errorf("parameter 'url' is required")
	}
	if urlType != httpURLType && urlType != httpsURLType {
		return statusEndpoint, fmt.Errorf(fmt.Sprintf("parameter 'urlType' should be %s or %s ", httpURLType, httpsURLType))
	}
	statusEndpoint = healthchecks.StatusEndpoint{
		Name:          name,
		Slug:          slug,
		Type:          urlType,
		StatusCheck:   httpsc.HttpStatusChecker{BaseUrl: url},
		IsTraversable: isTraversable,
	}

	if isTraversable {
		statusEndpoint.TraverseCheck = httpsc.HttpStatusChecker{BaseUrl: url}
	}

	return statusEndpoint, nil
}
