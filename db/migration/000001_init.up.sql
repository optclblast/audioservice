CREATE TABLE "Users" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "login" varchar NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "Authors" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "bio" varchar
);

CREATE TABLE "Collections" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "authon" bigint NOT NULL,
  "ft_authors" bigint,
  "type" varchar NOT NULL DEFAULT 'Album',
  "discription" varchar,
  "lenght" bigint,
  "label" varchar NOT NULL,
  "date" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Tracks" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "authon" bigint NOT NULL,
  "ft_authors" bigint,
  "album" bigint,
  "location" varchar NOT NULL
);

CREATE TABLE "Playlists" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "owner" bigint NOT NULL,
  "discription" varchar,
  "lenght" bigint,
  "date" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "UsersLikedTracks" (
  "id" bigserial PRIMARY KEY,
  "user" bigserial,
  "tracks" bigserial
);

CREATE TABLE "UsersLikedCollections" (
  "id" bigserial PRIMARY KEY,
  "user" bigserial,
  "collections" bigserial
);
