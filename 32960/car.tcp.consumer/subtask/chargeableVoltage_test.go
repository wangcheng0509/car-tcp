package subtask

import (
	"testing"

	"car.tcp.consumer/entity/model"
)

func Test_unmarshalChargeableVoltage(t *testing.T) {
	type args struct {
		data              []byte
		chargeableVoltage *model.ChargeableVoltage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "2items",
			args: args{
				data:              []byte{2, 1, 0, 200, 3, 252, 0, 2, 0, 1, 2, 78, 32, 78, 32, 2, 0, 200, 3, 252, 0, 2, 0, 1, 2, 117, 48, 117, 48},
				chargeableVoltage: &model.ChargeableVoltage{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := unmarshalChargeableVoltage(tt.args.data, tt.args.chargeableVoltage); (err != nil) != tt.wantErr {
				t.Errorf("unmarshalChargeableVoltage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
