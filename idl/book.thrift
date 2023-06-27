namespace go book

struct QueryBookReq {
    1: i32 Num (api.query="num", api.vd="$<100; msg:'num must less than 100'");
}

struct QueryBookResp {
    1: string ID;
    2: string Title;
    3: string Author;
    4: string Content; 
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
   QueryBookResponse queryBook(1: QueryBookRequest req) (api.get="book/query");
   InsertBookResponse insertBook(1: InsertBookRequest req) (api.post="book/insert");
}