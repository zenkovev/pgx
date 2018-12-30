package pgconn_test

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/pgconn"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	pgConn, err := pgconn.Connect(context.Background(), os.Getenv("PGX_TEST_DATABASE"))
	require.Nil(t, err)

	pgConn.SendExec("select current_database()")
	err = pgConn.Flush()
	require.Nil(t, err)

	result := pgConn.GetResult()
	require.NotNil(t, result)

	rowFound := result.NextRow()
	assert.True(t, rowFound)
	if rowFound {
		assert.Equal(t, "pgx_test", string(result.Value(0)))
	}

	_, err = result.Close()
	assert.Nil(t, err)

	err = pgConn.Close()
	assert.Nil(t, err)
}
