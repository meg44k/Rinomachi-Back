
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,     -- インデックス、主キー、自動増分
    user_id VARCHAR(255) NOT NULL UNIQUE,  -- ユーザID（ユニーク制約付き）
    user_name VARCHAR(50) NOT NULL,        -- ユーザ名（最大50文字）
    password VARCHAR(255) NOT NULL,        -- パスワード
    email VARCHAR(255) NOT NULL UNIQUE     -- メールアドレス（ユニーク制約付き）
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE buildings (
    id INT AUTO_INCREMENT PRIMARY KEY, -- インデックス、主キー、自動増分
    building_id VARCHAR(255) NOT NULL UNIQUE, -- 建物固有のID
    address VARCHAR(50) NOT NULL UNIQUE,  -- 建物の住所（最大50文字）
    structure VARCHAR(255) NOT NULL,  -- 建物の構造（例: 木造）
    floors INT NOT NULL,              -- 階数
    age INT DEFAULT NULL,             -- 築年数（NULL許容）
    area FLOAT DEFAULT NULL,          -- 床面積（NULL許容）
    contract VARCHAR(255) DEFAULT NULL, -- 契約（NULL許容）
    description TEXT DEFAULT NULL,    -- 建物説明（テキスト型、NULL許容）
    is_available BOOLEAN NOT NULL DEFAULT TRUE, -- 募集中かどうか（デフォルトはTRUE）
    price INT DEFAULT NULL            -- 値段（NULL許容）
    favorites INT DEFAULT 0 NOTNULL, --お気に入り数
    transportation TEXT DEFAULT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE favorites (
    id INT AUTO_INCREMENT PRIMARY KEY,     -- インデックス、主キー、自動増分
    user_id VARCHAR(255) NOT NULL,         -- ユーザID
    building_id VARCHAR(255) NOT NULL,     -- 建物固有のID
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE, -- 外部キー: ユーザテーブル
    FOREIGN KEY (building_id) REFERENCES buildings(building_id) ON DELETE CASCADE ON UPDATE CASCADE -- 外部キー: 空き家テーブル
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE histories (
    id INT AUTO_INCREMENT PRIMARY KEY,     -- インデックス、主キー、自動増分
    user_id VARCHAR(255) NOT NULL,         -- ユーザID
    building_id VARCHAR(255) NOT NULL,     -- 建物固有のID
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 追加・変更時間
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE, -- 外部キー: ユーザテーブル
    FOREIGN KEY (building_id) REFERENCES buildings(building_id) ON DELETE CASCADE ON UPDATE CASCADE -- 外部キー: 空き家テーブル
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
