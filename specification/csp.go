package specification

const CSPReportType = "csp-violation"

// CSPReport is an implementation for the [deprecated serialization] for reports.
//
// [deprecated serialization]: https://www.w3.org/TR/CSP3/#deprecated-serialize-violation
type CSPReportDeprecated struct {
	DocumentURI string `json:"document-uri"`

	Referrer string `json:"referrer"`

	BlockedURI string `json:"blocked-uri"`

	EffectiveDirective string `json:"effective-directive"`

	ViolatedDirective string `json:"violated-directive"`

	OriginalPolicy string `json:"original-policy"`

	Disposition string `json:"disposition"`

	StatusCode uint `json:"status-code"`

	ScriptSample string `json:"script-sample"`
}

// CSPReport will be an implementation for the [serialization] of reports.
//
// The specification is currently too unclear to implement this bit, so this doesn't currently work.
//
// [serialization]: https://www.w3.org/TR/CSP3/#reporting
type CSPReport struct {
	/*
		DocumentURL string `json:"document-uri TODO"`

		Referrer string `json:"referrer"`

		BlockedURL string `json:"blocked-uri TODO"`

		EffectiveDirective string `json:"effective-directive"`

		OriginalPolicy string `json:"original-policy"`

		SourceFile string `json:"source-file"`

		Sample string `json:"sample"`

		Disposition string `json:"disposition"`

		StatusCode uint `json:"status-code"`

		LineNumber uint `json:"line-number"`

		ColumnNumber uint `json:"column-number"`
	*/
}
