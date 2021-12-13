package reflection

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
			"Single string input",
			struct {
				Name string
			}{"Raf"},
			[]string{"Raf"},
		},
		{
			"Empty input",
			struct{}{},
			[]string{},
		},
		{
			"Struct with two string inputs",
			struct {
				Name    string
				Company string
			}{"Raf", "InfinityWorks"},
			[]string{"Raf", "InfinityWorks"},
		},
		{
			"Struct with mixed type inputs",
			struct {
				Name string
				Age  int
			}{"Raf", 27},
			[]string{"Raf"},
		},
		{
			"Struct with nested fields",
			Person{
				"Raf",
				Profile{"Birmingham", 27},
			},
			[]string{"Raf", "Birmingham"},
		},
		{
			"Using pointers",
			&Person{
				"Raf",
				Profile{"Birmingham", 27},
			},
			[]string{"Raf", "Birmingham"},
		},
		{
			"Using slices",
			[]Profile{
				{"Birmingham", 27},
				{"London", 29},
			},
			[]string{"Birmingham", "London"},
		},
		{
			"Using arrays",
			[2]Profile{
				{"Birmingham", 27},
				{"London", 29},
			},
			[]string{"Birmingham", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := []string{}

			fn := func(s string) {
				got = append(got, s)
			}
			Walk(test.Input, fn)

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Walk calls do not match expected, got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("using maps", func(t *testing.T) {
		input := map[string]string{
			"fruits":     "Apple",
			"vegetables": "Broccoli",
		}

		var got []string

		fn := func(s string) {
			got = append(got, s)
		}
		Walk(input, fn)

		assertContains(t, got, "Apple")
		assertContains(t, got, "Broccoli")
	})

	t.Run("using channels", func(t *testing.T) {
		aChann := make(chan Profile)

		go func() {
			aChann <- Profile{"Birmingham", 27}
			aChann <- Profile{"London", 29}
			close(aChann)
		}()

		got := []string{}
		want := []string{"Birmingham", "London"}

		fn := func(s string) {
			got = append(got, s)
		}
		Walk(aChann, fn)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("using a function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{"Berlin", 33}, Profile{"Katowice", 34}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		fn := func(s string) {
			got = append(got, s)
		}
		Walk(aFunction, fn)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

func assertContains(t testing.TB, got []string, s string) {
	t.Helper()
	contains := false
	for _, val := range got {
		if val == s {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("Expected %v to contain %q, but it does not", got, s)
	}
}
