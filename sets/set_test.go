package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		values []T
	}
	tests := []struct {
		name string
		args args
		want Set[T]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, New(tt.args.values...))
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args struct {
		v T
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.s.Add(tt.args.v)
		})
	}
}

func TestSet_Delete(t *testing.T) {
	type args struct {
		v T
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.s.Delete(tt.args.v)
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		from Set[T]
		f    func(T) U
	}
	tests := []struct {
		name string
		args args
		want Set[U]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, Map(tt.args.from, tt.args.f))
		})
	}
}

func TestFlatMap(t *testing.T) {
	type args struct {
		from Set[T]
		f    func(T) Set[U]
	}
	tests := []struct {
		name string
		args args
		want Set[U]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, FlatMap(tt.args.from, tt.args.f))
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		from Set[T]
		f    func(T) bool
	}
	tests := []struct {
		name string
		args args
		want Set[T]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, Filter(tt.args.from, tt.args.f))
		})
	}
}

func TestForEach(t *testing.T) {
	type args struct {
		from Set[T]
		f    func(T)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ForEach(tt.args.from, tt.args.f)
		})
	}
}

func TestAsSlice(t *testing.T) {
	type args struct {
		from Set[T]
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, AsSlice(tt.args.from))
		})
	}
}

func TestAsSliceSorted(t *testing.T) {
	type args struct {
		from Set[T]
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, AsSliceSorted(tt.args.from))
		})
	}
}

func TestAsSliceSortedFunc(t *testing.T) {
	type args struct {
		from Set[T]
		less func(a, b T) bool
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, AsSliceSortedFunc(tt.args.from, tt.args.less))
		})
	}
}
