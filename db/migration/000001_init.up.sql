CREATE TABLE "account" (
  "id" bigserial UNIQUE PRIMARY KEY,
  "login" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "rsakeys" (
  "id" bigserial UNIQUE PRIMARY KEY,
  "owner" bigserial,
  "key" varchar
);

CREATE TABLE "folder" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial,
  "parent" bigint,
  "name" varchar NOT NULL DEFAULT 'New Folder',
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "path" varchar NOT NULL,
  "tag" varchar
);

CREATE TABLE "file" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial,
  "parent" bigserial,
  "name" varchar NOT NULL DEFAULT 'New File',
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "path" varchar NOT NULL,
  "tag" varchar
);

CREATE INDEX ON "account" ("id");

CREATE INDEX ON "rsakeys" ("owner");

CREATE INDEX ON "folder" ("id");

CREATE INDEX ON "folder" ("owner");

CREATE INDEX ON "folder" ("tag");

CREATE INDEX ON "file" ("id");

CREATE INDEX ON "file" ("owner");

CREATE INDEX ON "file" ("parent");

CREATE INDEX ON "file" ("tag");

ALTER TABLE "rsakeys" ADD FOREIGN KEY ("owner") REFERENCES "account" ("id");

ALTER TABLE "folder" ADD FOREIGN KEY ("owner") REFERENCES "account" ("id");

ALTER TABLE "folder" ADD FOREIGN KEY ("parent") REFERENCES "folder" ("id");

ALTER TABLE "file" ADD FOREIGN KEY ("owner") REFERENCES "account" ("id");

ALTER TABLE "file" ADD FOREIGN KEY ("parent") REFERENCES "folder" ("id");
