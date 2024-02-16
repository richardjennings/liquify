package expr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseStatement(t *testing.T) {
	stmt, err := ParseStatement(AssignStatementSelector, "a = b")
	require.NoError(t, err)
	require.Equal(t, "a", stmt.(AssignmentStmt).Variable)

	stmt, err = Parse("%assign a = b")
	require.NoError(t, err)
	require.Equal(t, "a", stmt.(AssignmentStmt).Variable)

	stmt, err = Parse(`include.background and include.background != ''`)
	require.NoError(t, err)

}
