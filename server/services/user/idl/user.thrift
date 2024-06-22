namespace go user

include "./../../../shared/idl/base/base.thrift"

struct RegisterPhoneRequest {
    1: string phoneNumber,
    2: string fullName,
    3: string account,
    4: string passwd,
}

struct RegisterEmailRequest {
    1: string email,
    2: string fullName,
    3: string account,
    4: string passwd,
}

struct RegisterResponse {
    1: base.BaseResponse base_resp,
    2: string token,
}

struct LoginPhoneRequest {
    1: string account,
    2: string passwd,
}

struct LoginEmailRequest {
    1: string account,
    2: string passwd,
}

struct LoginResponse {
    1: base.BaseResponse base_resp,
    2: string token,
}

struct UpdateEmailRequest {
    1: string user_id,
    2: string email,
}

struct UpdateEmailResponse {
    1: base.BaseResponse base_resp,
}

struct UpdatePhoneRequest {
    1: string user_id,
    2: string phone_number,
}

struct UpdatePhoneResponse {
    1: base.BaseResponse base_resp,
}

struct UpdatePasswdRequest {
    1: string user_id,
    2: string old_passwd,
    3: string new_passwd,
}

struct UpdatePasswdResponse {
    1: base.BaseResponse base_resp,
}

struct UpdateBirthDayRequest {
    1: string user_id,
    2: i32 year,
    3: i32 month,
    4: i32 day,
}

struct UpdateBirthDayResponse {
    1: base.BaseResponse base_resp,
}

struct UploadAvatarRequest {
    1: string user_id,
}

struct UploadAvatarResponse {
    1: base.BaseResponse base_resp,
    2: string avatar_url,
    3: string avatar_id,
    4: string object_name
}

struct UpdateAvatarInfoRequest {
    1: string user_id,
    2: string avatar_id,
    3: string object_name,
    4: i8 blob_type,
}

struct UpdateAvatarInfoResponse {
    1: base.BaseResponse base_resp,
}

struct GetAvatarRequest {
    1: string user_id,
}

struct GetAvatarResponse {
    1: base.BaseResponse base_resp,
    2: string avatar_url,
}

service UserService {
    RegisterResponse RegisterPhone(1: RegisterPhoneRequest req)
    RegisterResponse RegisterEmail(1: RegisterEmailRequest req)
    LoginResponse LoginPhone(1: LoginPhoneRequest req)
    LoginResponse LoginEmail(1: LoginEmailRequest req)
    UpdateEmailResponse UpdateEmail(1: UpdateEmailRequest req)
    UpdatePhoneResponse UpdatePhone(1: UpdatePhoneRequest req)
    UpdatePasswdResponse UpdatePasswd(1: UpdatePasswdRequest req)
    UpdateBirthDayResponse UpdateBirthDay(1: UpdateBirthDayRequest req)
    UploadAvatarResponse UploadAvatar(1: UploadAvatarRequest req)
    UpdateAvatarInfoResponse UpdateAvatarInfo(1: UpdateAvatarInfoRequest req)
    GetAvatarResponse GetAvatar(1: GetAvatarRequest req)
}