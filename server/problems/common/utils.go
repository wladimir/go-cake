package common

import (
	"errors"
)

func Max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("empty array")
	}

	max := nums[0]

	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	return max, nil
}
