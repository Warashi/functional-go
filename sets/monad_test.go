package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonadImpl_Unit(t *testing.T) {
	type args struct {
		value U
	}
	tests := []struct {
		name string
		m    MonadImpl[T, U]
		args args
		want Set[U]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.m.Unit(tt.args.value))
		})
	}
}

func TestMonadImpl_Bind(t *testing.T) {
	type args struct {
		src Set[T]
		f   func(T) Set[U]
	}
	tests := []struct {
		name string
		m    MonadImpl[T, U]
		args args
		want Set[U]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.m.Bind(tt.args.src, tt.args.f))
		})
	}
}

func TestMonadImpl_Zero(t *testing.T) {
	tests := []struct {
		name string
		m    MonadImpl[T, U]
		want Set[U]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.m.Zero())
		})
	}
}

func TestMonadImpl_Plus(t *testing.T) {
	type args struct {
		a Set[T]
		b Set[T]
	}
	tests := []struct {
		name string
		m    MonadImpl[T, U]
		args args
		want Set[T]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.m.Plus(tt.args.a, tt.args.b))
		})
	}
}
