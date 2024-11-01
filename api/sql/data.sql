
insert into users (name, nickname, email, password) values 
('José', 'jose_123', 'jose@gmail.com', '123456'),
('Maria', 'maria_123', 'maria@gmail.com', '123456'),
('João', 'joao_123', 'joao@gmail.com', '12345678'),
('Mateus', 'mateus_123', 'mateus@gmail.com', '12345678');

insert into followers (user_id, follower_id) values
(1, 2),
(1, 3),
(3, 1),
(3, 2),
(3, 4);

insert into posts(title,content,author_id) values
('Post 1', 'Content 1', 1),
('Post 2', 'Content 2', 2),
('Post 3', 'Content 3', 3),
('Post 4', 'Content 4', 4);

insert into likes(user_id,post_id) values (5,1),(6,1),(7,1)

SELECT p.*, u.nickname, count(l.id) as likes
FROM posts p 
INNER JOIN users u ON u.id = p.author_id 
LEFT JOIN likes l on l.post_id = p.id
WHERE p.id = 1

