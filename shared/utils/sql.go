package utils

func SqlArray(num int, cb ...func(int)) string {
	sql := "("
	for i := 0; i < num; i++ {
		if len(cb) > 0 {
			cb[0](i)
		}
		if i > 0 {
			sql += ","
		}
		sql += "?"
	}
	sql += ")"
	return sql
}

func SqlValues(rows int, columns int, cb ...func(int)) string {
	sql := "VALUES "
	for i := 0; i < rows; i++ {
		if len(cb) > 0 {
			cb[0](i)
		}
		if i > 0 {
			sql += ","
		}
		sql += "("
		for j := 0; j < columns; j++ {
			if j > 0 {
				sql += ","
			}
			sql += "?"
		}
		sql += ")"
	}
	return sql
}
