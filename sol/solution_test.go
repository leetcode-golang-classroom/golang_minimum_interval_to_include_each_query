package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	intervals := [][]int{
		{2, 3}, {2, 5}, {1, 8}, {20, 25},
	}
	queries := []int{
		2, 19, 5, 22,
	}
	for idx := 0; idx < b.N; idx++ {
		minInterval(intervals, queries)
	}
}
func Test_minInterval(t *testing.T) {
	type args struct {
		intervals [][]int
		queries   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]",
			args: args{intervals: [][]int{
				{1, 4}, {2, 4}, {3, 6}, {4, 4},
			}, queries: []int{
				2, 3, 4, 5,
			},
			},
			want: []int{3, 3, 1, 4},
		},
		{
			name: "intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]",
			args: args{intervals: [][]int{
				{2, 3}, {2, 5}, {1, 8}, {20, 25},
			}, queries: []int{
				2, 19, 5, 22,
			},
			},
			want: []int{2, -1, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInterval(tt.args.intervals, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
