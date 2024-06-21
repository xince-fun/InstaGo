namespace go errno

enum Err {
    Success = 0,
    BadRequest = 10000,
    ParamsErr = 10002,
    ServiceErr = 20000,
    UserDBErr = 30000,
    UserSrvErr = 30001,
    UserPwdErr = 30002,
    UserPwdSameErr = 30003,
    BlobSrvErr = 40000,
    RecordNotFound = 80000,
    RecordExist = 80001,
    InvalidDate = 90001,
}