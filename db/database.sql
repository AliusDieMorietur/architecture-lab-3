CREATE TABLE "User" (
  "id" bigint generated always as identity,
  "name" varchar NULL
);

ALTER TABLE "User" ADD CONSTRAINT "pkUser" PRIMARY KEY ("id");

CREATE TABLE "Forum" (
  "id" bigint generated always as identity,
  "name" varchar NULL,
  "topicKeyword" varchar NULL
);

ALTER TABLE "Forum" ADD CONSTRAINT "pkForum" PRIMARY KEY ("id");

CREATE TABLE "UserForum" (
  "userId" bigint NOT NULL,
  "forumId" bigint NOT NULL
);

ALTER TABLE "UserForum" ADD CONSTRAINT "pkUserForum" PRIMARY KEY ("userId", "forumId");
ALTER TABLE "UserForum" ADD CONSTRAINT "fkUserForumUser" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE CASCADE;
ALTER TABLE "UserForum" ADD CONSTRAINT "fkUserForumForum" FOREIGN KEY ("forumId") REFERENCES "Forum" ("id") ON DELETE CASCADE;


