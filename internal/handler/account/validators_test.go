package account

import "testing"

func Test_validateAccountId(t *testing.T) {
	type args struct {
		accountIdStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Empty account ID",
			args:    args{accountIdStr: ""},
			wantErr: true,
		},
		{
			name:    "Invalid account ID",
			args:    args{accountIdStr: "abc"},
			wantErr: true,
		},
		{
			name:    "invalid account ID: hexadecimal form",
			args:    args{accountIdStr: "0x1"},
			wantErr: true,
		},
		{
			name:    "valid account ID precedes with 0",
			args:    args{accountIdStr: "0001"},
			wantErr: false,
		},
		{
			name:    "valid 10-base integer account ID",
			args:    args{accountIdStr: "123"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAccountId(tt.args.accountIdStr); (err != nil) != tt.wantErr {
				t.Errorf("validateAccountId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
