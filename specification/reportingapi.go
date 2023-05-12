package specification

import (
	"sort"

	"github.com/dunglas/httpsfv"
)

// ReportingEndpointsHeaderKey is the HTTP Header key for [Reporting Endpoints].
//
// [Reporting Endpoints]: https://www.w3.org/TR/reporting/#header
const ReportingEndpointsHeaderKey = "Reporting-Endpoints"

// MarshalEndpoints marshals the given map as a string to use in a [Reporting Endpoints] header.
//
// Keys will be used as endpoints names, and values as endpoints URLs.
// Note that URLs MUST be [potentially trustworthy].
// Non-secure endpoints will be ignored by clients.
//
// [Reporting Endpoints]: https://www.w3.org/TR/reporting/#header
// [potentially trustworthy]: https://w3c.github.io/webappsec-secure-contexts/#is-origin-trustworthy
func MarshalEndpoints(values map[string]string) (string, error) {
	type entry struct {
		name, url string
	}
	var es []entry
	for k, v := range values {
		es = append(es, entry{k, v})
	}
	sort.Slice(es, func(i, j int) bool { return es[i].name < es[j].name })

	dict := httpsfv.NewDictionary()
	for _, e := range es {
		dict.Add(e.name, httpsfv.NewItem(e.url))
	}
	return httpsfv.Marshal(dict)
}

type ReportsList []Report

// Report represents a [report] sent by a client.
//
// [report]: https://www.w3.org/TR/reporting/#serialize-reports
type Report struct {
	// Age is the number of milliseconds between the report’s [timestamp] and the current time according to the user agent.
	//
	// [timestamp]: https://www.w3.org/TR/reporting/#report-timestamp
	Age int `json:"timestamp"`

	// Type is a [report type].
	//
	// [report type]: https://www.w3.org/TR/reporting/#report-type
	Type string `json:"type"`

	// URL is typically the address of the Document or Worker from which the report was generated ([specification]).
	//
	// [specification]: https://www.w3.org/TR/reporting/#report-url
	URL string `json:"url"`

	// UserAgent is the value of the [User-Agent header] of the request from which the report was generated.
	//
	// [User-Agent header]: https://www.w3.org/TR/reporting/#report-user-agent
	UserAgent string `json:"user_agent"`

	// Body is the result of deserializing the report body object ([doc]).
	//
	// [doc]: https://www.w3.org/TR/reporting/#report-body
	Body any `json:"body"`
}
