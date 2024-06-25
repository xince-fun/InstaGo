namespace go follow

include "./../../../shared/idl/base/base.thrift"

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

service FollowService {
    FollowResponse Follow(1: FollowRequest req)
    UnfollowResponse Unfollow(1: UnfollowRequest req)
}