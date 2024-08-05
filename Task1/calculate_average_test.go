package main

import (
    "testing"
)

func TestCalculateAverage(t *testing.T) {
    tests := []struct {
        grades   []float32
        expected float32
    }{
        {[]float32{70, 80, 90}, 80},
        {[]float32{50, 60, 70, 80, 90}, 70},
        {[]float32{}, 0},
        {[]float32{100}, 100},
    }

    for _, test := range tests {
        result := calculateAverage(test.grades)
        if result != test.expected {
            t.Errorf("calculateAverage(%v) = %v; want %v", test.grades, result, test.expected)
        }
    }
}
