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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAccountId(tt.args.accountIdStr); (err != nil) != tt.wantErr {
				t.Errorf("validateAccountId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
