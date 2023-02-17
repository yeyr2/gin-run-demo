package sql

import (
	"douSheng/class"
)

func MessageChat(fromUserId int, toUserId int, recentTime int64) ([]class.Message, int64) {
	var messages []class.Message
	db.
		Where("((my_id = ? and to_user_id = ?) or (my_id = ? and to_user_id = ?)) and create_at > ?",
			fromUserId, toUserId, toUserId, fromUserId, recentTime).
		Order("create_at").
		Find(&messages)

	if len(messages) != 0 {
		recentTime = messages[len(messages)-1].CreateAt
	}

	return messages, recentTime
}

func InsertMessage(message class.Message) {
	db.Create(&message)
}