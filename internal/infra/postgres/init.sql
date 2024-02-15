CREATE DATABASE rinha;

CREATE TABLE IF NOT EXISTS clients (
  id SERIAL PRIMARY KEY,
  "limit" INTEGER NOT NULL,
  balance INTEGER NOT NULL DEFAULT 0
);

DO $$
  BEGIN
    IF NOT EXISTS (SELECT * FROM clients WHERE id BETWEEN 1 AND 5) THEN
      INSERT INTO clients ("limit") 
      VALUES 
      (1000 * 100),
      (800 * 100),
      (10000 * 100),
      (100000 * 100),
      (5000 * 100);
    END IF;
  END;
  $$;
