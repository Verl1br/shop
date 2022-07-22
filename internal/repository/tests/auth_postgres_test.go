package repository

import (
	"testing"

	"github.com/dhevve/shop"
	"github.com/dhevve/shop/internal/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestAuthPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewAuthPostgres(db)

	type args struct {
		item shop.User
	}

	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				item: shop.User{
					Name:     "test name",
					Username: "test username",
					Password: "qwerty",
				},
			},
			want: 1,
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(args.item.Name, args.item.Username, args.item.Password).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO baskets").WithArgs(id).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			user := shop.User{
				Name:     tt.input.item.Name,
				Username: tt.input.item.Username,
				Password: tt.input.item.Password,
			}

			got, err := r.CreateUser(user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
