package main

import "testing"

func TestGenerateRandomElementsSizeAndValues(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "negative size", size: -1, want: 0},
		{name: "zero size", size: 0, want: 0},
		{name: "one element", size: 1, want: 1},
		{name: "multiple elements", size: 64, want: 64},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := generateRandomElements(test.size)
			if len(data) != test.want {
				t.Fatalf("generateRandomElements(%d) returned %d elements, want %d", test.size, len(data), test.want)
			}

			for index, value := range data {
				if value <= 0 {
					t.Errorf("element at index %d = %d, want a positive integer", index, value)
				}
			}
		})
	}
}

func TestMaximumEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{7}, want: 7},
		{name: "repeated maximum", data: []int{4, 9, 2, 9, 1}, want: 9},
		{name: "maximum at the end", data: []int{2, 1, 8}, want: 8},
		{name: "negative values", data: []int{-9, -2, -5}, want: -2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum(%v) = %d, want %d", test.data, got, test.want)
			}
		})
	}
}

func TestMaxChunksEdgeCases(t *testing.T) {
	unevenData := make([]int, 2*CHUNKS+1)
	for index := range unevenData {
		unevenData[index] = index + 1
	}

	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", want: 0},
		{name: "fewer elements than chunks", data: []int{-9, -2, -5}, want: -2},
		{name: "maximum in the remainder", data: unevenData, want: len(unevenData)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maxChunks(test.data); got != test.want {
				t.Errorf("maxChunks(%v) = %d, want %d", test.data, got, test.want)
			}
		})
	}
}
