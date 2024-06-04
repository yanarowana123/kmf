CREATE TABLE r_currency (
    id INT NOT NULL IDENTITY(1,1) PRIMARY KEY,
    title varchar(60) not null,
    code varchar(3) not null,
    value numeric(18,2) not null,
    a_date date not null
);