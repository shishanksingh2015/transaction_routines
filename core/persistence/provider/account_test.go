package provider

import (
	"context"
	"reflect"
	"routines/core/domain"
	"routines/db"
	"testing"
)

func Test_accountProvider_GetAccountByDocumentNumber(t *testing.T) {

}

func Test_accountProvider_GetAccountById(t *testing.T) {
	type fields struct {
		dB db.BaseDB
	}
	type args struct {
		ctx  context.Context
		data int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &accountProvider{
				dB: tt.fields.dB,
			}
			got, err := a.GetAccountById(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
