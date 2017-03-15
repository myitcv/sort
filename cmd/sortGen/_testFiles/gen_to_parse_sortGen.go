// Code generated by sortGen - do not edit

package main

import "sort"
import "github.com/myitcv/sorter"

import "github.com/myitcv/sorter/cmd/sortGen/_testFiles/internal/other"
import "bytes"

func sortByName(vs []person) {
	sort.Sort(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(orderByName(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func stableSortByName(vs []person) {
	sort.Stable(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(orderByName(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func sortMySlice(vs *MySlice) *MySlice {
	theVs := vs.AsMutable()

	sort.Sort(&sorter.Wrapper{
		LenFunc: func() int {
			return theVs.Len()
		},
		LessFunc: func(i, j int) bool {
			return bool(orderMySlice(theVs, i, j))
		},
		SwapFunc: func(i, j int) {
			jPrev := theVs.Get(j)
			iPrev := theVs.Get(i)

			theVs.Set(j, iPrev)
			theVs.Set(i, jPrev)
		},
	})

	return theVs.AsImmutable(vs)
}
func stableSortMySlice(vs *MySlice) *MySlice {
	theVs := vs.AsMutable()

	sort.Stable(&sorter.Wrapper{
		LenFunc: func() int {
			return theVs.Len()
		},
		LessFunc: func(i, j int) bool {
			return bool(orderMySlice(theVs, i, j))
		},
		SwapFunc: func(i, j int) {
			jPrev := theVs.Get(j)
			iPrev := theVs.Get(i)

			theVs.Set(j, iPrev)
			theVs.Set(i, jPrev)
		},
	})

	return theVs.AsImmutable(vs)
}
func sortOtherMySlice(vs *other.MySlice) *other.MySlice {
	theVs := vs.AsMutable()

	sort.Sort(&sorter.Wrapper{
		LenFunc: func() int {
			return theVs.Len()
		},
		LessFunc: func(i, j int) bool {
			return bool(orderOtherMySlice(theVs, i, j))
		},
		SwapFunc: func(i, j int) {
			jPrev := theVs.Get(j)
			iPrev := theVs.Get(i)

			theVs.Set(j, iPrev)
			theVs.Set(i, jPrev)
		},
	})

	return theVs.AsImmutable(vs)
}
func stableSortOtherMySlice(vs *other.MySlice) *other.MySlice {
	theVs := vs.AsMutable()

	sort.Stable(&sorter.Wrapper{
		LenFunc: func() int {
			return theVs.Len()
		},
		LessFunc: func(i, j int) bool {
			return bool(orderOtherMySlice(theVs, i, j))
		},
		SwapFunc: func(i, j int) {
			jPrev := theVs.Get(j)
			iPrev := theVs.Get(i)

			theVs.Set(j, iPrev)
			theVs.Set(i, jPrev)
		},
	})

	return theVs.AsImmutable(vs)
}
func sortPointerByName(vs []*person) {
	sort.Sort(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(orderPointerByName(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func stableSortPointerByName(vs []*person) {
	sort.Stable(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(orderPointerByName(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func sortBufferByContents(vs []bytes.Buffer) {
	sort.Sort(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(orderBufferByContents(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func stableSortBufferByContents(vs []bytes.Buffer) {
	sort.Stable(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(orderBufferByContents(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func sortMap(vs []map[string]bool) {
	sort.Sort(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(orderMap(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func stableSortMap(vs []map[string]bool) {
	sort.Stable(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(orderMap(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func (e *example) sortBanana(vs []string) {
	sort.Sort(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(e.orderBanana(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
func (e *example) stableSortBanana(vs []string) {
	sort.Stable(&sorter.Wrapper{
		LenFunc: func() int {
			return len(vs)
		},
		LessFunc: func(i, j int) bool {
			return bool(e.orderBanana(vs, i, j))
		},
		SwapFunc: func(i, j int) {
			vs[i], vs[j] = vs[j], vs[i]
		},
	})
}
