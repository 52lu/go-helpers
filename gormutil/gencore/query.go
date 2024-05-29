package gencore

import "gorm.io/gen"

type QueryMethodInterface interface {
	// SELECT * FROM @@table WHERE id = @id
	FindOneById(id int64) (gen.T, error)
	// SELECT * FROM @@table order by id desc
	FindLastOne() (gen.T, error)
	// SELECT * FROM @@table order by id asc
	FindFirstOne() (gen.T, error)
}
