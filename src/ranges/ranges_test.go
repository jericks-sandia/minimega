package ranges

import "testing"
import "fmt"

func TestSplitRange(t *testing.T) {
	r, _ := NewRange("kn", 1, 520)

	expected := []string{"kn1", "kn2", "kn3", "kn100"}
	input := "kn[1-3,100]"

	res, _ := r.SplitRange(input)

	es := fmt.Sprintf("%v", expected)
	rs := fmt.Sprintf("%v", res)

	if es != rs {
		t.Fatal("SplitRange returned: ", res, ", expected: ", expected)
	}
}

func TestUnsplitRange(t *testing.T) {
	r, _ := NewRange("kn", 1, 520)

	expected := "kn[1-5]"
	input := []string{"kn1", "kn2", "kn3", "kn4", "kn5"}

	res, err := r.UnsplitRange(input)
	if err != nil {
		t.Fatal("UnsplitRange returned error: ", err)
	}
	if expected != res {
		t.Fatal("UnsplitRange returned: ", res)
	}

	expected = "kn[1-5,20]"
	input = []string{"kn1", "kn2", "kn3", "kn4", "kn5", "kn20"}

	res, err = r.UnsplitRange(input)
	if err != nil {
		t.Fatal("UnsplitRange returned error: ", err)
	}
	if expected != res {
		t.Fatal("UnsplitRange returned: ", res)
	}

	expected = "kn[1-5,20,44-45]"
	input = []string{"kn44", "kn45", "kn1", "kn2", "kn3", "kn4", "kn5", "kn20"}

	res, err = r.UnsplitRange(input)
	if err != nil {
		t.Fatal("UnsplitRange returned error: ", err)
	}
	if expected != res {
		t.Fatal("UnsplitRange returned: ", res)
	}
}
