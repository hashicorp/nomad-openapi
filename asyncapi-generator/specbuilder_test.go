package asyncapigenerator

import (
	"os"
	"path"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
)

func TestBuildSpec(t *testing.T) {
	builder := specBuilder{
		logger: hclog.Default(),
	}

	s, err := builder.buildSpec()
	require.NoError(t, err)
	require.NotNil(t, s)

	yaml, err := s.ToYAML()
	require.NoError(t, err)
	require.NotEmpty(t, yaml)

	outputPath := path.Join(t.TempDir(), "test-build.yaml")
	err = os.WriteFile(outputPath, []byte(yaml), 0644)
	require.NoError(t, err)

	outputPath = "./test-build.yaml"
	err = os.WriteFile(outputPath, []byte(yaml), 0644)
	require.NoError(t, err)
	// require.Contains(t, yaml, jobResponseSchema)

	//loader := asyncapi.Reflector{}
	//model, err := loader.LoadFromFile(outputPath)
	//require.NoError(t, err)
	//require.NotNil(t, model)
	//
	//err = model.Validate(context.Background())
	//require.NoError(t, err)
}
