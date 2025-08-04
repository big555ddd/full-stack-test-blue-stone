package migrations

func Models() []any {
	return []any{
		// (*model.User)(nil),
		// (*model.UserOtp)(nil),
		// (*model.ActivityLog)(nil),
	}
}

func RawBeforeQueryMigrate() []string {
	return []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	}
}

func RawAfterQueryMigrate() []string {
	return []string{}
}
