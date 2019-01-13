package space

// Source: exercism/problem-specifications
// Commit: 8d4df79 space-age: Apply new "input" policy
// Problem Specifications Version: 1.1.0

var testCases = []struct {
	description string
	planet      Planet
	seconds     float64
	expected    float64
}{
	{
		description: "age on Earth",
		planet:      Planet{"Earth", 1},
		seconds:     1000000000,
		expected:    31.69,
	},
	{
		description: "age on Mercury",
		planet:      Planet{"Mercury", 0.2408467},
		seconds:     2134835688,
		expected:    280.88,
	},
	{
		description: "age on Venus",
		planet:      Planet{"Venus", 0.61519726},
		seconds:     189839836,
		expected:    9.78,
	},
	{
		description: "age on Mars",
		planet:      Planet{"Mars", 1.8808158},
		seconds:     2329871239,
		expected:    39.25,
	},
	{
		description: "age on Jupiter",
		planet:      Planet{"Jupiter", 11.862615},
		seconds:     901876382,
		expected:    2.41,
	},
	{
		description: "age on Saturn",
		planet:      Planet{"Saturn", 29.447498},
		seconds:     3000000000,
		expected:    3.23,
	},
	{
		description: "age on Uranus",
		planet:      Planet{"Uranus", 84.016846},
		seconds:     3210123456,
		expected:    1.21,
	},
	{
		description: "age on Neptune",
		planet:      Planet{"Neptune", 164.79132},
		seconds:     8210123456,
		expected:    1.58,
	},
}
