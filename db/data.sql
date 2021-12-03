INSERT INTO "User" (
  "name"
) VALUES (
  'Kirill'
),(
  'Roma'
),(
  'Ilja'
);

INSERT INTO "Forum" (
  "name",
  "topicKeyword"
) VALUES (
  'Anime',
  'anime'
), (
  'Cartoons',
  'cartoons'
), (
  'Squid Game',
  'squid_game'
);

INSERT INTO "UserForum" (
  "userId",
  "forumId"
) VALUES (1, 1), (1, 2), (1, 3), (2, 1), (2, 2), (3, 1);
