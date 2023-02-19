package main

import "testing"

func Test_uniqString(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_1",
			args: args{src: "abcd"},
			want: true,
		},
		{
			name: "test_2",
			args: args{src: "abCdefAaf"},
			want: false,
		},
		{
			name: "test_3",
			args: args{src: "aabcd"},
			want: false,
		},
		{
			name: "test_4",
			args: args{src: "Acda"},
			want: false,
		},
		{
			name: "test_5",
			args: args{src: "不客气! 不用谢!"},
			want: false,
		},
		{
			name: "test_6",
			args: args{src: "客气不用谢!"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqString(tt.args.src); got != tt.want {
				t.Errorf("uniqString() = %v, want %v", got, tt.want)
			}
		})
	}
}
