package scujwc

import (
	"net/url"
	"reflect"
	"testing"
)

func TestJwc_GetCourse(t *testing.T) {
	type args struct {
		params url.Values
	}
	p := url.Values{}
	p.Set("pageNumber", "2")
	tests := []struct {
		name    string
		j       *Jwc
		args    args
		want    []Course
		wantErr bool
	}{
		{
			name: "test",
			j:    j,
			args: args{
				params: p,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.j
			got, err := j.GetCourse(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jwc.getCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Jwc.getCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}
