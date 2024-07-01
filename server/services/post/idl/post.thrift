namespace go post

include "./../../../shared/idl/base/base.thrift"

struct PostPhotoRequest {
    1: string user_id,
    2: string title,
    3: binary photo,
}

struct PostPhotoResponse {
    1: base.BaseResponse base_resp,
    2: string photo_url,
    3: string post_id,
    4: string object_name,
}

service PostService {
    PostPhotoResponse PostPhoto(1: PostPhotoRequest req)
}