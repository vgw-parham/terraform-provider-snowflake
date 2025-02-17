package snowflake

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExternalFunctionCreate(t *testing.T) {
	r := require.New(t)
	s := ExternalFunction("test_function", "test_db", "test_schema")
	s.WithArgs([]map[string]string{{"name": "data", "type": "varchar"}})
	s.WithArgTypes("varchar")
	s.WithReturnType("varchar")
	s.WithNullInputBehavior("RETURNS NULL ON NULL INPUT")
	s.WithReturnBehavior("IMMUTABLE")
	s.WithAPIIntegration("test_api_integration_01")
	s.WithURLOfProxyAndResource("https://123456.execute-api.us-west-2.amazonaws.com/prod/test_func")

	r.Equal(`"test_db"."test_schema"."test_function"`, s.QualifiedName())
	r.Equal(`"test_db"."test_schema"."test_function" (varchar)`, s.QualifiedNameWithArgTypes())

	expected := `CREATE EXTERNAL FUNCTION "test_db"."test_schema"."test_function" (data varchar) RETURNS varchar NULL RETURNS NULL ON NULL INPUT IMMUTABLE API_INTEGRATION = 'test_api_integration_01' AS 'https://123456.execute-api.us-west-2.amazonaws.com/prod/test_func'`
	r.Equal(expected, s.Create())
}

func TestExternalFunctionDrop(t *testing.T) {
	r := require.New(t)

	// Without arg
	s := ExternalFunction("test_function", "test_db", "test_schema")
	r.Equal(`DROP FUNCTION "test_db"."test_schema"."test_function" ()`, s.Drop())

	// With arg
	s = ExternalFunction("test_function", "test_db", "test_schema").WithArgTypes("varchar")
	r.Equal(`DROP FUNCTION "test_db"."test_schema"."test_function" (varchar)`, s.Drop())
}

func TestExternalFunctionShow(t *testing.T) {
	r := require.New(t)
	s := ExternalFunction("test_function", "test_db", "test_schema")
	r.Equal(`SHOW EXTERNAL FUNCTIONS LIKE 'test_function' IN SCHEMA "test_db"."test_schema"`, s.Show())
}
