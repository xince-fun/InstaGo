namespace go user

include "./../../shared/idl/base/base.thrift"

struct LoginRequest {
    1: string account (api.body = "account" api.vd = "len($) > 0 && len($) < 100>")
    2: string passwd (api.body = "passwd" api.vd = "len($) > 0 && len($) < 33>")
}

struct RegisterRequest {
    1: string account (api.body = "account" api.vd = "len($) > 0 && len($) < 33")
    2: string passwd (api.body = "passwd" api.vd = "len($) > 0 && len($) < 33")
    3: string fullName (api.body = "fullname" api.vd = "len($) > 0 && len($) < 33")
    4: string phoneOrEmail (api.body = "phone_or_email" api.vd = "len($) > 0 && len($) < 100")
}

struct UpdateEmailRequest {
    1: string email (api.body = "email" api.vd = "len($) > 0 && len($) < 100")
}

struct UpdatePhoneRequest {
    1: string phone (api.body = "phone" api.vd = "len($) > 0 && len($) < 33")
}

struct UpdatePasswdRequest {
    1: string oldPasswd (api.body = "old_passwd" api.vd = "len($) > 0 && len($) < 33")
    2: string newPasswd (api.body = "new_passwd" api.vd = "len($) > 0 && len($) < 33")
}

struct UpdateBirthDayRequest {
    1: i32 year (api.body = "year" api.vd = "$ > 0 && $ < 3000")
    2: i32 month (api.body = "month" api.vd = "$ > 0 && $ < 13")
    3: i32 day (api.body = "day" api.vd = "$ > 0 && $ < 32")
}

struct UploadAvatarRequest {}

struct UpdateAvatarInfoRequest {
    1: string avatar_id (api.body = "avatar_id")
    2: string object_name (api.body = "object_name")
    3: i8 blob_type (api.body = "blob_type")
}

struct GetAvatarRequest {}

service UserService {
    base.NilResponse Login(1: LoginRequest req) (api.post = "/api/v1/user/login")
    base.NilResponse Register(1: RegisterRequest req) (api.post = "/api/v1/user/register")
    base.NilResponse UpdateEmail(1: UpdateEmailRequest req) (api.put = "/api/v1/user/email")
    base.NilResponse UpdatePhone(1: UpdatePhoneRequest req) (api.put = "/api/v1/user/phone")
    base.NilResponse UpdatePasswd(1: UpdatePasswdRequest req) (api.put = "/api/v1/user/passwd")
    base.NilResponse UpdateBirthDay(1: UpdateBirthDayRequest req) (api.put = "/api/v1/user/birthday")
    base.NilResponse UploadAvatar(1: UploadAvatarRequest req) (api.post = "/api/v1/user/avatar")
    base.NilResponse UpdateAvatarInfo(1: UpdateAvatarInfoRequest req) (api.put = "/api/v1/user/avatar")

    base.NilResponse GetAvatar(1: GetAvatarRequest req) (api.get = "/api/v1/user/avatar")
}