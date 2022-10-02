package main

import "testing"


func Sum(a, b int) int{
	return a + b
}
func TestSum(t *testing.T) {
	t.Run("Sum of two numbers", func(t *testing.T) {
		got := Sum(6, 4)
	want := 10

	assertCorrectMessage(t, got, want)
	})
}





func assertCorrectMessage(t testing.TB, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}