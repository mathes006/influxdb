package tsm1

import "testing"

// TestByteSliceSlabPool is a sanity check, verifying reference counting (and
// thus re-use).
func TestByteSliceSlabPool(t *testing.T) {
	p := NewByteSliceSlabPool()
	checkRefs := func(n int64) {
		if p.refs != n {
			t.Fatalf("bad refs %d", n)
		}
	}
	checkRefs(0)
	empty := p.Get(0)
	checkRefs(0)
	if empty != nil {
		t.Fatal("bad bufA")
	}

	bufA := p.Get(1)
	checkRefs(1)

	p.Inc(bufA)
	checkRefs(2)

	bufA[0] = 'x'

	bufB := p.Get(1)
	checkRefs(3)
	bufB[0] = 'y'

	if bufA[0] != 'x' {
		t.Fatal("bad overwrite")
	}

	p.Dec(bufA)
	checkRefs(2)
	p.Dec(bufA)
	checkRefs(1)

	bufC := p.Get(1)
	checkRefs(2)
	bufC[0] = 'z'

	if bufC[0] != 'z' {
		t.Fatal("bad reuse")
	}
}

func TestStringSlabPool(t *testing.T) {
	p := NewStringSlabPool()
	s0, b0 := p.Get(1)
	b0[0] = 'x'
	if s0 != "x" {
		t.Fatal("bad string write s0")
	}

	s1, b1 := p.Get(1)
	b1[0] = 'y'
	if s1 != "y" {
		t.Fatal("bad string write s1")
	}
	if s0 != "x" {
		t.Fatal("bad overwrite s0")
	}


}
