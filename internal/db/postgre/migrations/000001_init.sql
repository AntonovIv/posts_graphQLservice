-- +goose Up
-- +goose StatementBegin

drop table if exists Posts;

drop table if exists Comments;

create table Posts (
    id serial primary key, 
    name varchar(255) not null,
    content text not null,
    author varchar(64),
    comments_allowed boolean
);

create table Comments (
    id serial primary key,
    content varchar(2000) not null,
    author varchar(64),
    postid int not null,
    foreign key (postid) references Posts(id), 
    reply_to int, 
    foreign key(reply_to) references Comments(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table Posts;

drop table Comments;
-- +goose StatementEnd