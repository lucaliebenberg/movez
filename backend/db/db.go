package db

type Store struct {
	User UserStore
}

type NewSqlUserStore struct {}