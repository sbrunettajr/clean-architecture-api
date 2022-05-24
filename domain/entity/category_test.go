package entity

import "testing"

func TestCategory_IsActive(t *testing.T) {
	testCases := []struct {
		desc   string
		active bool
	}{
		{
			desc:   "Active",
			active: true,
		},
		{
			desc:   "Inactive",
			active: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			category := Category{
				Active: tC.active,
			}
			if category.IsActive() != tC.active {
				t.Fail()
			}
		})
	}
}
