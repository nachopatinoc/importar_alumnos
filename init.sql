CREATE TABLE IF NOT EXISTS alumnos (
  id SERIAL PRIMARY KEY,
  apellido TEXT,
  nombre TEXT,
  nro_documento TEXT,
  tipo_documento TEXT,
  fecha_nacimiento DATE,
  sexo CHAR(1),
  nro_legajo INT UNIQUE NOT NULL,
  fecha_ingreso DATE
);

