package types

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestQuota_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
	}{
		{
			"empty account address",
			fields{
				Address: "",
			},
			nil,
		},
		{
			"20 bytes account address",
			fields{
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Quota{
				Address: tt.fields.Address,
			}
			if got := m.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuota_Validate(t *testing.T) {
	type fields struct {
		Address        string
		AllocatedBytes sdk.Int
		ConsumedBytes  sdk.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty account address",
			fields{
				Address: "",
			},
			true,
		},
		{
			"invalid account address",
			fields{
				Address: "invalid",
			},
			true,
		},
		{
			"invalid prefix account address",
			fields{
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"10 bytes account address",
			fields{
				Address:        "sent1qypqxpq9qcrsszgslawd5s",
				AllocatedBytes: sdk.NewInt(0),
				ConsumedBytes:  sdk.NewInt(0),
			},
			false,
		},
		{
			"20 bytes account address",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(0),
				ConsumedBytes:  sdk.NewInt(0),
			},
			false,
		},
		{
			"30 bytes account address",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				AllocatedBytes: sdk.NewInt(0),
				ConsumedBytes:  sdk.NewInt(0),
			},
			false,
		},
		{
			"nil allocated",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.Int{},
			},
			true,
		},
		{
			"negative allocated",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero allocated",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(0),
				ConsumedBytes:  sdk.NewInt(0),
			},
			false,
		},
		{
			"positive allocated",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(1000),
				ConsumedBytes:  sdk.NewInt(0),
			},
			false,
		},
		{
			"nil consumed",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(1000),
				ConsumedBytes:  sdk.Int{},
			},
			true,
		},
		{
			"negative consumed",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(1000),
				ConsumedBytes:  sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero consumed",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(1000),
				ConsumedBytes:  sdk.NewInt(0),
			},
			false,
		},
		{
			"positive consumed",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(1000),
				ConsumedBytes:  sdk.NewInt(1000),
			},
			false,
		},
		{
			"allocated less than consumed",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(1000),
				ConsumedBytes:  sdk.NewInt(2000),
			},
			true,
		},
		{
			"allocated equals to consumed",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(2000),
				ConsumedBytes:  sdk.NewInt(2000),
			},
			false,
		},
		{
			"allocated greater than consumed",
			fields{
				Address:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				AllocatedBytes: sdk.NewInt(2000),
				ConsumedBytes:  sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Quota{
				Address:        tt.fields.Address,
				AllocatedBytes: tt.fields.AllocatedBytes,
				ConsumedBytes:  tt.fields.ConsumedBytes,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
