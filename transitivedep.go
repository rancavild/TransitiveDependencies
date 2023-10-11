package transitivedep

import (
	"sort"
)

type Dependency interface {
	AddDirect(key string, dependencies []string)
	DependencyFor(key string) []string
	GetDependency() map[string][]string
}

type TransitiveDep interface {
	getDependenciesFor(key string, dependency map[string][]string) map[string]struct{}
}

type TransDep struct {}

func (t TransDep) getDependenciesFor(key string, dependency map[string][]string) map[string]struct{} {
	result := make(map[string]struct{})
	dependencies, ok := dependency[key]

	if ok {
		for _, v := range dependencies {
			result[v] = struct{}{}
			for x := range t.getDependenciesFor(v, dependency) {
				result[x] = struct{}{}
			}
		}
		return result 
	} 
	return result
}

type TransitiveDependency struct {
	dependency map[string][]string
	transitiveDep TransitiveDep
}

func (t *TransitiveDependency) AddDirect(key string, dependencies []string) {
	t.dependency[key] = dependencies
}

func (t *TransitiveDependency) DependencyFor(key string) (result []string) {
	for d := range t.transitiveDep.getDependenciesFor(key, t.dependency) {
		result = append(result, d)
	}
	sort.Strings(result)
	return
}

func (t *TransitiveDependency) GetDependency() map[string][]string {
	return t.dependency
}

func NewTransitiveDependency(transitiveDep TransitiveDep) Dependency {
	return &TransitiveDependency{dependency: make(map[string][]string), transitiveDep: transitiveDep}
}
