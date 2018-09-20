package migrate

type migrateMission func() error

var configs = [...]migrateMission{
	func() error { return nil }, // 0
	missionCreateBgHistoryTable, // 1
}
