package specification_test

import (
	"testing"

	. "github.com/empijei/SecurityCollector/specification"
)

func TestDeprecatedMarshalGroups(t *testing.T) {
	got := DeprecatedMarshalGroups(
		NewDeprecatedGroup("foo", "https://deprecated.empijei.science", "https://other.empijei.science"),
		NewDeprecatedGroup("bar", "https://another.empijei.science"),
	)
	want := `{"group":"foo","max_age":432000,"endpoints":[{"url":"https://deprecated.empijei.science"},{"url":"https://other.empijei.science"}]},{"group":"bar","max_age":432000,"endpoints":[{"url":"https://another.empijei.science"}]}`
	if got != want {
		t.Errorf(`DeprecatedMarshalGroups: got %v, want %v`, got, want)
	}
}
