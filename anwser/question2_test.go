package anwser

import (
	"reflect"
	"testing"
	"time"
)

func TestPipe(t *testing.T) {
	type args struct {
		value interface{}
		args  []func(interface{}) interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "string and Increment once",
			args: args{"test", []func(interface{}) interface{}{Increment}},
			want: "test1",
		},
		{
			name: "int and Increment once",
			args: args{int64(1), []func(interface{}) interface{}{Increment}},
			want: int64(2),
		},
		{
			name: "float and Increment once",
			args: args{3.6, []func(interface{}) interface{}{Increment}},
			want: 4.6,
		},
		{
			name: "string and no Increment",
			args: args{"test", []func(interface{}) interface{}{}},
			want: "test",
		},
		{
			name: "int and no Increment",
			args: args{int64(0), []func(interface{}) interface{}{}},
			want: int64(0),
		},
		{
			name: "float and no Increment",
			args: args{2.4, []func(interface{}) interface{}{}},
			want: 2.4,
		},
		{
			name: "nil and Increment once",
			args: args{nil, []func(interface{}) interface{}{Increment}},
			want: nil,
		},
		{
			name: "Unexpected type and Increment once",
			args: args{time.Now(), []func(interface{}) interface{}{Increment}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pipe(tt.args.value, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pipe() = %v, want %v", got, tt.want)
			}
		})
	}
}
