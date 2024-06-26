// Code generated by hertz generator. DO NOT EDIT.

package relation

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	relation "github.com/xince-fun/InstaGo/server/api/biz/handler/relation"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			{
				_relation := _v1.Group("/relation", _relationMw()...)
				_relation.POST("/follow", append(_followMw(), relation.Follow)...)
				_relation.GET("/is_follow", append(_isfollowMw(), relation.IsFollow)...)
				_relation.POST("/unfollow", append(_unfollowMw(), relation.Unfollow)...)
				{
					_followee := _relation.Group("/followee", _followeeMw()...)
					_followee.GET("/count", append(_countfolloweelistMw(), relation.CountFolloweeList)...)
					_followee.GET("/list", append(_getfolloweelistMw(), relation.GetFolloweeList)...)
				}
				{
					_follower := _relation.Group("/follower", _followerMw()...)
					_follower.GET("/count", append(_countfollowerlistMw(), relation.CountFollowerList)...)
					_follower.GET("/list", append(_getfollowerlistMw(), relation.GetFollowerList)...)
				}
			}
		}
	}
}
