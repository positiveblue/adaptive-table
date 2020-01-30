package adaptivetable

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestAdaptiveTableSize(t *testing.T) {
	at := NewAdaptiveTable(8)

	if at.Size() != 0 {
		t.Error("An empty table should have size 0")
	}

	// Insert elements until the table is full
	for i := 1; i <= 8; i++ {
		at.Insert(uint64(i))
		if at.Size() != i {
			t.Error(fmt.Sprintf("Table should have size %d", i))
		}
	}

	at.Insert(1)
	if at.Size() != 8 {
		t.Error("Insert elements that are in the table should not change its size")
	}

	at.Insert(10)
	if at.Size() != 8 {
		t.Error("Insert elements that are bigger than Max() in the table should not change its size")
	}

	at.Insert(0)
	if at.Size() != 8 {
		t.Error("Insert elements that are smaller than Max() in the table should not change its size")
	}
}

func TestAdaptiveTableIsEmpty(t *testing.T) {
	at := NewAdaptiveTable(2)

	if !at.IsEmpty() {
		t.Error("AdaptiveTable are empty by default")
	}

	at.Insert(0)
	if at.IsEmpty() {
		t.Error("AdaptiveTable are not empty after one insertion")
	}

	at.Pop()
	if !at.IsEmpty() {
		t.Error("AdaptiveTable is empty after same name of Insertion() and Pop()")
	}

	at.Insert(0)
	at.Insert(1)
	at.Pop()
	if at.IsEmpty() {
		t.Error("AdaptiveTable is not empty when number of Insertion() > Pop()")
	}
}

func TestAdaptiveTableMin(t *testing.T) {
	at := NewAdaptiveTable(2)

	at.Insert(1)
	if at.Min() != 1 {
		t.Error("A table with just one element has that element as Min()")
	}

	at.Insert(2)
	if at.Min() != 1 {
		t.Error("A table with 1 and 2 has Min() == 1")
	}

	at.Insert(0)
	if at.Min() != 0 {
		t.Error("After inserting 0 to a table with one the Min() == 0")
	}
}

func TestAdaptiveTableMax(t *testing.T) {
	at := NewAdaptiveTable(2)

	at.Insert(1)
	if at.Max() != 1 {
		t.Error("A table with just one element has that element as Max()")
	}

	at.Insert(3)
	if at.Max() != 3 {
		t.Error("A table with 1 and 3 has Max() == 3")
	}

	at.Insert(2)
	if at.Max() != 2 {
		t.Error("After inserting 2 to a table with 1 and 3 the Max() == 3")
	}
}

func TestAdaptiveTablePop(t *testing.T) {
	at := NewAdaptiveTable(3)

	at.Insert(0)
	at.Insert(1)
	at.Insert(2)

	for at.Size() > 0 {
		if at.Max() != at.Pop() {
			t.Error("Pop() should always yield the biggest value in the table")
		}
	}
}

func TestAdaptiveTableContains(t *testing.T) {
	at := NewAdaptiveTable(6)

	if at.Contains(0) {
		t.Error("An empty table should not contain any element")
	}

	for i := 0; i < 11; i += 2 {
		at.Insert(uint64(i))
	}

	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			if !at.Contains(uint64(i)) {
				t.Error("The table should contain all the even elements up to 10")
			}
		} else {
			if at.Contains(uint64(i)) {
				t.Error("The table should not contain any odd element")
			}
		}
	}

	if at.Contains(11) {
		t.Error("The table should not contain any elment > 10")
	}

	max := at.Pop()

	if at.Contains(max) {
		t.Error("The table should not contain its max after Pop()")
	}
}

func TestAdaptiveTableIsNewRecord(t *testing.T) {
	at := NewAdaptiveTable(2)

	if !at.IsNewRecord(0) {
		t.Error("Any element is record for an empty table")
	}
	at.Insert(0)

	if !at.IsNewRecord(2) {
		t.Error("If the table is not full any element which is not in the table is a record")
	}
	at.Insert(2)

	if at.IsNewRecord(0) {
		t.Error("An element which is in the table is not a new record")
	}

	if !at.IsNewRecord(1) {
		t.Error("Any element smaller than Max() is a new record")
	}
	at.Insert(1)

	if at.IsNewRecord(2) {
		t.Error("When the table is full, an element is only a new record if element < Max()")
	}
}

func TestAdaptiveTableInsert(t *testing.T) {
	at := NewAdaptiveTable(3)

	var pos int

	pos = at.Insert(10)
	if pos != 0 {
		t.Error("Inserting an element to an empty table returns position 0")
	}

	pos = at.Insert(10)
	if pos != -1 {
		t.Error("Inserting an element which is already in the table returns -1")
	}

	pos = at.Insert(9)
	if pos != 0 {
		t.Error("Inserting a new top record inserts in position 0")
	}

	pos = at.Insert(11)
	if pos != 2 {
		t.Error("Inserting a new max returns the last position")
	}
}

func TestAdaptiveTableInsertStatistics(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	set := make(map[uint64]bool)

	minThreshold := 32

	at := AdaptiveTable{
		initSize:  minThreshold,
		maxSize:   math.MaxInt64,
		threshold: minThreshold}

	for i := 10000000; i > 0; i-- {
		k := rand.Uint64()
		if _, ok := set[k]; !ok {
			set[k] = true
			at.Insert(k)
		}
	}

	// Size of the table should be ~ k*log(n)
	// This is just a first approximation, further testing
	// will be included in v1.0.0
	if at.Size() < 300 || at.Size() > 600 {
		t.Error("The size of the table is not in the range, try to run this test again")
	}
}

func TestValues(t *testing.T) {

	at := NewAdaptiveTable(10)

	for i := 0; i < 10; i++ {
		at.Insert(uint64(i))
	}

	values := at.Values()
	for i := range values {
		if values[i] != uint64(i) {
			t.Error("Values() does not return the correct elements in the table")
		}
	}
}
