package repository

import (
	"database/sql"
	"testing"

	"github.com/dhevve/shop"
	"github.com/dhevve/shop/internal/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestItemPostgres_CreateItem(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewItemPostgres(db)

	type args struct {
		item shop.Item
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
				item: shop.Item{
					Name:    "test name",
					Price:   300,
					BrandId: 1,
				},
			},
			want: 1,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO items").
					WithArgs(args.item.Name, args.item.Price, args.item.BrandId).WillReturnRows(rows)
			},
		},
		{
			name: "Empty Fields",
			mock: func(args args, id int) {

				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO items").
					WithArgs("", 300, 1).WillReturnRows(rows)

			},
			input: args{
				item: shop.Item{
					Name:    "",
					Price:   300,
					BrandId: 1,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			item := shop.Item{
				Name:    tt.input.item.Name,
				Price:   tt.input.item.Price,
				BrandId: tt.input.item.BrandId,
			}

			got, err := r.CreateItem(item)
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

func TestItemPostgres_GetItems(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewItemPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		want    []shop.Item
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "price", "brandid"}).
					AddRow(1, "name1", 100, 1).
					AddRow(2, "name2", 100, 1).
					AddRow(3, "name3", 100, 1)

				mock.ExpectQuery("SELECT (.+) FROM items").WillReturnRows(rows)
			},
			want: []shop.Item{
				{Id: 1, Name: "name1", Price: 100, BrandId: 1},
				{Id: 2, Name: "name2", Price: 100, BrandId: 1},
				{Id: 3, Name: "name3", Price: 100, BrandId: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetAllItems()
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

func TestItemPostgres_GetByIdItem(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewItemPostgres(db)

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		mock    func()
		want    shop.Item
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "price", "brandid"}).
					AddRow(1, "name1", 100, 1)

				mock.ExpectQuery("SELECT (.+) FROM items WHERE (.+)").
					WithArgs(1).WillReturnRows(rows)
			},
			input: args{
				id: 1,
			},
			want: shop.Item{Id: 1, Name: "name1", Price: 100, BrandId: 1},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "price", "brand_id"})

				mock.ExpectQuery("SELECT (.+) FROM items WHERE (.+)").
					WithArgs(404).WillReturnRows(rows)
			},
			input: args{
				id: 404,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetById(tt.input.id)
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

func TestItemPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewItemPostgres(db)

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec("DELETE FROM items WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				id: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("DELETE FROM items WHERE (.+)").
					WithArgs(404).WillReturnError(sql.ErrNoRows)
			},
			input: args{
				id: 404,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.DeleteItem(tt.input.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
