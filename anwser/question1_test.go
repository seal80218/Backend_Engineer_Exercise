package anwser

import "testing"

func TestGetAmountFormat(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 2.35",
			args: args{str: "2.35"},
			want: "2.35",
		},
		{
			name: "Test 2222.35",
			args: args{str: "2222.35"},
			want: "2,222.35",
		},
		{
			name: "Test 222222.35",
			args: args{str: "222222.35"},
			want: "222,222.35",
		},
		{
			name: "Test not num",
			args: args{str: "asdad"},
			want: "asdad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAmountFormat(tt.args.str); got != tt.want {
				t.Errorf("GetAmountFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
