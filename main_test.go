package main

import (
	"testing"
)

func TestGetCellInfo(t *testing.T) {
	chi, cidx, cbidx := getCellInfo(0, 0, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 1 {
		t.Errorf("expecting 1,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 0, 1)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 1 {
		t.Errorf("expecting 1,got %d", cidx)
	}
	chi, cidx, cbidx = getCellInfo(0, 0, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 517 {
		t.Errorf("expecting 517,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 0, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 517 {
		t.Errorf("expecting 517,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 1, 0)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 1 {
		t.Errorf("expecting 1,got %d", cidx)
	}
	if cbidx != 3 {
		t.Errorf("expecting 3,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(0, 1, 1)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 1 {
		t.Errorf("expecting 1,got %d", cidx)
	}
	if cbidx != 7 {
		t.Errorf("expecting 7,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 2, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 515 {
		t.Errorf("expecting 515,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 2, 1)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 3 {
		t.Errorf("expecting 3,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 2, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 2, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 3, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 3 {
		t.Errorf("expecting 3,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(0, 3, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 3 {
		t.Errorf("expecting 3,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(0, 3, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(0, 3, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(1, 0, 0)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 1 {
		t.Errorf("expecting 1,got %d", cidx)
	}
	if cbidx != 2 {
		t.Errorf("expecting 2,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 0, 1)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 1 {
		t.Errorf("expecting 1,got %d", cidx)
	}
	if cbidx != 6 {
		t.Errorf("expecting 6,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 0, 2)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 5 {
		t.Errorf("expecting 5,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 0, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 0 {
		t.Errorf("expecting 0,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 1, 0)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 1 {
		t.Errorf("expecting 1,got %d", cidx)
	}
	if cbidx != 4 {
		t.Errorf("expecting 4,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 1, 1)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 1 {
		t.Errorf("expecting 1,got %d", cidx)
	}
	if cbidx != 8 {
		t.Errorf("expecting 8,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 1, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 1, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 5 {
		t.Errorf("expecting 5,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 2, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 0 {
		t.Errorf("expecting 0,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 2, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 2, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 2, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 3, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 3 {
		t.Errorf("expecting 3,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(1, 3, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 3 {
		t.Errorf("expecting 3,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 3, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(1, 3, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 7 {
		t.Errorf("expecting 7,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 0, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 514 {
		t.Errorf("expecting 514,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(2, 0, 1)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 2 {
		t.Errorf("expecting 2,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 0, 2)
	if chi != 10 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 0, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(2, 1, 0)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 4 {
		t.Errorf("expecting 4,got %d", cidx)
	}
	if cbidx != 2 {
		t.Errorf("expecting 2,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 1, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 1, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 1, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 2, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 4 {
		t.Errorf("expecting 4,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 2, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 2, 2)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 520 {
		t.Errorf("expecting 520,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 2, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 3, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 4 {
		t.Errorf("expecting 4,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(2, 3, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 4 {
		t.Errorf("expecting 4,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 3, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 0,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(2, 3, 3)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 7 {
		t.Errorf("expecting 7,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 0, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 2 {
		t.Errorf("expecting 2,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 0, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 2 {
		t.Errorf("expecting 2,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 0, 2)
	if chi != 10 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 0, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(3, 1, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 2 {
		t.Errorf("expecting 2,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 1, 1)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 2 {
		t.Errorf("expecting 2,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 1, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 1, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 6 {
		t.Errorf("expecting 6,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 2, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 4 {
		t.Errorf("expecting 4,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 2, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 4 {
		t.Errorf("expecting 4,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 2, 2)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 2, 3)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 8,got %d", cidx)
	}
	if cbidx != 6 {
		t.Errorf("expecting 6,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 3, 0)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 4 {
		t.Errorf("expecting 4,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}

	chi, cidx, cbidx = getCellInfo(3, 3, 1)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 4 {
		t.Errorf("expecting 4,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 3, 2)
	if chi != 11 {
		t.Errorf("expecting 11,got %d", chi)
	}
	if cidx != 8 {
		t.Errorf("expecting 0,got %d", cidx)
	}
	if cbidx != 4 {
		t.Errorf("expecting 4,got %d", cbidx)
	}
	chi, cidx, cbidx = getCellInfo(3, 3, 3)
	if chi != 10 {
		t.Errorf("expecting 10,got %d", chi)
	}
	if cidx != 520 {
		t.Errorf("expecting 520,got %d", cidx)
	}
	if cbidx != 0 {
		t.Errorf("expecting 0,got %d", cbidx)
	}
}
