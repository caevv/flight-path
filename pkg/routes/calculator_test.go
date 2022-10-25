package routes

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		paths []FlightRoute
	}
	tests := []struct {
		name    string
		args    args
		want    FlightRoute
		wantErr bool
	}{
		{
			name: "single",
			args: args{
				paths: []FlightRoute{{"SFO", "EWR"}},
			},
			wantErr: false,
			want:    FlightRoute{"SFO", "EWR"},
		},
		{
			name: "two routes",
			args: args{
				paths: []FlightRoute{{"ATL", "EWR"}, {"SFO", "ATL"}},
			},
			wantErr: false,
			want:    FlightRoute{"SFO", "EWR"},
		},
		{
			name: "multiple routes",
			args: args{
				paths: []FlightRoute{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			},
			wantErr: false,
			want:    FlightRoute{"SFO", "EWR"},
		},
		{
			name: "loop",
			args: args{
				paths: []FlightRoute{{"IND", "EWR"}, {"EWR", "IND"}},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "invalid same path 1",
			args: args{
				paths: []FlightRoute{{"IND", "IND"}},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "invalid same path multiple",
			args: args{
				paths: []FlightRoute{{"IND", "IND"}, {"IND", "IND"}},
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.paths)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v, got = %v", err, tt.wantErr, got)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}

			t.Logf("sent: %q, received: %q", tt.args.paths, got)
		})
	}
}
