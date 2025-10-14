package walk

import (
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         any
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
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{
					Age:  33,
					City: "London",
				},
				{
					Age:  34,
					City: "Reyjavik",
				},
			},
			[]string{"London", "Reyjavik"},
		},
		{
			"array",
			[2]Profile{
				{
					Age:  33,
					City: "London",
				},
				{
					Age:  34,
					City: "Reyjavik",
				},
			},
			[]string{"London", "Reyjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		want := []string{"Moo", "Baa"}
		m := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(m, func(input string) {
			got = append(got, input)
		})

		for _, v := range want {
			if !slices.Contains(got, v) {
				t.Errorf("missing %q", v)
			}
		}
	})

	t.Run("channel", func(t *testing.T) {
		profileChan := make(chan Profile)

		go func() {
			profileChan <- Profile{33, "Berlin"}
			profileChan <- Profile{34, "Katowice"}
			close(profileChan)
		}()

		want := []string{"Berlin", "Katowice"}

		var got []string
		walk(profileChan, func(s string) {
			got = append(got, s)
		})

		for _, v := range want {
			if !slices.Contains(got, v) {
				t.Errorf("missing %q", v)
			}
		}
	})
}
