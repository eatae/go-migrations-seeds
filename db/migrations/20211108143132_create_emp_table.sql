-- +goose Up
CREATE TABLE emp (
    empno serial PRIMARY KEY,
    ename VARCHAR(50) NOT NULL,
    job varchar(50) NOT NULL,
    mgr int,
    hiredate DATE NOT NULL,
    sal int NOT NULL,
    comm int,
    deptno int NOT NULL,
    FOREIGN KEY (mgr) REFERENCES emp (empno),
    FOREIGN KEY (deptno) REFERENCES dept (deptno)
);

-- +goose Down
DROP TABLE IF EXISTS emp;