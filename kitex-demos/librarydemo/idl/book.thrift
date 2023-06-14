// IDL for librarydemo API

namespace go librarydemo

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct Book {
    1: i64 book_id
    2: i64 user_id
    3: string user_name
    4: string title 
}

struct CreateBookRequest {
    1: string title
    2: i64 user_id
}

struct CreateBookResponse {
    1: BaseResp base_resp
}

struct DeleteNoteRequest {
    1: i64 book_id
    2: i64 user_id
}

struct DeleteBookResponse {
    1: BaseResp base_resp
}

struct MGetBookRequest {
    2: list<i64> book_ids
}

struct MGetBookResponse {
    1: list<Book> books 
    2: BaseResp base_resp
}

struct QueryBookRequest {
    1: i64 user_id
    2: optional string search_key
    3: i64 offset
    4: i64 limit 
}

struct QueryBookResponse {
    1: list<Book> books
    2: i64 total
    3: BaseResp base_resp
}

service BookService {
    CreateBookResponse CreateBook(1: CreateBookRequest req)
    MGetBookResponse MGetBook(1: MGetBookRequest req)
    DeleteBookResponse DeleteBook(1: DeleteNoteRequest req)
    QueryBookResponse QueryBook(1: QueryBookRequest req)
}