package scujwc

import (
	"testing"
)

func TestJwc_getEvaList(t *testing.T) {
	tests := []struct {
		name    string
		j       *Jwc
		wantErr bool
	}{
		{
			name:    "test",
			j:       j,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetEvaList()
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jwc.getEvaList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
