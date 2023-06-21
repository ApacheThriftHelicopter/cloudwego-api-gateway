namespace go bookapi

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct User {
    1: i64 user_id
    2: string username
    3: string avatar
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

struct CreateUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct CreateUserResponse {
    1: BaseResp base_resp
}

struct CheckUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct CheckUserResponse {
    1: BaseResp base_resp
}

struct CreateBookRequest {
    1: string title (api.vd="len($) > 0")
    2: string content (api.vd="len($) > 0")
    3: i64 user_id
}

struct CreateBookResponse {
    1: BaseResp base_resp
}

struct QueryBookRequest {
    1: i64 user_id
    2: optional string search_key (api.query="search_key", api.vd="len($) > 0")
    3: i64 offset (api.query="offset", api.vd="len($) >= 0")
    4: i64 limit (api.query="limit", api.vd="len($) >= 0")
}

struct QueryBookResponse {
    1: list<Book> books
    2: i64 total
    3: BaseResp base_resp
}

struct UpdateBookRequest {
    1: i64 book_id (api.path="book_id")
    2: i64 user_id
    3: optional string title (api.form="title", api.vd="len($) > 0")
    4: optional string content (api.form="content", api.vd="len($) > 0")
}

struct UpdateBookResponse {
    1: BaseResp base_resp
}

struct DeleteBookRequest {
    1: i64 book_id (api.path="book_id")
    2: i64 user_id
}

struct DeleteBookResponse {
    1: BaseResp base_resp
}

service ApiService {
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/v1/user/register")
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/v1/user/login")
    CreateBookResponse CreateBook(1: CreateBookRequest req) (api.post="/v1/book")
    QueryBookResponse QueryBook(1: QueryBookRequest req) (api.get="/v1/book/query")
    UpdateBookResponse UpdateBook(1: UpdateBookRequest req) (api.put="/v1/book/:book_id")
    DeleteBookResponse DeleteBook(1: DeleteBookRequest req) (api.delete="/v1/book/:book_id")
}