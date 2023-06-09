namespace go book

struct QueryBookReq {
    1: i32 Num (api.query="num", api.vd="$<100; msg:'num must less than 100'");
}

struct QueryBookResp {
    1: bool   Exist;
    2: string ID;
    3: string Title;
    4: string Author;
}

struct InsertBookReq {
    1: string ID (api.form="id");
    2: string Title (api.form="title");
    3: string Author (api.form="author");
}

struct InsertBookResp {
    1: bool Ok; 
    2: string Msg; 
}

service BookSvc {
   QueryBookResp queryBook(1: QueryBookReq req) (api.get="/book/query");
   InsertBookResp insertBook(1: InsertBookReq req) (api.post="/book/insert");
}