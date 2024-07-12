package status

type IStatusRepository interface {
	GetStatusDatabase() StatusDatabase
}
