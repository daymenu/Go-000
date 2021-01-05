package sliding

import (
	"testing"
	"time"
)

func TestSliding(t *testing.T) {
	data := [][]int64{
		{5},
		{6},
		{6, 7},
		{6, 7, 8},
	}
	var c *Counter
	// c = NewCounter()
	// sum := c.Sum(time.Now())
	// if sum != 0 {
	// 	t.Errorf("sum is except 0 but get %d", sum)
	// }

	for _, reqs := range data {
		c = NewCounter()
		tt := time.Now()
		var want int64
		for _, r := range reqs {
			want += r
			c.Inc(r)
			tt = time.Now()
		}
		var sum int64
		if sum = c.Sum(tt); sum != want {
			t.Errorf("sum is except %d but get %d", want, sum)
			continue
		}
		t.Logf("normal sum is except %d but get %d", want, sum)
	}

}
