package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount_IsDocumentValid(t *testing.T) {
	type fields struct {
		Id             int
		DocumentNumber string
		IsValid        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "successfully validate the document number",
			fields: fields{DocumentNumber: "1234567890"},
			want:   true,
		},
		{
			name:   "fail validate the document number",
			fields: fields{DocumentNumber: "123456789"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Id:             tt.fields.Id,
				DocumentNumber: tt.fields.DocumentNumber,
				IsValid:        tt.fields.IsValid,
			}
			assert.Equalf(t, tt.want, a.IsDocumentValid(), "IsDocumentValid()")
		})
	}
}
