package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateProcessVariable(t *testing.T){
	arg := CreateProcessVariableParams{
		Name: "flow",
		Alias: "f",
	}
	processvariable, err := testQueries.CreateProcessVariable(context.Background(),arg)
	require.NoError(t, err)
	require.NotEmpty(t, processvariable)

	require.Equal(t, arg.Name, processvariable.Name)
	require.Equal(t, arg.Alias, processvariable.Alias)
	require.Equal(t, arg.Description, processvariable.Description)

	require.NotZero(t, processvariable.ID)
	require.NotZero(t,processvariable.CreatedAt)
}