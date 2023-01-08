package dao

const (
	getLeaderboard = iota
	listLeaderboard
	updateLeaderboard
)

const (
	getLeaderboardStatement    = `select`
	listLeaderboardStatement   = `select`
	updateLeaderboardStatement = `update leaderboard`
)

var mapStatement = map[int]string{
	getLeaderboard:    getLeaderboardStatement,
	listLeaderboard:   listLeaderboardStatement,
	updateLeaderboard: updateLeaderboardStatement,
}
