namespace go blob

include "./../../../shared/idl/base/base.thrift"

struct GeneratePutPreSignedUrlRequest {
    1: string user_id,
    2: i8 blob_type,
    3: i32 timeout,
}

struct GeneratePutPreSignedUrlResponse {
    1: base.BaseResponse base_resp,
    2: string url,
    3: string id,
}

struct GenerateGetPreSignedUrlRequest {
    1: string blob_id,
    2: i32 timeout,
}

struct GenerateGetPreSignedUrlResponse {
    1: base.BaseResponse base_resp,
    2: string url,
}

struct NotifyBlobUploadRequest {
    1: string blob_id,
    2: string user_id,
    3: string url,
}

struct NotifyBlobUploadResponse {
    1: base.BaseResponse base_resp,
}

service BlobService {
    GeneratePutPreSignedUrlResponse GeneratePutPreSignedUrl(1: GeneratePutPreSignedUrlRequest req),
    GenerateGetPreSignedUrlResponse GenerateGetPreSignedUrl(1: GenerateGetPreSignedUrlRequest req),
    NotifyBlobUploadResponse NotifyBlobUpload(1: NotifyBlobUploadRequest req),
}