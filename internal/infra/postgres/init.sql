CREATE DATABASE rinha;

CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  "limit" INTEGER NOT NULL,
  balance INTEGER NOT NULL DEFAULT 0
);

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
