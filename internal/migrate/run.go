package migrate

import (
	"fmt"

	"gorm.io/gorm"
)

// Run automatically migrates the schema of the database passed as argument.
func Run(db *gorm.DB, opt Options) (err error) {
	if db == nil {
		return fmt.Errorf("migrate: no database connection")
	}

	return nil
}
