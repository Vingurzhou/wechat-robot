CREATE TABLE msg (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    msg_id INTEGER,
    from_user_name TEXT,
    to_user_name TEXT,
    msg_type INTEGER,
    content TEXT,
    status INTEGER,
    img_status INTEGER,
    img_buf_i_len INTEGER,
    create_time INTEGER,
    msg_source TEXT,
    push_content TEXT,
    new_msg_id INTEGER,
    msg_seq INTEGER,
    wxid TEXT,
    appid TEXT,
    type_name TEXT
);

CREATE TABLE chatroom (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chatroom_id TEXT,
    nick_name TEXT,
    py_initial TEXT,
    quan_pin TEXT,
    sex INTEGER,
    remark TEXT,
    remark_py_initial TEXT,
    remark_quan_pin TEXT,
    chat_room_notify INTEGER,
    chat_room_owner TEXT,
    small_head_img_url TEXT
);

CREATE TABLE member (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    wxid TEXT,
    nick_name TEXT NOT NULL,
    inviter_user_name TEXT,
    member_flag INTEGER NOT NULL,
    display_name TEXT,
    big_head_img_url TEXT,
    small_head_img_url TEXT,
    chatroom_id TEXT,
    is_owner BOOLEAN NOT NULL DEFAULT 0,
    is_admin BOOLEAN NOT NULL DEFAULT 0
);
CREATE UNIQUE INDEX idx_chatroom_id_wxid ON member (chatroom_id, wxid);
