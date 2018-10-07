package compound

import "testing"

func TestInterest(t *testing.T) {
	type args struct {
		p   float64
		pmt float64
		r   float64
		n   float64
		t   float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"1000 initial. 200/month. 5 years. 12% interest. compounded monthly.",
			args{1000, 200, 12, 12, 5},
			18150.630669845916,
		},
		{
			"1000 initial. 1234/month. 1.5 years. 0.87% interest. compounded quarterly.",
			args{1000, 1234, .87, 4, 1.5},
			23346.249742819935,
		},
		{
			"1000 initial. 1234/month. 15 years. 5% interest. compounded bi-annualy.",
			args{1000, 1234, 5, 2, 15},
			327153.1817999441,
		},
		{
			"1000 initial. 100/month. 10 years. 5% interest. compounded annualy.",
			args{1000, 100, 5, 1, 10},
			16722.36566943604,
		},
		{
			"1000000 initial. 10000/month. 50 years. 12% interest. compounded every other year.",
			args{1000000, 10000, 12, .5, 50},
			4.320839854547448e+08,
		},
		{
			"10000 initial. 100/month. 10 years. 5% interest. compounded monthly",
			args{10000, 100, 5, 12, 10},
			31998.322921469724,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interest(tt.args.p, tt.args.pmt, tt.args.r, tt.args.n, tt.args.t); got != tt.want {
				t.Errorf("Interest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInterest(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Interest(20000, 1000, 6.75, 12, 30)
	}
}
