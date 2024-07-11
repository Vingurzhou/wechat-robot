CREATE TABLE IF NOT EXISTS `group`  (
    id TEXT PRIMARY KEY , 
    username TEXT NOT NULL ,
    nickname TEXT NOT NULL, 
    display_name TEXT NOT NULL, 
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `user` (
    id TEXT PRIMARY KEY , 
    group_id TEXT NOT NULL ,
    username TEXT NOT NULL ,
    nickname TEXT NOT NULL ,
    display_name TEXT NOT NULL, 
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_group_id_nickname ON user(group_id, nickname);

CREATE TABLE IF NOT EXISTS `msg` (
    id TEXT PRIMARY KEY , 
    group_id TEXT NOT NULL ,
    username TEXT NOT NULL ,
    nickname TEXT NOT NULL ,
    display_name TEXT NOT NULL, 
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_group_id_nickname ON user(group_id, nickname);