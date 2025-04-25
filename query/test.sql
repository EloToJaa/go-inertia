-- name: GetTestById :one
select *
from test_table
where id = $1
limit 1
;

