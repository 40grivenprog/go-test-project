CREATE TABLE positions (
   id INTEGER primary key,
   name VARCHAR(50) NOT NULL,
   salary INTEGER NOT NULL,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
);
