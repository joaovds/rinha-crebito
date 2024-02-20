CREATE DATABASE rinha;

CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  "limit" INTEGER NOT NULL,
  balance INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  value INTEGER NOT NULL,
  type_transaction char(1) NOT NULL CHECK (type_transaction IN ('c', 'd')),
  description VARCHAR(10) NOT NULL,
  realized_at timestamp NOT NULL DEFAULT now(),
  account_id INTEGER NOT NULL,
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE INDEX idx_realized_at ON transactions (realized_at);

DO $$
  BEGIN
    IF NOT EXISTS (SELECT * FROM accounts WHERE id BETWEEN 1 AND 5) THEN
      INSERT INTO accounts ("limit") 
      VALUES 
      (1000 * 100),
      (800 * 100),
      (10000 * 100),
      (100000 * 100),
      (5000 * 100);
    END IF;
  END;
  $$;
