# 📥 Trabajo Práctico – Importador de Alumnos desde CSV

Este proyecto importa registros de alumnos desde un archivo `.csv` hacia una base de datos PostgreSQL, procesando datos en lotes (batch) con validaciones por campo. Fue desarrollado aplicando buenas prácticas como:

- DRY (Don't Repeat Yourself)
- KISS (Keep It Simple, Stupid)
- YAGNI (You Ain't Gonna Need It)
- TDD (Test Driven Development)

---

## 🧠 Objetivo del proyecto

Leer millones de registros desde un CSV y cargarlos en la tabla `alumnos` de la base `DEV_SYSACAD` de forma eficiente y robusta, con control de errores y validaciones estrictas.

---

## ⚙️ Paso a paso para usar el proyecto

### 1. Requisitos previos

- Go 1.21+
- Docker + Docker Compose
- Git (opcional, si se clona desde un repo)

---

### 2. Crear el archivo `.env`

En la raíz del proyecto, crear un archivo llamado `.env` con el siguiente contenido:

DB_HOST=localhost
DB_PORT=5434
DB_USER=user_sysacad
DB_PASS=postgres
DB_NAME=DEV_SYSACAD

---

### 3. Levantar la base de datos con Docker

Desde la raíz del proyecto, correr:

docker-compose up -d

---

### 4. Colocar el archivo CSV real

Ubicar el archivo original (el que contiene los 2.5 millones de registros) en la siguiente carpeta del proyecto:

/data/alumnos.csv

---

### ▶️ 5. Ejecutar el programa

Desde la raíz del proyecto, correr:

go run main.go

---

### ✅ 6. Verificar que los datos fueron insertados

Una vez finalizada la ejecución, podés entrar a la base de datos y consultar la cantidad de alumnos cargados.

Entrar al contenedor de PostgreSQL:

docker exec -it dev_sysacad psql -U user_sysacad -d DEV_SYSACAD

En el caso de querer eliminar los alumnos ingresados en la tabla ejecutar dentro del prompt de PostgreSQL:

DELETE FROM ALUMNOS

---

### 🧪 7. Ejecutar los tests

Para ejercutar los test usar:

go test ./test

Para ejecutarlos desde la raíz:

go test ./...

Qué testea este proyecto:

- ✅ Lectura y parseo correcto de un CSV

- ✅ Validación de cada campo (legajo, fechas, sexo, tipo documento)

- ✅ Inserción en la base de datos (batch insert)

- ✅ Rechazo de registros duplicados (por legajo o documento)

- ✅ Conexión correcta a la base (ConectarDB())

---

### 🚨 8. Notas y advertencias importantes

- ❌ No se utiliza COPY de PostgreSQL (no está permitido por la consigna)

- ✅ El insert se hace por lote (batch insert), lo que acelera la carga

- 🔒 Los campos nro_legajo está validado y controlado para evitar duplicados (también puede agregarse nro_documento)

- 🧪 El proyecto fue desarrollado usando TDD: cada función crítica tiene su test

