package specification

import (
	"encoding/json"
	"io"
	"strings"
)

const (
	// DeprecatedReportToDefaultMaxAge is used as default cache duration for report groups.
	DeprecatedReportToDefaultMaxAge = 5 * 24 * 60 * 60

	// DeprecatedReportToHeaderKey is the HTTP header key for the Reporting API.
	DeprecatedReportToHeaderKey = "Report-To"
)

// DeprecatedEndpoint is the Go representation of [endpoints] in the deprecated draft for Report-To headers format.
//
// [deprecated draft]: https://www.w3.org/TR/2018/WD-reporting-1-20180925/#concept-endpoints
type DeprecatedEndpoint struct {
	// URL defines the location of the endpoint.
	URL string `json:"url"`
	// Priority forms failover classes.
	// Failover classes allow the developer to provide backup collectors (those with higher priority values)
	// that will only receive reports if all of the primary collectors (those with lower priority values) fail.
	Priority uint `json:"priority,omitempty"`
	// Weight determines how report traffic is balanced across the failover class.
	Weight uint `json:"weight,omitempty"`
}

// DeprecatedGroup is the Go representation of [groups] in the deprecated draft Report-To headers format.
//
// [groups]: https://www.w3.org/TR/2018/WD-reporting-1-20180925/#id-member
type DeprecatedGroup struct {
	// Name associates a name with the endpoint group.
	// If no member named "group" is present in the object,
	// the endpoint group will be given the name "default".
	Name string `json:"group,omitempty"`
	// IncludeSubdomains enables this endpoint group for all subdomains of the current origin’s host.
	// If no member named "include_subdomains" is present in the object, or its value is not "true",
	// the endpoint group will not be enabled for subdomains
	IncludeSubdomains bool `json:"include_subdomains,omitempty"`
	// MaxAgeSeconds defines the endpoint group’s lifetime, as a non-negative integer number of seconds.
	// A value of 0 will cause the endpoint group to be removed from the user agent’s reporting cache.
	MaxAgeSeconds uint `json:"max_age"`
	// Endpoints is the list of endpoints that belong to this endpoint group.
	Endpoints []DeprecatedEndpoint `json:"endpoints"`
}

// NewDeprecatedGroup creates a new Group with MaxAge set to DeprecatedReportToDefaultMaxAge
// and all optional values with increasing priority.
func NewDeprecatedGroup(name string, url string, otherUrls ...string) DeprecatedGroup {
	es := []DeprecatedEndpoint{{URL: url}}
	for i, u := range otherUrls {
		es = append(es, DeprecatedEndpoint{URL: u, Priority: uint(i)})
	}
	return DeprecatedGroup{
		Name:          name,
		MaxAgeSeconds: DeprecatedReportToDefaultMaxAge,
		Endpoints:     es,
	}
}

// DeprecatedMarshalGroups returns the wire representation of a list of groups.
func DeprecatedMarshalGroups(groups ...DeprecatedGroup) string {
	var sb strings.Builder
	for i, g := range groups {
		if i >= 1 {
			io.WriteString(&sb, ",")
		}
		buf, err := json.Marshal(g)
		if err != nil {
			// This doesn't happen because we know that DeprecatedGroup is serializable.
			panic(err)
		}
		sb.Write(buf)
	}
	return sb.String()
}
