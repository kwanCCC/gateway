package protocol

import "testing"

func TestRegister(t *testing.T) {
	type args struct {
		name     string
		protocol Protocol
	}
	tcpArgs := args{"TCP", nil}
	tcpArgsII := args{"TCP", nil}
	tests := []struct {
		name       string
		args       args
		wantStatus bool
	}{
		{"TCP", tcpArgs, true,},
		{"TCP_II", tcpArgsII, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStatus := Register(tt.args.name, tt.args.protocol); gotStatus != tt.wantStatus {
				t.Errorf("Register() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestBoot(t *testing.T) {
	type args struct {
		name string
	}
	tcp := args{"TCP"}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Boot Tcp", tcp, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Boot(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Boot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
