package so

import "testing"

func Test_sistemaOP(t *testing.T) {
	tests := []struct {
		name string

		sistema string
		want    string
		wantErr bool
	}{
		{name: "test 1", sistema: "windows", want: "windows", wantErr: false},
		{name: "test 2", sistema: "linux", want: "linux", wantErr: false},
		{name: "test 3", sistema: "darwin", want: "osx", wantErr: false},
		{name: "falso 1", sistema: "test", want: "", wantErr: true},
		{name: "falso 2", sistema: "macos", want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := sistemaOP(tt.sistema)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("sistemaOP() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("sistemaOP() succeeded unexpectedly")
			}

			if got != tt.want {
				t.Errorf("sistemaOP() = %v, want %v", got, tt.want)
			}
		})
	}
}
