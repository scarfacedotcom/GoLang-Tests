package prices

import "testing"

func TestTotalPrice(t *testing.T) {
	type parameters struct {
		nights  uint
		rate    uint
		cityTax uint
	}

	type testCase struct {
		name string
		args parameters
		want uint
	}
	tests := []testCase{
		{
			name: "test 0 nights",
			args: parameters{nights: 0, rate: 150, cityTax: 12},
			want: 12,
		},
		{
			name: "test 1 nights",
			args: parameters{nights: 1, rate: 100, cityTax: 12},
			want: 112,
		},
		{
			name: "test 2 nights",
			args: parameters{nights: 2, rate: 100, cityTax: 12},
			want: 212,
		},
		{
			name: "test 3 nights",
			args: parameters{nights: 3, rate: 100, cityTax: 12},
			want: 312,
		},
		{
			name: "test 4 nights",
			args: parameters{nights: 4, rate: 100, cityTax: 12},
			want: 412,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalPrice(tt.args.nights, tt.args.rate, tt.args.cityTax); got != tt.want {
				t.Errorf("TotalPrice()= %v,want v", tt.want)

			}
		})
	}
}
