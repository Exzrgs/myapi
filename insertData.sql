INSERT INTO articles (title, contents, username, nice, created_at) VALUES
    ('firstPost','This is my first blog','saki',2,now());

INSERT INTO articles (title, contents,username,nice) VALUES
    ('second','This is second blog','saki',1);

INSERT INTO comments (article_id, message, created_at) VALUES
    (1,'1st comment',now());

INSERT INTO comments (article_id,message) VALUES
    (1,'hello');