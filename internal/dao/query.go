package dao

import "errors"

const (
	getPlayerRank = iota
	getPlayersRanks
	listPlayerRank
	updatePlayerRank
	auth
	getPlayerScore
	addPlayerScore
	resetPlayerScore
	resetAllPlayerScore
)

const (
	getPlayerRankStatement       = "select player_id, player_rank, update_time from `rank` where player_id = ?"
	getPlayersRanksStatement     = "select player_id, player_rank, update_time from `rank` where player_id in (?)"
	listPlayerRankStatement      = "select player_id, player_rank, update_time from `rank`"
	updatePlayerRankStatement    = "update `rank` set player_rank = ?, update_time = ? where player_id = ?"
	authStatement                = "select username from `account` where username = ? and password = ?"
	getPlayerScoreStatement      = "select player_id, player_score, update_time from `score` where player_id = ?"
	addPlayerScoreStatement      = "update `score` set player_score = ?, update_time = ? where player_id = ?"
	resetPlayerScoreStatement    = "update `score` set player_score = 0, update_time = ? where player_id = ?"
	resetAllPlayerScoreStatement = "update `score` set player_score = 0, update_time = ?"
)

var mapStatement = map[int]string{
	getPlayerRank:       getPlayerRankStatement,
	getPlayersRanks:     getPlayersRanksStatement,
	listPlayerRank:      listPlayerRankStatement,
	updatePlayerRank:    updatePlayerRankStatement,
	auth:                authStatement,
	getPlayerScore:      getPlayerScoreStatement,
	addPlayerScore:      addPlayerScoreStatement,
	resetPlayerScore:    resetPlayerScoreStatement,
	resetAllPlayerScore: resetAllPlayerScoreStatement,
}

func getQueryStatement(key int) (string, error) {
	value, ok := mapStatement[key]
	if ok {
		return value, nil
	}

	return "", errors.New("failed to query statement")
}
