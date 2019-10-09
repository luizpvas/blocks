create table users (
	user_id           uuid primary key,
	registration_date timestamp with time zone,
	password_hash     text
);

create table sessions (
	user_id uuid,
	token   text
);

create table records (
	resource_id text,
	record_id   uuid primary key,
	attributes  jsonb,
	created_at  timestamp with time zone,
	updated_at  timestamp with time zone
);