package reflecion

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{

			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with non string field",
			struct {
				Name  string
				Value int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		// For this last test, we have created two types
		// This helps greatly for readability
		{
			"struct with nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"pointer to a struct",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		}, {
			"slices",
			[]Profile{
				{22, "London"},
				{33, "Paris"},
				{44, "Bruxelles"},
			},
			[]string{"London", "Paris", "Bruxelles"},
		},
		{
			"arrays",
			[2]Person{
				{"Chris", Profile{33, "London"}},
				{"John", Profile{44, "Paris"}},
			},
			[]string{"Chris", "London", "John", "Paris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})

	}
	t.Run("with maps", func(t *testing.T) {

		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string

		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "London"}
			aChannel <- Profile{44, "Bruxelles"}
			close(aChannel)
		}()

		var got []string
		want := []string{"London", "Bruxelles"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{44, "Brussels"}
		}

		var got []string
		want := []string{"Berlin", "Brussels"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it did not", haystack, needle)
	}

}
