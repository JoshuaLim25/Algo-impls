package sorting

type Solution struct {
	Data []int
}

func (s Solution) BinarySearch(data []int, needle int) (bool, int) {
	if len(data) == 0 {
		return false, 0
	}

	low, high := 0, len(data)-1
	for low <= high {
		// mid := (high + low) >> 2
		mid := (high + low) / 2
		switch {
		case data[mid] > needle:
			{
				high = mid - 1
			}
		case data[mid] < needle:
			{
				low = mid + 1
			}
		case data[mid] == needle:
			{
				return true, mid
			}
		}
	}

	return false, 0
}
