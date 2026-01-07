-- name: CreateUser :one
insert into users(id, created_at, updated_at, email, hashed_password)
values(
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2
)
returning *;

-- name: GetUserByID :one
select *
from users
where id = $1;

-- name: GetUserByEmail :one
select *
from users
where email = $1;

-- name: GetUserFromRefreshToken :one
select users.id, users.created_at, users.updated_at, users.email, users.hashed_password
from users
join refresh_tokens on users.id = refresh_tokens.user_id
where refresh_tokens.token = $1
    and revoked_at is null
    and expires_at > now();

-- name: UpdateUserEmailPassword :exec
update users
set hashed_password = $1,
email = $2
where id = $3;

-- name: UpgradeUser :exec
update users
set is_chirpy_red = true
where id = $1;