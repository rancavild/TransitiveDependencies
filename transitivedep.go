package transitivedep

import (
	"sort"
)

type Dependency interface {
	AddDirect(key string, dependencies []string)
	DependencyFor(key string) []string
}

type TransitiveDependency struct {
	dependency map[string][]string	
}

func (t TransitiveDependency) AddDirect(key string, dependencies []string) {
	t.dependency[key] = dependencies
}

func (t TransitiveDependency) getDependencyFor(key string) map[string]struct{} {
	result := make(map[string]struct{})
	dependencies, ok := t.dependency[key]

	if ok {
		for _, v := range dependencies {
			result[v] = struct{}{}
			for x := range t.getDependencyFor(v) {
				result[x] = struct{}{}
			}
		}
		return result 
	} 
	return result
}

func (t TransitiveDependency) DependencyFor(key string) (result []string) {
	for d := range t.getDependencyFor(key) {
		result = append(result, d)
	}
	sort.Strings(result)
	return
}

func NewTransitiveDependency() Dependency {
	return TransitiveDependency{dependency: make(map[string][]string)}
}
