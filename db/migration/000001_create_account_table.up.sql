CREATE TABLE IF NOT EXISTS accounts(
    id serial primary key,
    document_number VARCHAR(255) not null,
    created_at TIMESTAMP DEFAULT NOW()
);