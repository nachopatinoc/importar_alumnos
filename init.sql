CREATE TABLE IF NOT EXISTS alumnos (
  apellido TEXT,
  nombre TEXT,
  nro_documento TEXT UNIQUE,
  tipo_documento TEXT,
  fecha_nacimiento DATE,
  sexo CHAR(1),
  nro_legajo INT NOT NULL PRIMARY KEY,
  fecha_ingreso DATE
);

