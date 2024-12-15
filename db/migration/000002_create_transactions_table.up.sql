CREATE TABLE IF NOT EXISTS transactions(
id uuid primary key,
account_id INTEGER not null,
operation_type INTEGER not null,
amount NUMERIC(10,2) not null,
event_date TIMESTAMP not null,
FOREIGN KEY(account_id) REFERENCES accounts(id)
);