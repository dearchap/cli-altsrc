package yaml

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	altsrc "github.com/urfave/cli-altsrc/v3"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return altsrc.MustTestdataDir(ctx)
	}()
)

func TestYAML(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "test_config.yaml")
	altConfigPath := filepath.Join(testdataDir, "test_alt_config.yaml")

	vsc := YAML(
		"water_fountain.water",
		"/dev/null/nonexistent.yaml",
		configPath,
		altConfigPath,
	)
	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)

	yvs := vsc.Chain[0].(*yamlValueSource)
	r.Equal("yaml file \"/dev/null/nonexistent.yaml\" at key \"water_fountain.water\"", yvs.String())
	r.Equal("&yamlValueSource{file:\"/dev/null/nonexistent.yaml\",keyPath:\"water_fountain.water\"}", yvs.GoString())
}