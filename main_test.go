package main

import (
	"sort"
	"testing"
)

type Packages []PackageInfo

type ByPackageSize struct {
	Packages
}

func (p ByPackageSize) Len() int {
	return len(p.Packages)
}

func (p ByPackageSize) Swap(i, j int) {
	p.Packages[i], p.Packages[j] = p.Packages[j], p.Packages[i]
}

func (p ByPackageSize) Less(i, j int) bool {
	return p.Packages[i].PackageSize > p.Packages[j].PackageSize
}

func Test_calculate_defaultPackageSize(t *testing.T) {
	items1 := []int{250, 500, 1000, 2000, 5000}

	testCases := []struct {
		name         string
		itemsOrdered int
		want         []PackageInfo
	}{
		{
			name:         "test1",
			itemsOrdered: 1,
			want: []PackageInfo{
				{
					PackageSize: 250,
					Quantity:    1,
				},
			},
		},
		{
			name:         "test2",
			itemsOrdered: 251,
			want: []PackageInfo{
				{
					PackageSize: 500,
					Quantity:    1,
				},
			},
		},
		{
			name:         "test3",
			itemsOrdered: 501,
			want: []PackageInfo{
				{
					PackageSize: 500,
					Quantity:    1,
				},
				{
					PackageSize: 250,
					Quantity:    1,
				},
			},
		},
		{
			name:         "test4",
			itemsOrdered: 12001,
			want: []PackageInfo{
				{
					PackageSize: 5000,
					Quantity:    2,
				},
				{
					PackageSize: 2000,
					Quantity:    1,
				},
				{
					PackageSize: 250,
					Quantity:    1,
				},
			},
		},
		{
			name:         "test5",
			itemsOrdered: 700,
			want: []PackageInfo{
				{
					PackageSize: 500,
					Quantity:    1,
				},
				{
					PackageSize: 250,
					Quantity:    1,
				},
			},
		},
		{
			name:         "test6",
			itemsOrdered: 999,
			want: []PackageInfo{
				{
					PackageSize: 1000,
					Quantity:    1,
				},
			},
		},
		{
			name:         "test7",
			itemsOrdered: 499,
			want: []PackageInfo{
				{
					PackageSize: 500,
					Quantity:    1,
				},
			},
		},
		{
			name:         "test8",
			itemsOrdered: 1499,
			want: []PackageInfo{
				{
					PackageSize: 1000,
					Quantity:    1,
				},
				{
					PackageSize: 500,
					Quantity:    1,
				},
			},
		},
		{
			name:         "test8",
			itemsOrdered: 800,
			want: []PackageInfo{
				{
					PackageSize: 1000,
					Quantity:    1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Calculate(tc.itemsOrdered, items1)
			sort.Sort(ByPackageSize{got})
			if len(got) != len(tc.want) {
				t.Errorf("Calculate() = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i].PackageSize != tc.want[i].PackageSize {
					t.Errorf("Calculate() = %v, want %v", got, tc.want)
				}
				if got[i].Quantity != tc.want[i].Quantity {
					t.Errorf("Calculate() = %v, want %v", got, tc.want)
				}
			}

		})
	}
}

func Test_calculate_customPackages(t *testing.T) {
	items2 := []int{23, 31, 53}
	testCases := []struct {
		name         string
		itemsOrdered int
		want         []PackageInfo
	}{
		{
			name:         "test1",
			itemsOrdered: 263,
			want: []PackageInfo{
				{
					PackageSize: 53,
					Quantity:    4,
				},
				{
					PackageSize: 31,
					Quantity:    1,
				},
				{
					PackageSize: 23,
					Quantity:    1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Calculate(tc.itemsOrdered, items2)
			sort.Sort(ByPackageSize{got})
			if len(got) != len(tc.want) {
				t.Errorf("Calculate() = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i].PackageSize != tc.want[i].PackageSize {
					t.Errorf("Calculate() = %v, want %v", got, tc.want)
				}
				if got[i].Quantity != tc.want[i].Quantity {
					t.Errorf("Calculate() = %v, want %v", got, tc.want)
				}
			}

		})
	}
}

func Test_calculate_customPackages1(t *testing.T) {
	items1 := []int{250, 500, 1000}

	testCases := []struct {
		name         string
		itemsOrdered int
		want         []PackageInfo
	}{
		{
			name:         "test1",
			itemsOrdered: 990,
			want: []PackageInfo{
				{
					PackageSize: 1000,
					Quantity:    1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Calculate(tc.itemsOrdered, items1)
			sort.Sort(ByPackageSize{got})
			if len(got) != len(tc.want) {
				t.Errorf("Calculate() = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i].PackageSize != tc.want[i].PackageSize {
					t.Errorf("Calculate() = %v, want %v", got, tc.want)
				}
				if got[i].Quantity != tc.want[i].Quantity {
					t.Errorf("Calculate() = %v, want %v", got, tc.want)
				}
			}

		})
	}
}

func Test_calculate_customPackages2(t *testing.T) {
	itemsPackage := []int{490, 500, 1000, 2000, 5000}

	testCases := []struct {
		name         string
		itemsOrdered int
		want         []PackageInfo
	}{
		{
			name:         "test1",
			itemsOrdered: 991,
			want: []PackageInfo{
				{
					PackageSize: 1000,
					Quantity:    1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Calculate(tc.itemsOrdered, itemsPackage)
			sort.Sort(ByPackageSize{got})
			if len(got) != len(tc.want) {
				t.Errorf("Calculate() = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i].PackageSize != tc.want[i].PackageSize {
					t.Errorf("Calculate() = %v, want %v", got, tc.want)
				}
				if got[i].Quantity != tc.want[i].Quantity {
					t.Errorf("Calculate() = %v, want %v", got, tc.want)
				}
			}

		})
	}
}
