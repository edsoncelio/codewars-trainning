package main
import (
	"testing"
)
// usage: go test, using cover: go test -cover

func TestIsTriangle(t *testing.T){
	result := IsTriangle(2,3,4)
	if result != true {
		t.Errorf("Expected:  %d, got: %d",true, result)
	} 
}

func TestRemoveChar(t *testing.T){
	word string = testing
	result := RemoveChar("testing")
	if result != word[1:(len(word)-1)] {
		t.Errorf("Expected: %s, got: %s", word[1:(len(word)-1)], result)
	}
}

func TestEvenOrOdd(t *testing.T){
	result := EvenOrOdd(25)
	if result != "Odd"{
		t.Errorf("Expected: %s, got: %s", "Odd", result)
	}
}

func TestPositiveSum(t *testing.T){
	result := PositiveSum(1,2,3,-4)
	if result != 2{
		t.Errorf("Expected: %d, got: %d", 2, result)
	}
}

func TestNegative(t *testing.T){
	result := Negative(-2)
	if result != 2 {
		t.Errorf("Expected: %d, got: %d",2, result)
	}
}
