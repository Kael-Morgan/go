CREATE TABLE users (
  id smallserial PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  middle_name VARCHAR(50),
  last_name VARCHAR(50) NOT NULL,
  birthday TIMESTAMP,
  username VARCHAR(50) NOT NULL,
  password TEXT NOT NULL,
  email VARCHAR(100) UNIQUE,
  phone_number VARCHAR(15),
  gender VARCHAR(7),
  is_active BOOLEAN DEFAULT TRUE,
  is_deleted BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP ,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE conversation (
  id smallserial PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  is_active BOOLEAN DEFAULT TRUE,
  is_deleted BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE participant (
  id smallserial PRIMARY KEY,
  user_id INTEGER NOT NULL,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE message (
  id smallserial PRIMARY KEY,
  participant_id INTEGER NOT NULL,
  conversation_id INTEGER NOT NULL,
  type INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (participant_id) REFERENCES participant(id),
  FOREIGN KEY (conversation_id) REFERENCES conversation(id)
);