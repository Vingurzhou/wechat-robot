CREATE TABLE IF NOT EXISTS msg (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    MsgId INTEGER,
    FromUserName TEXT,
    ToUserName TEXT,
    MsgType INTEGER,
    Content TEXT,
    Status INTEGER,
    ImgStatus INTEGER,
    ImgBuf_iLen INTEGER,
    CreateTime INTEGER,
    MsgSource TEXT,
    PushContent TEXT,
    NewMsgId INTEGER,
    MsgSeq INTEGER,
    Wxid TEXT,
    Appid TEXT,
    TypeName TEXT
);