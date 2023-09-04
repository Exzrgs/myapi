CREATE TABLE IF NOT EXISTS articles (
    article_id integer unsigned auto_increment primary key,
    title varchar(100) not null,
    contents text not null,
    username varchar(100) not null,
    nice integer not null,
    created_at  datetime
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id integer unsigned auto_increment primary key,
    article_id integer unsigned not null,
    message text not null,
    created_at datetime,
    foreign key (article_id) references articles(article_id)
);

INSERT INTO articles (title, contents, username, nice, created_at) VALUES
    ('firstPost','This is my first blog','saki',2,now());

INSERT INTO articles (title, contents,username,nice) VALUES
    ('second','This is second blog','saki',1);

INSERT INTO comments (article_id, message, created_at) VALUES
    (1,'1st comment',now());

INSERT INTO comments (article_id,message) VALUES
    (1,'hello');