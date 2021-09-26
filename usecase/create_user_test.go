package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ofiliobi/urban-octo-fortnight/domain/entity"
	"github.com/ofiliobi/urban-octo-fortnight/domain/vo"
)

type stubUserRepoCreator struct {
	result entity.User
	err    error
}

func (c stubUserRepoCreator) Create(context.Context, entity.User) (entity.User, error) {
	return c.result, c.err
}

type stubCreateUserPresenter struct {
	result CreateUserOutput
}

func (c stubCreateUserPresenter) Output(_ entity.User) CreateUserOutput {
	return c.result
}

func TestCreateUserInteractor_Execute(t *testing.T) {
	type fields struct {
		repo entity.UserRepositoryCreator
		pre  CreateUserPresenter
	}

	type args struct {
		input CreateUserInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    CreateUserOutput
		wantErr bool
	}{
		{
			name: "Create common user success",
			fields: fields{
				repo: stubUserRepoCreator{
					result: entity.NewCommonUser(
						vo.NewUuidStaticTest(),
						vo.NewFullName("Test testing"),
						vo.NewEmailTest("test@testing.com"),
						vo.NewPassword("passw"),
						vo.NewDocumentTest(vo.CNPJ, "98.521.079/0001-09"),
						nil,
						time.Time{},
					),
					err: nil,
				},
				pre: stubCreateUserPresenter{
					result: CreateUserOutput{
						ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
						FullName: "Test testing",
						Document: CreateUserDocumentOutput{
							Type:  vo.CNPJ.String(),
							Value: "34018708000191",
						},
						Email:     "test@testing.com",
						Password:  "passw",
						Wallet:    CreateUserWalletOutput{},
						Type:      vo.COMMON.String(),
						CreatedAt: time.Time{}.String(),
					},
				},
			},
			args: args{
				input: CreateUserInput{
					FullName: vo.NewFullName("Test testing"),
					Document: vo.NewDocumentTest(vo.CNPJ, "98.521.079/0001-09"),
					Email:    vo.NewEmailTest("test@testing.com"),
					Password: vo.NewPassword("passw"),
					Wallet:   nil,
					Type:     "COMMON",
				},
			},
			want: CreateUserOutput{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				FullName: "Test testing",
				Document: CreateUserDocumentOutput{
					Type:  vo.CNPJ.String(),
					Value: "34018708000191",
				},
				Email:     "test@testing.com",
				Password:  "passw",
				Wallet:    CreateUserWalletOutput{},
				Type:      vo.COMMON.String(),
				CreatedAt: time.Time{}.String(),
			},
			wantErr: false,
		},
		{
			name: "Create merchant user success",
			fields: fields{
				repo: stubUserRepoCreator{
					result: entity.NewMerchantUser(
						vo.NewUuidStaticTest(),
						vo.NewFullName("Test testing"),
						vo.NewEmailTest("test@testing.com"),
						vo.NewPassword("passw"),
						vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
						vo.NewWallet(vo.NewMoneyNGN(vo.NewAmountTest(100))),
						time.Now(),
					),
					err: nil,
				},
				pre: stubCreateUserPresenter{
					result: CreateUserOutput{
						ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
						FullName: "Test testing",
						Document: CreateUserDocumentOutput{
							Type:  vo.CNPJ.String(),
							Value: "34018708000191",
						},
						Email:     "test@testing.com",
						Password:  "passw",
						Wallet:    CreateUserWalletOutput{},
						Type:      vo.COMMON.String(),
						CreatedAt: time.Time{}.String(),
					},
				},
			},
			args: args{
				input: CreateUserInput{
					FullName: vo.NewFullName("Test testing"),
					Document: vo.NewDocumentTest(vo.CNPJ, "98.521.079/0001-09"),
					Email:    vo.NewEmailTest("test@testing.com"),
					Password: vo.NewPassword("passw"),
					Wallet:   nil,
					Type:     "COMMON",
				},
			},
			want: CreateUserOutput{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				FullName: "Test testing",
				Document: CreateUserDocumentOutput{
					Type:  vo.CNPJ.String(),
					Value: "34018708000191",
				},
				Email:     "test@testing.com",
				Password:  "passw",
				Wallet:    CreateUserWalletOutput{},
				Type:      vo.COMMON.String(),
				CreatedAt: time.Time{}.String(),
			},
			wantErr: false,
		},
		{
			name: "Create common user error",
			fields: fields{
				repo: stubUserRepoCreator{
					result: entity.User{},
					err:    errors.New("failed created user"),
				},
				pre: stubCreateUserPresenter{
					result: CreateUserOutput{},
				},
			},
			args: args{
				input: CreateUserInput{
					FullName: vo.NewFullName("Test testing"),
					Document: vo.NewDocumentTest(vo.CNPJ, "98.521.079/0001-09"),
					Email:    vo.NewEmailTest("test@testing.com"),
					Password: vo.NewPassword("passw"),
					Wallet:   nil,
					Type:     "COMMON",
				},
			},
			want:    CreateUserOutput{},
			wantErr: true,
		},
		{
			name: "Create common user database error",
			fields: fields{
				repo: stubUserRepoCreator{
					result: entity.NewCommonUser(
						vo.NewUuidStaticTest(),
						vo.NewFullName("Test testing"),
						vo.NewEmailTest("test@testing.com"),
						vo.NewPassword("passw"),
						vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
						nil,
						time.Now(),
					),
					err: errors.New("fail database"),
				},
				pre: stubCreateUserPresenter{
					result: CreateUserOutput{},
				},
			},
			args: args{
				input: CreateUserInput{
					FullName: vo.NewFullName("Test testing"),
					Document: vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
					Email:    vo.NewEmailTest("test@testing.com"),
					Password: vo.NewPassword("passw"),
					Wallet:   nil,
					Type:     vo.COMMON,
				},
			},
			want:    CreateUserOutput{},
			wantErr: true,
		},
		{
			name: "Create common user type user error",
			fields: fields{
				repo: stubUserRepoCreator{
					result: entity.NewCommonUser(
						vo.NewUuidStaticTest(),
						vo.NewFullName("Test testing"),
						vo.NewEmailTest("test@testing.com"),
						vo.NewPassword("passw"),
						vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
						nil,
						time.Now(),
					),
					err: errors.New("fail database"),
				},
				pre: stubCreateUserPresenter{
					result: CreateUserOutput{},
				},
			},
			args: args{
				input: CreateUserInput{
					FullName: vo.NewFullName("Test testing"),
					Document: vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
					Email:    vo.NewEmailTest("test@testing.com"),
					Password: vo.NewPassword("passw"),
					Wallet:   nil,
					Type:     "Test",
				},
			},
			want:    CreateUserOutput{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCreateUserInteractor(
				tt.fields.repo,
				tt.fields.pre,
			)

			got, err := c.Execute(context.Background(), tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}