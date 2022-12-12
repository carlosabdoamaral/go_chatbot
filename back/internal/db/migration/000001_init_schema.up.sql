CREATE DATABASE ia_using_go;

use ia_using_go;
CREATE TABLE message (
    "id" bigserial PRIMARY KEY,
    "message" VARCHAR(255) NOT NULL,
    "answer" VARCHAR(255) NOT NULL
);

INSERT INTO message(message, answer) VALUES ('OI', 'EAI, TUDO BEM?');