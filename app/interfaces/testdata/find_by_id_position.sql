CREATE TABLE positions (
   id BIGSERIAL primary key,
   name VARCHAR(50) NOT NULL,
   salary INTEGER NOT NULL,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
);

INSERT INTO positions (name, salary) VALUES ('Ruby Engineer', 100);
