package vo

import (
	"reflect"
	"testing"
)

func TestNewDocument(t *testing.T) {
	type args struct {
		typeDoc TypeDocument
		value   string
	}
	tests := []struct {
		name    string
		args    args
		want    Document
		wantErr bool
	}{
		{
			name: "Test new valid document",
			args: args{
				typeDoc: CPF,
				value:   "07091054954",
			},
			want: Document{
				typeDoc: CPF,
				value:   "07091054954",
			},
			wantErr: false,
		},
		{
			name: "Test new valid document",
			args: args{
				typeDoc: CPF,
				value:   "070.910.549-54",
			},
			want: Document{
				typeDoc: CPF,
				value:   "070.910.549-54",
			},
			wantErr: false,
		},
		{
			name: "Test new valid document",
			args: args{
				typeDoc: CNPJ,
				value:   "20.770.438/0001-66",
			},
			want: Document{
				typeDoc: CNPJ,
				value:   "20.770.438/0001-66",
			},
			wantErr: false,
		},
		{
			name: "Test new valid document",
			args: args{
				typeDoc: CNPJ,
				value:   "20770438000166",
			},
			want: Document{
				typeDoc: CNPJ,
				value:   "20770438000166",
			},
			wantErr: false,
		},
		{
			name: "Test new invalid document",
			args: args{
				typeDoc: "FAKER",
				value:   "2077043800016655",
			},
			want:    Document{},
			wantErr: true,
		},
		{
			name: "Test new invalid document",
			args: args{
				typeDoc: "FAK",
				value:   "954554",
			},
			want:    Document{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDocument(tt.args.typeDoc, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestDocument_Equals(t *testing.T) {
	type fields struct {
		typeDoc TypeDocument
		value   string
	}
	type args struct {
		value Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test document value equals",
			fields: fields{
				typeDoc: CPF,
				value:   "07091054954",
			},
			args: args{
				value: Document{
					typeDoc: CPF,
					value:   "07091054954",
				},
			},
			want: true,
		},
		{
			name: "Test document value equals",
			fields: fields{
				typeDoc: CNPJ,
				value:   "20770.438/0001-66",
			},
			args: args{
				value: Document{
					typeDoc: CNPJ,
					value:   "20770.438/0001-66",
				},
			},
			want: true,
		},
		{
			name: "Test document value not equals",
			fields: fields{
				typeDoc: CNPJ,
				value:   "20.770.438/0001-66",
			},
			args: args{
				value: Document{
					typeDoc: CPF,
					value:   "07091054954",
				},
			},
			want: false,
		},
		{
			name: "Test document value not equals",
			fields: fields{
				typeDoc: CPF,
				value:   "07091054954",
			},
			args: args{
				value: Document{
					typeDoc: CPF,
					value:   "070.910.549-54",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := NewDocument(tt.fields.typeDoc, tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := d.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}
