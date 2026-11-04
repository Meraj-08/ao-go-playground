package format

import "testing"

func TestTags(t *testing.T) {
	cases := []struct {
		name string
		tags []string
		want string
	}{
		{"none", nil, ""},
		{"one", []string{"work"}, "#work"},
		{"several", []string{"work", "home"}, "#work #home"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Tags(c.tags)
			if err != nil {
				t.Fatal(err)
			}
			if got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestTagsRejectsBlankTag(t *testing.T) {
	for _, blank := range []string{"", "   "} {
		_, err := Tags([]string{"work", blank})
		if err == nil {
			t.Fatalf("expected an error for blank tag %q", blank)
		}
		if want := "tag 2 is blank"; err.Error() != want {
			t.Fatalf("got error %q, want %q", err, want)
		}
	}
}
