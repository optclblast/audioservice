CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY,
  "login" varchar NOT NULL,
  "password" varchar NOT NULL,
  "signupdate" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Artists" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "bio" varchar,
  "picture" varchar
);

CREATE TABLE "Collections" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "artist" bigint NOT NULL,
  "ft_artists" bigint,
  "type" varchar NOT NULL DEFAULT 'Album',
  "discription" varchar,
  "lenght" bigint,
  "label" varchar NOT NULL,
  "cover" varchar,
  "date" timestamp NOT NULL DEFAULT 'now()',
  "upload_date" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Tracks" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "artist" bigint NOT NULL,
  "ft_artists" bigint,
  "album" bigint,
  "cover" varchar,
  "location" varchar NOT NULL,
  "upload_date" timestamp NOT NULL DEFAULT 'now()'
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
  "user_id" bigserial,
  "tracks_id" bigserial
);

CREATE TABLE "UsersLikedCollections" (
  "user_id" bigserial,
  "collections" bigserial
);

CREATE TABLE "UsersLikedPlaylists" (
  "user_id" bigserial,
  "playlists" bigserial
);

CREATE INDEX ON "Users" ("login");

CREATE INDEX ON "Artists" ("name");

CREATE INDEX ON "Collections" ("name");

CREATE INDEX ON "Collections" ("artist");

CREATE INDEX ON "Collections" ("ft_artists");

CREATE INDEX ON "Collections" ("type");

CREATE INDEX ON "Tracks" ("name");

CREATE INDEX ON "Tracks" ("artist");

CREATE INDEX ON "Tracks" ("ft_artists");

CREATE INDEX ON "Playlists" ("name");

CREATE INDEX ON "Playlists" ("owner");

ALTER TABLE "Collections" ADD FOREIGN KEY ("artist") REFERENCES "Artists" ("id");

ALTER TABLE "Collections" ADD FOREIGN KEY ("ft_artists") REFERENCES "Artists" ("id");

ALTER TABLE "Tracks" ADD FOREIGN KEY ("artist") REFERENCES "Artists" ("id");

ALTER TABLE "Tracks" ADD FOREIGN KEY ("ft_artists") REFERENCES "Artists" ("id");

ALTER TABLE "Tracks" ADD FOREIGN KEY ("album") REFERENCES "Collections" ("id");

ALTER TABLE "Playlists" ADD FOREIGN KEY ("owner") REFERENCES "Users" ("id");

ALTER TABLE "UsersLikedTracks" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "UsersLikedTracks" ADD FOREIGN KEY ("tracks_id") REFERENCES "Tracks" ("id");

ALTER TABLE "UsersLikedCollections" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "UsersLikedCollections" ADD FOREIGN KEY ("collections") REFERENCES "Collections" ("id");

ALTER TABLE "UsersLikedPlaylists" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "UsersLikedPlaylists" ADD FOREIGN KEY ("playlists") REFERENCES "Playlists" ("id");
