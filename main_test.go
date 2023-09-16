package main

import (
	"testing"
)

// 	input5 := []string{"abcd", "bnrt", "crmz", "dtye"}
// 	expected5 := false
// 	actual5, err5 := checkIfStringsFormSameSet(input5)
// 	if err5 != nil {
// 		t.Errorf("unexpected error: %v", err5)
// 	}
// 	if actual5 != expected5 {
// 		t.Errorf("expected %v but got %v", expected5, actual5)
// 	}
// }

func Test_checkIfStringsFormSameSet(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{array: []string{"abcd", "bnrt", "crmy", "dtye"}},
			want: true,
		},
		{
			name:    "invalid input (non-square array)",
			args:    args{array: []string{"abcd", "bnrt", "crmy", "dty"}},
			wantErr: true,
		},
		{
			name:    "invalid input (non-lowercase letter)",
			args:    args{array: []string{"abcd", "bnrt", "crmy", "DtYe"}},
			wantErr: true,
		},
		{
			name:    "invalid input (empty array)",
			args:    args{array: []string{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkIfStringsFormSameSet(tt.args.array)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkIfStringsFormSameSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkIfStringsFormSameSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
