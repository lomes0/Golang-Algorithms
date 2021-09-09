package main

import (
	"math/rand"
	"reflect"
	//"sort"
	"sync"
)

const (
	SIZE = (1024)
)

type Arr []int

/*
* TODO::Impl properly
 */
func weak_partition(arr []int, i int, j int, piv int) int {

	sort.Ints(arr[i : j+1])
	for l := j; l >= -1; l-- {
		if arr[l] == piv {
			return l
		}
	}

	return i
}

func pivots2right(arr []int, piv int) int {

	swapf := reflect.Swapper(arr)

	var smt queue = queue{make([]int, 0)} // smaller than piv queue

	i := len(arr) - 1
	lm := 0 // left most pivot

	/*
	* Scan valus from right to left.
	*  - Remeber the position of right most smaller than pivot.
	*  - For each pivot, try swap with right most smaller than pivot.
	*
	 */
	for i >= 0 {

		if arr[i] == piv {

			if smt.size() == 0 {
				lm = i
				i--
				continue
			}

			lm = smt.pop()
			swapf(i, lm)
			continue
		}

		smt.push(i)
		i--
	}

	return lm
}

/*
* @Desc: Partition implementation.
* @PROC: 1. Do soft partitioning with weak_partition():
*	    - All smaller or equales to pivot are on the left.
*	    - All greater than pivot are on the right.
*	 2. Sort dangling pivots to the right with pivots2right().
 */
func partition(arr []int, i int, j int, piv int) (int, int) {

	//r_edge := weak_partition(arr, i, j, piv)     // r_edge := the right most pivot value index
	var r_edge int
	l_edge := pivots2right(arr[0:r_edge+1], piv) // l_edge := the left most pivot value index

	/*
	* a_i < pivot
	* b_i > pivot
	* 		       l_edge	      r_edge
	* (a_1,	a_2, ..., a_k, piv, ........, piv,     b_1, b_2, ..., b_m)
	 */
	return l_edge, r_edge
}

/*
* @Desc: QuickSort implementation.
* @PROC: 1. Pick random pivot.
*	 2. Orginize elements around pivot.
*	 3. Run again over smaller & largerer elements.
 */
func qsort_impl(arr []int, wg *sync.WaitGroup) {

	defer wg.Done()
	if len(arr) <= 1 {
		return
	}

	m := rand.Intn(len(arr))
	piv := arr[m]
	i, j := 0, len(arr)-1
	i, j = partition(arr, i, j, piv)

	wg.Add(2)
	go qsort_impl(arr[0:i], wg)
	go qsort_impl(arr[j:], wg)
}

/*
* @Desc: Warrper method.
 */
func (arr Arr) qsort() {

	var wg sync.WaitGroup
	wg.Add(1)
	go qsort_impl(arr[0:], &wg)
	wg.Wait()
}
