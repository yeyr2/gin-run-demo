package controller

import (
	"douSheng/class"
	"douSheng/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserFriendListResponse struct {
	class.Response
	UserList []class.FriendUser `json:"user_list"`
}

type UserListResponse struct {
	class.Response
	UserList []class.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 0, 64)
	state, _ := strconv.Atoi(c.Query("action_type"))
	if user, exist := sql.FindUser(token); exist {
		if user.Id == toUserId && state == 1 {
			c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "不能关注自己"})
			return
		}
		if err := sql.RelationAction(token, toUserId, state); err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "关注错误"})
			return
		}
		c.JSON(http.StatusOK, class.Response{StatusCode: 0})
		return
	}
	c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

func FollowList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")

	c.JSON(http.StatusOK, UserListResponse{
		Response: class.Response{
			StatusCode: 0,
		},
		UserList: sql.FindFollowUsers(userId, token),
	})
}

func FollowerList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")

	c.JSON(http.StatusOK, UserListResponse{
		Response: class.Response{
			StatusCode: 0,
		},
		UserList: sql.FindFollowerUsers(userId, token),
	})
}

// FriendList 好友 : 我关注且关注我的
func FriendList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")

	c.JSON(http.StatusOK, UserFriendListResponse{
		Response: class.Response{
			StatusCode: 0,
		},
		UserList: sql.FindFriends(userId, token),
	})
}