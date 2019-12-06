package main

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"One to one multiply",
			args{
				a: "1",
				b: "2",
			},
			"2",
			false,
		},
		{
			"Wrong parameters",
			args{
				a: "A",
				b: "2",
			},
			"",
			true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := multiply(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("multiply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"1+1",
			args{
				"1",
				"1",
			},
			"2",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sum(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sum() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alignNumbers(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			"111, 1",
			args{
				x: "1",
				y: "111",
			},
			"001",
			"111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := alignNumbers(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("alignNumbers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("alignNumbers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_alignNumbers1(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := alignNumbers(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("alignNumbers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("alignNumbers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
