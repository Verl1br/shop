package repository

import (
	"database/sql"
	"testing"

	"github.com/dhevve/shop"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestItemPostgres_CreateItem(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewItemPostgres(db)

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
					Name:  "test name",
					Price: 300,
				},
			},
			want: 1,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO items").
					WithArgs(args.item.Name, args.item.Price).WillReturnRows(rows)
			},
		},
		{
			name: "Empty Fields",
			mock: func(args args, id int) {

				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO items").
					WithArgs("", 300).WillReturnRows(rows)

			},
			input: args{
				item: shop.Item{
					Name:  "",
					Price: 300,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			item := shop.Item{
				Name:  tt.input.item.Name,
				Price: tt.input.item.Price,
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

	r := NewItemPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		want    []shop.Item
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(1, "name1", 100).
					AddRow(2, "name2", 100).
					AddRow(3, "name3", 100)

				mock.ExpectQuery("SELECT (.+) FROM items").WillReturnRows(rows)
			},
			want: []shop.Item{
				{Id: 1, Name: "name1", Price: 100},
				{Id: 2, Name: "name2", Price: 100},
				{Id: 3, Name: "name3", Price: 100},
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

	r := NewItemPostgres(db)

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
				rows := sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(1, "name1", 100)

				mock.ExpectQuery("SELECT (.+) FROM items WHERE (.+)").
					WithArgs(1).WillReturnRows(rows)
			},
			input: args{
				id: 1,
			},
			want: shop.Item{Id: 1, Name: "name1", Price: 100},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "price"})

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

	r := NewItemPostgres(db)

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
