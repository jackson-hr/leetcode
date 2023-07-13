package easy

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	sorted := make([]int, m+n)
	for i <= m && j <= n {
		if i >= m {
			copy(sorted[i+j:], nums2[j:])
			break
		}
		if j >= n {
			copy(sorted[i+j:], nums1[i:])
			break
		}

		if nums1[i] <= nums2[j] {
			sorted[i+j] = nums1[i]
			i++
		}
		if nums2[j] <= nums1[i] {
			sorted[i+j] = nums2[j]
			j++
		}
	}
	copy(nums1, sorted)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	//nums1 = append(nums1, nums2...)
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	tail, i, j := m+n-1, m-1, n-1
	for tail >= 0 && i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[tail] = nums2[j]
			j--
		} else {
			nums1[tail] = nums1[i]
			i--
		}
		tail--
	}
	if tail >= 0 && j >= 0 {
		copy(nums1, nums2[:j+1])
	}
}
