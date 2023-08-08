package main

//hello admin

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	gen "car/presenters/graph_admin/generator/plugins"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var jsonTagRegexp = regexp.MustCompile(`json:".*?"`)
var jsonTagGroupRegexp = regexp.MustCompile(`json:"(.*?)"`)

func snakeCaseMutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			jsonTagGrouped := jsonTagGroupRegexp.FindStringSubmatch(field.Tag)
			snakeCase := ToSnakeCase(jsonTagGrouped[1])

			field.Tag = jsonTagRegexp.ReplaceAllString(field.Tag, fmt.Sprintf(`json:"%s"`, snakeCase))
		}
	}

	return b
}
func main() {
	path, err := findCfg("gqlgen_admin.yml")
	cfg, err := config.LoadConfig(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	cfg.Directives["hoo"] = config.DirectiveConfig{
		SkipRuntime: true,
	}
	p := modelgen.Plugin{
		MutateHook: snakeCaseMutateHook,
	}
	err = api.Generate(
		cfg,
		api.ReplacePlugin(&p),
		api.ReplacePlugin(gen.New()), // This is the magic line
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}

func findCfg(cfgName string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("unable to get working dir to findCfg: %w", err)
	}

	cfg := findCfgInDir(dir, cfgName)

	for cfg == "" && dir != filepath.Dir(dir) {
		dir = filepath.Dir(dir)
		cfg = findCfgInDir(dir, cfgName)
	}

	if cfg == "" {
		return "", os.ErrNotExist
	}

	return cfg, nil
}

func findCfgInDir(dir, cfgName string) string {
	path := filepath.Join(dir, cfgName)
	if _, err := os.Stat(path); err == nil {
		return path
	}
	return ""
}
