-- +goose Up
-- +goose StatementBegin

insert into Posts (
    name, 
    content, 
    author, 
    comments_allowed) 
    values (
        'post name 1', 
        'post content 1', 
        'post author 1', 
        true);
        
insert into Posts (
    name, 
    content, 
    author, 
    comments_allowed) 
    values (
        'post name 2', 
        'post content 2', 
        'post author 2', 
        true);
 
insert into Posts (
    name, 
    content, 
    author, 
    comments_allowed) 
    values (
        'post name 3', 
        'post content 3', 
        'post author 3', 
        true);
 
insert into Comments (
    content, 
    author, 
    postid, 
    reply_to) 
    values (
        '1 comment for post  1', 
        'comment author 1', 
        1, 
        null);

insert into Comments (
    content, 
    author, 
    postid, 
    reply_to) 
    values (
        '2 comment for post  1', 
        'comment author 1', 
        1, 
        null);

insert into Comments (
    content, 
    author, 
    postid, 
    reply_to) 
    values (
        '1 comment for post  2', 
        'comment author 1', 
        2, 
        null);

insert into Comments (
    content, 
    author, 
    postid, 
    reply_to) 
    values (
        '2 comment for post  2', 
        'comment author 1', 
        2, 
        null);
insert into Comments (
    content, 
    author, 
    postid, 
    reply_to) 
    values (
        '1 comment for 1 commentfor post 2', 
        'comment author 1', 
        2, 
        4);
insert into Comments (
    content, 
    author, 
    postid, 
    reply_to) 
    values (
        '2 comment for 1 commentfor post 2', 
        'comment author 1', 
        2, 
        4)


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table Posts, Comments;
-- +goose StatementEnd