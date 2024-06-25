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
    UserNotExistErr = 30004,
    BlobSrvErr = 40000,
    RelationDBErr = 50000,
    RelationSrvErr = 50001,
    RelationSelfErr = 50002,
    RelationExistErr = 50003,
    RelationNotExistErr = 50004,
    RecordNotFound = 80000,
    RecordExist = 80001,
    InvalidDate = 90001,
}