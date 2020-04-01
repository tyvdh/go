package gql

import (
	"context"
	"testing"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/stellar/go/services/ticker/internal/tickerdb"
	"github.com/stellar/go/support/db/dbtest"
	hlog "github.com/stellar/go/support/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func int32Ptr(i int32) *int32 {
	return &i
}

func TestResolver_Markets_ValidateNumHoursAgoIn28Days(t *testing.T) {
	db := dbtest.Postgres(t)
	defer db.Close()

	var session tickerdb.TickerSession
	session.DB = db.Open()
	session.Ctx = context.Background()
	defer session.DB.Close()

	// Run migrations to make sure the tests are run
	// on the most updated schema version
	migrations := &migrate.FileMigrationSource{
		Dir: "../tickerdb/migrations",
	}
	_, err := migrate.Exec(session.DB.DB, "postgres", migrations, migrate.Up)
	require.NoError(t, err)
	r := New(&session, hlog.New())

	// ✅ Less than 28 days.
	_, err = r.Markets(
		struct {
			BaseAssetCode      *string
			BaseAssetIssuer    *string
			CounterAssetCode   *string
			CounterAssetIssuer *string
			NumHoursAgo        *int32
		}{
			NumHoursAgo: int32Ptr(671),
		},
	)
	assert.NoError(t, err)

	// ✅ Exactly 28 days.
	_, err = r.Markets(
		struct {
			BaseAssetCode      *string
			BaseAssetIssuer    *string
			CounterAssetCode   *string
			CounterAssetIssuer *string
			NumHoursAgo        *int32
		}{
			NumHoursAgo: int32Ptr(672),
		},
	)
	assert.NoError(t, err)

	// ❌ More than 28 days.
	_, err = r.Markets(
		struct {
			BaseAssetCode      *string
			BaseAssetIssuer    *string
			CounterAssetCode   *string
			CounterAssetIssuer *string
			NumHoursAgo        *int32
		}{
			NumHoursAgo: int32Ptr(673),
		},
	)
	assert.EqualError(t, err, "numHoursAgo cannot be greater than 672 (28 days)")
}
