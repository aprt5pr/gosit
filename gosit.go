package gosit

import "fmt"
import "strings"

type scale struct {
	name, scale_type, full_name string
	notes                       [8]string
}

func IsValidNoteName(note string) bool {
	n := strings.ToUpper(string(note[0]))
	valid_notes := [7]string{"A", "B", "C", "D", "E", "F", "G"}

	for idx := range valid_notes {
		if n == valid_notes[idx] {
			return true
		}
	}
	return false
}

func ResolveNoteName(note string) string {
	// If note doesn't have a sign - assuem natural
	if len(note) == 1 {
		return strings.ToUpper(note)
	} else if len(note) == 2 {
		valid_signs := [2]string{"#", "b"}
		sign := string(note[1])
		for idx := range valid_signs {
			if sign == valid_signs[idx] {
				break
			}
			// Reached end of array with no match
			if idx == 1 {
				return "?"
			}
		}
		// e.g. c-sharp is also d-flat
		switch strings.ToLower(note) {
		case "a#":
			return "A#/Bb"
		case "bb":
			return "A#/Bb"
		case "c#":
			return "C#/Db"
		case "db":
			return "C#/Db"
		case "d#":
			return "D#/Eb"
		case "eb":
			return "D#/Eb"
		case "f#":
			return "F#/Gb"
		case "gb":
			return "F#/Gb"
		case "g#":
			return "G#/Ab"
		case "ab":
			return "G#/Ab"
		}
	} else {
		return "?"
	}
	return "?"
}

func NewScale(name, s_type string) *scale {
	s := scale{name: name, scale_type: s_type}
	s.full_name = s.name + s_type

	// All the possible notes
	notes := [12]string{"A", "A#/Bb", "B", "C", "C#/Db", "D", "D#/Eb", "E", "F", "F#/Gb", "G", "G#/Ab"}
	for idx := range notes {
		// Get index of starting point in scale
		if notes[idx] == s.name {
			start_at := idx

			// Step sizes in the major scale relative to root tone
			// +1 == half step
			// +2 == whole step
			// whole, whole, half, whole, whole, whole
			scale_steps := [6]int{2, 4, 5, 7, 9, 11}

			// Find tones 2 thru 7 i.e. elements 1 thru 6 in `notes` array
			for step_idx := range scale_steps {
				step := scale_steps[step_idx]
				if start_at+step <= 11 {
					// We add `1` to s.notes[x] to offset for element 0 which we already know
					// and will add later
					s.notes[step_idx+1] = notes[start_at+step]
				} else {
					// We add `1` to s.notes[x] to offset for element 0 which we already know
					// and will add later
					// Start over when end of array is reached by subtracting max num elements `12`
					s.notes[step_idx+1] = notes[start_at+step-12]
				}
			}
			// Do not check any further notes after match is found
			break
		}
	}
	s.notes[0] = name
	s.notes[7] = s.notes[0]
	return &s
}
