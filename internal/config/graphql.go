package config

// GraphQLConfig toggles the GraphQL HTTP endpoint (gqlgen).
type GraphQLConfig struct {
	Enabled            bool
	Path               string
	PlaygroundEnabled  bool
}

// LoadGraphQLConfig reads GRAPHQL_* environment variables.
func LoadGraphQLConfig() GraphQLConfig {
	defaultPath := normalizeBasePath(GetEnv("API_BASE_PATH", "")) + "/graphql"
	path := GetEnv("GRAPHQL_PATH", defaultPath)
	if path == "" {
		path = defaultPath
	}
	return GraphQLConfig{
		Enabled:           parseBool(GetEnv("GRAPHQL_ENABLED", "false")),
		Path:              path,
		PlaygroundEnabled: parseBool(GetEnv("GRAPHQL_PLAYGROUND_ENABLED", "false")),
	}
}
