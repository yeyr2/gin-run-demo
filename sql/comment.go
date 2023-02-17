package sql

import (
	"douSheng/class"
	"log"
)

func FindComments(videoId int, token string) []class.Comment {
	tx := getDB()
	var comments []class.Comment

	users := make(map[int64]class.User)
	myUser, _ := FindUser(token)

	result := tx.Table("comment c").Preload("Author").
		Select("c.user_token, c.id as c_id , c.content, c.create_date, c.video_id,u.name,u.follow_count,u.follower_count,u.token").
		Joins("left join user u on c.user_token = u.token").
		Where("c.video_id = ?", videoId).Order("create_date").Find(&comments)

	for i := range comments {
		if comments[i].Id == 0 {
			comments[i].Id = comments[i].CId
		}

		user, ok := users[comments[i].Id] // 自己与目标用户关系
		if !ok {
			if myUser.Id == comments[i].Author.Id {
				comments[i].Author.IsFollow = false
			} else {
				result := db.Where("my_id = ? and other_user_id = ? and state = 1", myUser.Id, comments[i].Author.Id).Find(&class.Relation{}).RowsAffected
				if result == 0 {
					comments[i].Author.IsFollow = false
				} else {
					comments[i].Author.IsFollow = true
				}
			}
			users[comments[i].Author.Id] = comments[i].Author
		} else {
			comments[i].Author.IsFollow = user.IsFollow
		}
	}

	if result.Error != nil {
		log.Println(result.Error)
	}
	return comments
}

// ReviseComment 根据actionType状态添加或者删除comment到数据库
func ReviseComment(comment class.Comment) (int64, error) {
	tx := getDB()
	var err error
	if comment.Type == 1 { // 添加
		err = tx.Create(&comment).Error
	} else { // 删除
		err = tx.Where("id = ?", comment.Id).Delete(&class.Comment{}).Error
	}

	return comment.Id, err
}