-- +goose Up
CREATE TABLE dept (
    deptno serial PRIMARY KEY,
    dname VARCHAR(50) NOT NULL,
    loc VARCHAR(50) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS dept;