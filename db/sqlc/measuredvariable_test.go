package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMeasuredVariable(t *testing.T) {
	arg := CreateMeasuredVariableParams{
		Variable: "transmiter",
		Alias:    "t",
	}
	measuredvariable, err := testQueries.CreateMeasuredVariable(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, measuredvariable)

	require.Equal(t, arg.Variable, measuredvariable.Variable)
	require.Equal(t, arg.Alias, measuredvariable.Alias)

	require.NotZero(t, measuredvariable.ID)
	require.NotZero(t, measuredvariable.CreatedAt)
}
