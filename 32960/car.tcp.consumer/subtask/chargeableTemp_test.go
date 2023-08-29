package subtask

import (
	"testing"

	"car.tcp.consumer/entity/model"
)

func Test_unmarshalChargeableTemp(t *testing.T) {
	type args struct {
		data []byte
		v    *model.ChargeableTemp
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "2items",
			args: args{
				data: []byte{2, 1, 0, 2, 50, 60, 2, 0, 2, 60, 65},
				v:    &model.ChargeableTemp{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := unmarshalChargeableTemp(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("unmarshalChargeableTemp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
