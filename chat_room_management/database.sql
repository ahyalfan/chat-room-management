-- Active: 1723868453146@@127.0.0.1@5432@server_room_chat

CREATE DATABASE server_chat_room;

CREATE TABLE users (
    "id" bigserial PRIMARY KEY,
    "username" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL
)