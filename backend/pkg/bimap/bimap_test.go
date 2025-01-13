package bimap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Parallel()
	type args[K, V comparable] struct {
		fmap map[K]V
	}

	successTests := []struct {
		name string
		args args[string, int]
		want bimap[string, int]
	}{
		{
			name: "success",
			args: args[string, int]{
				fmap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			want: bimap[string, int]{
				forwardMap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				reverseMap: map[int]string{
					1: "a",
					2: "b",
					3: "c",
				},
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := New(tt.args.fmap)
			if err != nil {
				t.Errorf("New() error = %v, wantErr %v", err, nil)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("New() mismatch (-got +want):\n%s", diff)
			}
		})
	}

	failureTests := []struct {
		name string
		args args[string, int]
		want bimap[string, int]
	}{
		{
			name: "failure: duplicate value",
			args: args[string, int]{
				fmap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 2,
				},
			},
			want: bimap[string, int]{},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := New(tt.args.fmap)
			if err == nil {
				t.Errorf("New() error = %v, wantErr = is not nil", err)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("New() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestNewMust(t *testing.T) {
	t.Parallel()
	type args[K, V comparable] struct {
		forwardMap map[K]V
	}

	tests := []struct {
		name string
		args args[string, int]
		want bimap[string, int]
	}{
		{
			name: "success",
			args: args[string, int]{
				forwardMap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			want: bimap[string, int]{
				forwardMap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				reverseMap: map[int]string{
					1: "a",
					2: "b",
					3: "c",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewMust(tt.args.forwardMap)
			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewMust() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	t.Parallel()
	type args[K, V comparable] struct {
		forwardMap map[K]V
	}

	successTests := []struct {
		name string
		args args[string, int]
		want map[int]string
	}{
		{
			name: "success",
			args: args[string, int]{
				forwardMap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			want: map[int]string{
				1: "a",
				2: "b",
				3: "c",
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := reverse(tt.args.forwardMap)
			if err != nil {
				t.Errorf("reverse() error = %v, wantErr %v", err, nil)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("reverse() mismatch (-got +want):\n%s", diff)
			}
		})
	}

	failureTests := []struct {
		name string
		args args[string, int]
		want map[int]string
	}{
		{
			name: "failure: duplicate value",
			args: args[string, int]{
				forwardMap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 2,
				},
			},
			want: nil,
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := reverse(tt.args.forwardMap)
			if err == nil {
				t.Errorf("reverse() error = %v, wantErr = is not nil", err)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("reverse() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_bimap_GetForwardMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    *bimap[string, int]
		want map[string]int
	}{
		{
			name: "success",
			m: &bimap[string, int]{
				forwardMap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				reverseMap: map[int]string{
					1: "a",
					2: "b",
					3: "c",
				},
			},
			want: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.m.GetForwardMap()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetForwardMap() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_bimap_GetReverseMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    *bimap[string, int]
		want map[int]string
	}{
		{
			name: "success",
			m: &bimap[string, int]{
				forwardMap: map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				reverseMap: map[int]string{
					1: "a",
					2: "b",
					3: "c",
				},
			},
			want: map[int]string{
				1: "a",
				2: "b",
				3: "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.m.GetReverseMap()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetReverseMap() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
