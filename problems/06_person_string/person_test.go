package person

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		name string
		person Person
		want string
	}{
		{
			name:   "normal case",
			person: Person{FirstName: "John", LastName: "Doe", Age: 30},
			want:   "John Doe is 30 years old.",
		},
		{
			name:   "empty names",
			person: Person{FirstName: "", LastName: "", Age: 25},
			want:   "  is 25 years old.",
		},
		{
			name:   "age is zero",
			person: Person{FirstName: "Alice", LastName: "Smith", Age: 0},
			want:   "Alice Smith is 0 years old.",
		},
		{
			name:   "Japanese name",
			person: Person{FirstName: "太郎", LastName: "山田", Age: 20},
			want:   "太郎 山田 is 20 years old.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.String()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
