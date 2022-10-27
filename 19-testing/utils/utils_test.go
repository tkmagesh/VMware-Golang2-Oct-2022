package utils

import "testing"

/*
func Test_IsPrime_7(t *testing.T) {
	//Arrange
	no := 7
	expectedResult := true

	//Act
	actualResult := IsPrime(no)

	//Assert
	if actualResult != expectedResult {

		t.Errorf("Expected %t but got %t", expectedResult, actualResult)
	}
}

func Test_IsPrime_10(t *testing.T) {
	//Arrange
	no := 10
	expectedResult := false

	//Act
	actualResult := IsPrime(no)

	//Assert
	if actualResult != expectedResult {
		t.Errorf("Expected %t but got %t", expectedResult, actualResult)
	}
}
*/

//data driven test
type TestData struct {
	name     string
	no       int
	expected bool
	actual   bool
}

func Test_IsPrime(t *testing.T) {
	testCaseData := []TestData{
		{name: "Test_7_IsPrime", no: 7, expected: true},
		{name: "Test_10_IsPrime", no: 10, expected: false},
		{name: "Test_13_IsPrime", no: 13, expected: false},
		{name: "Test_15_IsPrime", no: 15, expected: false},
		{name: "Test_17_IsPrime", no: 17, expected: true},
	}
	for _, testData := range testCaseData {
		t.Run(testData.name, func(t *testing.T) {
			testData.actual = IsPrime(testData.no)
			//Assert
			if testData.actual != testData.expected {
				t.Errorf("Expected %t but got %t", testData.expected, testData.actual)
			}
		})
	}
}
