package config

import "fmt"

func Validate(c *Config) error {
	for _, sql := range c.SQL {
		sqlGo := sql.Gen.Go
		if sqlGo == nil {
			continue
		}
		if sqlGo.EmitMethodsWithDBArgument && sqlGo.EmitPreparedQueries {
			return fmt.Errorf("invalid config: emit_methods_with_db_argument and emit_prepared_queries settings are mutually exclusive")
		}
		if sql.Database != nil {
			if sql.Database.URL == "" {
				return fmt.Errorf("invalid config: database must have a non-empty URL")
			}
		}
	}
	return nil
}
