CREATE TABLE users 
(
	id serial not null unique,
	name varchar not null,
	username varchar not null unique,
	password_hash varchar not null
);

CREATE TABLE channels
(
	id serial not null unique,
	link varchar not null
);

CREATE TABLE videos
(
	id serial not null unique,
	channel_id int references channels (id) on delete cascade not null,
	link varchar not null
);
