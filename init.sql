CREATE TABLE IF NOT EXISTS alumnos (
  id SERIAL PRIMARY KEY,
  nombre TEXT,
  apellido TEXT,
  dni TEXT,
  nacimiento DATE,
  sexo TEXT,
  email TEXT
);

