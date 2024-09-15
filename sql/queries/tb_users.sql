-- name: GetUserByEmail :one
select usr_email, usr_id from `tb_users` where usr_email = ? limit 1;

-- name: UpdateUserStatusByUserId :exec
update `tb_users` 
set 
    usr_status = ?,
    usr_updated_at = ?
where usr_id = ?;
