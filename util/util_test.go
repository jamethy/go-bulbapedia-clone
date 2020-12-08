package util

import (
	"reflect"
	"testing"
)

func Test_doesMatch(t *testing.T) {
	type args struct {
		route        string
		currentRoute string
	}
	tests := []struct {
		name  string
		args  args
		want map[string]string
		want1  bool
	}{
		{
			name: "simple-value",
			args: args{
				route:        "/test/{value}",
				currentRoute: "/test/1",
			},
			want: map[string]string{
				"value": "1",
			},
			want1: true,
		},
		{
			name: "complex-values",
			args: args{
				route:        "/test/{value}-{second}",
				currentRoute: "/test/1-2",
			},
			want: map[string]string{
				"value": "1",
				"second": "2",
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rgx := RouteToRegex(tt.args.route)
			got, got1 := ParseRoute(rgx, tt.args.currentRoute)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRoute() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseRoute() got1 = %v, want1 %v", got1, tt.want1)
			}
		})
	}
}
