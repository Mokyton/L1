package main

import "testing"

func Test_binSearch(t *testing.T) {
	type args struct {
		data []int
		val  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test_1",
			args:    args{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, val: 5},
			want:    4,
			wantErr: false,
		},
		{
			name:    "test_2",
			args:    args{data: []int{10, 15, 18, 20, 23, 25, 88, 125}, val: 10},
			want:    0,
			wantErr: false,
		},
		{
			name:    "test_3",
			args:    args{data: []int{10, 15, 18, 20, 23, 25, 88, 125}, val: 88},
			want:    6,
			wantErr: false,
		},
		{
			name:    "test_err_val_not_in_slice",
			args:    args{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, val: 12},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "test_err_unsorted_slice",
			args:    args{data: []int{5, 4, 1, 8, 9, 3, 2, 7}, val: 5},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := binSearch(tt.args.data, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("binSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("binSearch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
