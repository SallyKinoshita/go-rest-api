package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_newName(t *testing.T) {
	t.Parallel()
	type args struct {
		first     string
		last      string
		firstKana string
		lastKana  string
	}

	successTests := []struct {
		name string
		args args
		want Name
	}{
		{
			name: "success",
			args: args{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
			want: Name{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := newName(tt.args.first, tt.args.last, tt.args.firstKana, tt.args.lastKana)
			if err != nil {
				t.Errorf("newName() error = %v, wantErr %v", err, nil)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("newName() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestName_validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		first     string
		last      string
		firstKana string
		lastKana  string
	}

	successTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Name{
				first:     tt.fields.first,
				last:      tt.fields.last,
				firstKana: tt.fields.firstKana,
				lastKana:  tt.fields.lastKana,
			}

			if err := n.validate(); err != nil {
				t.Errorf("Name.validate() error = %v, wantErr %v", err, nil)
			}
		})
	}

	failureTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "failure: first name is below the min length required",
			fields: fields{
				first:     "",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
		},
		{
			name: "failure: last name is below the min length required",
			fields: fields{
				first:     "太郎",
				last:      "",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
		},
		{
			name: "failure: first name kana is below the min length required",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "",
				lastKana:  "ヤマダ",
			},
		},
		{
			name: "failure: last name kana is below the min length required",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "",
			},
		},
		{
			name: "failure: first name exceeds the max length allowed",
			fields: fields{
				first:     "太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗太朗",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
		},
		{
			name: "failure: last name exceeds the max length allowed",
			fields: fields{
				first:     "太郎",
				last:      "山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
		},
		{
			name: "failure: first name kana exceeds the max length allowed",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウタロウタロウタロウタロウタロウタロウタロウタロウタロウタロウタロウタロウタロウタロウタロウタロウ",
				lastKana:  "ヤマダ",
			},
		},
		{
			name: "failure: last name kana exceeds the max length allowed",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダヤマダ",
			},
		},
		{
			name: "failure: first name kana contains kana characters",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "太朗",
				lastKana:  "ヤマダ",
			},
		},
		{
			name: "failure: last name kana contains kana characters",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "山田",
			},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Name{
				first:     tt.fields.first,
				last:      tt.fields.last,
				firstKana: tt.fields.firstKana,
				lastKana:  tt.fields.lastKana,
			}
			if err := n.validate(); err == nil {
				t.Errorf("Name.validate() error = %v, wantErr = is not nil", err)
			}
		})
	}
}

func TestName_Full(t *testing.T) {
	t.Parallel()
	type fields struct {
		first     string
		last      string
		firstKana string
		lastKana  string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
			want: "山田太郎",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Name{
				first:     tt.fields.first,
				last:      tt.fields.last,
				firstKana: tt.fields.firstKana,
				lastKana:  tt.fields.lastKana,
			}

			got := n.Full()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Name.Full() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestName_FullKana(t *testing.T) {
	t.Parallel()
	type fields struct {
		first     string
		last      string
		firstKana string
		lastKana  string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
			want: "ヤマダタロウ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Name{
				first:     tt.fields.first,
				last:      tt.fields.last,
				firstKana: tt.fields.firstKana,
				lastKana:  tt.fields.lastKana,
			}
			got := n.FullKana()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Name.FullKana() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestName_First(t *testing.T) {
	t.Parallel()
	type fields struct {
		first     string
		last      string
		firstKana string
		lastKana  string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
			want: "太郎",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Name{
				first:     tt.fields.first,
				last:      tt.fields.last,
				firstKana: tt.fields.firstKana,
				lastKana:  tt.fields.lastKana,
			}

			got := n.First()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Name.First() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestName_Last(t *testing.T) {
	t.Parallel()
	type fields struct {
		first     string
		last      string
		firstKana string
		lastKana  string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
			want: "山田",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Name{
				first:     tt.fields.first,
				last:      tt.fields.last,
				firstKana: tt.fields.firstKana,
				lastKana:  tt.fields.lastKana,
			}

			got := n.Last()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Name.Last() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestName_FirstKana(t *testing.T) {
	t.Parallel()
	type fields struct {
		first     string
		last      string
		firstKana string
		lastKana  string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
			want: "タロウ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Name{
				first:     tt.fields.first,
				last:      tt.fields.last,
				firstKana: tt.fields.firstKana,
				lastKana:  tt.fields.lastKana,
			}

			got := n.FirstKana()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Name.FirstKana() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestName_LastKana(t *testing.T) {
	t.Parallel()
	type fields struct {
		first     string
		last      string
		firstKana string
		lastKana  string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
			want: "ヤマダ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Name{
				first:     tt.fields.first,
				last:      tt.fields.last,
				firstKana: tt.fields.firstKana,
				lastKana:  tt.fields.lastKana,
			}

			got := n.LastKana()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Name.LastKana() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
