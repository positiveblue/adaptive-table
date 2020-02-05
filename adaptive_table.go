package adaptivetable

type AdaptiveTable struct {
	values             []uint64
	initSize           int
	maxSize            int
	threshold          int
	relativePercentage bool
}

func NewAdaptiveTable(initSize int) AdaptiveTable {
	return AdaptiveTable{
		initSize:  initSize,
		maxSize:   initSize,
		threshold: initSize}
}

func NewAdaptiveTableComplete(initSize, maxSize, threshold int, relativePercentage bool) AdaptiveTable {
	return AdaptiveTable{
		initSize:           initSize,
		maxSize:            maxSize,
		threshold:          threshold,
		relativePercentage: relativePercentage,
	}
}

func (at *AdaptiveTable) Size() int {
	return len(at.values)
}

func (at *AdaptiveTable) IsEmpty() bool {
	return len(at.values) == 0
}

func (at *AdaptiveTable) Min() uint64 {
	return at.values[0]
}

func (at *AdaptiveTable) Max() uint64 {
	lastIndex := len(at.values) - 1
	return at.values[lastIndex]
}

func (at *AdaptiveTable) Pop() uint64 {
	last := at.Max()
	at.values = at.values[:len(at.values)-1]
	return last
}

func (at *AdaptiveTable) Contains(value uint64) bool {
	if len(at.values) == 0 || at.Max() < value {
		return false
	}

	for _, element := range at.values {
		if element == value {
			return true
		}
	}

	return false
}

func (at *AdaptiveTable) IsNewRecord(value uint64) bool {
	if (len(at.values) < at.maxSize || at.Max() > value) && !at.Contains(value) {
		return true
	}

	return false
}

func (at *AdaptiveTable) currentThreshold() int {
	if !at.relativePercentage {
		return at.threshold
	}

	return int(float32(len(at.values)*at.threshold) / 100.00)
}

func (at *AdaptiveTable) Insert(value uint64) int {
	if !at.IsNewRecord(value) {
		return -1
	}

	at.values = append(at.values, value)

	index := len(at.values) - 1
	done := false
	for done != true {
		if index == 0 || at.values[index-1] < at.values[index] {
			done = true
		} else if at.values[index-1] > at.values[index] {
			at.values[index-1], at.values[index] = at.values[index], at.values[index-1]
			index--
		} else {
			done = true
		}
	}

	ct := at.currentThreshold()
	if index > ct || at.Size() > at.maxSize {
		at.Pop()
	}

	return index
}

func (at *AdaptiveTable) Values() []uint64 {
	return at.values
}
