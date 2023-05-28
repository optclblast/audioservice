CREATE TABLE "account" (
  "id" bigserial UNIQUE PRIMARY KEY,
  "login" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "RSAKeys" (
  "owner" bigserial PRIMARY KEY,
  "key" varchar
);

CREATE TABLE "folder" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial,
  "parent" bigint,
  "name" varchar NOT NULL DEFAULT 'New Folder',
  "access_level" varchar NOT NULL DEFAULT 'DEFAULT',
  "created_at" timestamp NOT NULL DEFAULT 'now()',
);

CREATE TABLE "file" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial,
  "parent" bigserial,
  "name" varchar NOT NULL DEFAULT 'New File',
  "content" bigint NOT NULL,
  "tag" varchar
);

CREATE INDEX ON "account" ("id");

CREATE INDEX ON "RSAKeys" ("owner");

CREATE INDEX ON "folder" ("id");

CREATE INDEX ON "folder" ("owner");

CREATE INDEX ON "folder" ("tag");

CREATE INDEX ON "file" ("id");

CREATE INDEX ON "file" ("owner");

CREATE INDEX ON "file" ("parent");

CREATE INDEX ON "file" ("tag");

ALTER TABLE "RSAKeys" ADD FOREIGN KEY ("owner") REFERENCES "account" ("id");

ALTER TABLE "folder" ADD FOREIGN KEY ("owner") REFERENCES "account" ("id");

ALTER TABLE "folder" ADD FOREIGN KEY ("parent") REFERENCES "folder" ("id");

ALTER TABLE "file" ADD FOREIGN KEY ("owner") REFERENCES "account" ("id");

ALTER TABLE "file" ADD FOREIGN KEY ("parent") REFERENCES "folder" ("id");
