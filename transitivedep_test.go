package transitivedep

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDependencyFor(t *testing.T) {
	testCases := []struct{
		expected []string
		given string
	}{
		{[]string{"B","C","E","F","G","H"}, "A"},
		{[]string{"C","E","F","G","H"}, "B"},
		{[]string{"G"}, "C"},
		{[]string{"A","B","C","E","F","G","H"}, "D"},
		{[]string{"F", "H"}, "E"},
		{[]string{"H"}, "F"},
	}
	td := NewTransitiveDependency(Trans{})
        td.AddDirect("A", []string{"B", "C"})
        td.AddDirect("B", []string{"C", "E"})
        td.AddDirect("C", []string{"G"})
        td.AddDirect("D", []string{"A", "F"})
        td.AddDirect("E", []string{"F"})
        td.AddDirect("F", []string{"H"})

	for i, testCase := range testCases {
		testName := fmt.Sprintf("Test-DependencyFor-%d",i)

		t.Run(testName, func(t *testing.T) {
			actual := td.DependencyFor(testCase.given)
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("Failed : expected %v, actual %v",testCase.expected, actual)
			}										
		})
	}
}
