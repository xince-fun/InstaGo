namespace go relation

include "./../../shared/idl/base/base.thrift"

struct FollowRequest {
    1: string followeeID (api.body = "followee_id")
}

struct UnfollowRequest {
    1: string followeeID (api.body = "followee_id")
}

struct CountFolloweesRequest {
}

struct CountFollowersRequest {
}

struct GetFolloweesRequest {
    1: i32 offset (api.query = "offset")
    2: i32 limit (api.query = "limit")
}

struct GetFollowersRequest {
    1: i32 offset (api.query = "offset")
    2: i32 limit (api.query = "limit")
}

service RelationService {
    base.NilResponse Follow(1: FollowRequest req) (api.post = "/api/v1/relation/follow")
    base.NilResponse Unfollow(1: UnfollowRequest req) (api.post = "/api/v1/relation/unfollow")
    base.NilResponse CountFolloweeList(1: CountFolloweesRequest req) (api.get = "/api/v1/relation/followee/count")
    base.NilResponse CountFollowerList(1: CountFollowersRequest req) (api.get = "/api/v1/relation/follower/count")
    base.NilResponse GetFolloweeList(1: GetFolloweesRequest req) (api.get = "/api/v1/relation/followee/list")
    base.NilResponse GetFollowerList(1: GetFollowersRequest req) (api.get = "/api/v1/relation/follower/list")
}