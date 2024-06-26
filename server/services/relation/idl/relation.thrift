namespace go relation

include "./../../../shared/idl/base/base.thrift"
include "./../../../shared/idl/common/user.thrift"

struct FollowRequest {
    1: string follower_id
    2: string followee_id
}

struct FollowResponse {
    1: base.BaseResponse base_resp,
}

struct UnfollowRequest {
    1: string follower_id
    2: string followee_id
}

struct UnfollowResponse {
    1: base.BaseResponse base_resp,
}

struct CountFolloweeListRequest {
    1: string user_id
}

struct CountFolloweeListResponse {
    1: base.BaseResponse base_resp,
    2: i32 count
}

struct CountFollowerListRequest {
    1: string user_id
}

struct CountFollowerListResponse {
    1: base.BaseResponse base_resp,
    2: i32 count
}

struct GetFolloweeListRequest {
    1: string user_id,
    2: i32 offset,
    3: i32 limit,
}

struct GetFolloweeListResponse {
    1: base.BaseResponse base_resp,
    2: list<user.UserInfo> followee_list
}

struct GetFollowerListRequest {
    1: string user_id,
    2: i32 offset,
    3: i32 limit,
}

struct GetFollowerListResponse {
    1: base.BaseResponse base_resp,
    2: list<user.UserInfo> follower_list
}

service RelationService {
    FollowResponse Follow(1: FollowRequest req)
    UnfollowResponse Unfollow(1: UnfollowRequest req)
    CountFolloweeListResponse CountFolloweeList(1: CountFolloweeListRequest req)
    CountFollowerListResponse CountFollowerList(1: CountFollowerListRequest req)
    GetFolloweeListResponse GetFolloweeList(1: GetFolloweeListRequest req)
    GetFollowerListResponse GetFollowerList(1: GetFollowerListRequest req)
}