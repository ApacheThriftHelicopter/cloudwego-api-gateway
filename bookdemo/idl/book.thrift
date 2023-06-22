namespace go book

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct Book {
    1: i64 book_id
    2: i64 user_id
    3: string username
    4: string user_avatar
    5: string title
    6: string content
    7: i64 create_time
}

struct CreateBookRequest {
    1: string title (vt.min_size = "1")
    2: string content (vt.min_size = "1")
    3: i64 user_id (vt.gt = "0")
}

struct CreateBookResponse {
    1: BaseResp base_resp
}

struct DeleteBookRequest {
    1: i64 book_id (vt.gt = "0")
    2: i64 user_id
}

struct DeleteBookResponse {
    1: BaseResp base_resp
}

struct UpdateBookRequest {
    1: i64 book_id (vt.gt = "0")
    2: i64 user_id
    3: optional string title
    4: optional string content
}

struct UpdateBookResponse {
    1: BaseResp base_resp
}

struct QueryBookRequest {
    1: i64 user_id (vt.gt = "0")
    2: optional string search_key
    3: i64 offset (vt.ge = "0")
    4: i64 limit (vt.ge = "0")
}

struct QueryBookResponse {
    1: list<Book> books
    2: i64 total
    3: BaseResp base_resp
}

struct MGetBookRequest {
    1: list<i64> book_ids (vt.min_size = "1")
}

struct MGetBookResponse {
    1: list<Book> books
    2: BaseResp base_resp
}

service BookService {
    CreateBookResponse CreateBook(1: CreateBookRequest req)
    DeleteBookResponse DeleteBook(1: DeleteBookRequest req)
    UpdateBookResponse UpdateBook(1: UpdateBookRequest req)
    QueryBookResponse QueryBook(1: QueryBookRequest req)
    MGetBookResponse MGetBook(1: MGetBookRequest req)
}