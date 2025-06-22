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

#### Si estas en Linux ejecutá lo siguiente para las instalaciones:

##### Actualizar e instalar dependencias básicas
```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y curl wget git ca-certificates gnupg lsb-release software-properties-common
```

##### 👉 Instalar Go 1.21+
```bash
wget https://go.dev/dl/go1.21.10.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.10.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

##### Verificar Go
```bash
go version
```

##### 👉 Instalar Docker
```bash
curl -fsSL https://get.docker.com | sudo sh
```

##### Verificar Docker
```bash
sudo docker version
```

##### 👉 Instalar Docker Compose (plugin moderno para Docker)
```bash
sudo apt install -y docker-compose-plugin
```

##### Verificar Docker Compose
```bash
docker compose version
```

##### 👉 Instalar Git (opcional si ya lo tenés)
```bash
sudo apt install -y git
git --version
```

---

### 2. Crear el archivo `.env`

En la raíz del proyecto, crear un archivo llamado `.env` con el siguiente contenido:

```bash
DB_HOST=localhost
DB_PORT=5434
DB_USER=user_sysacad
DB_PASS=pass_sysacad
DB_NAME=DEV_SYSACAD
```

---

### 3. Levantar la base de datos con Docker

Desde la raíz del proyecto, correr:
```bash
docker-compose up -d
```

---

### 4. Colocar el archivo CSV real

Ubicar el archivo original (el que contiene los 2.5 millones de registros) en la siguiente carpeta del proyecto:

```bash
/data/alumnos.csv
```

---

### ▶️ 5. Ejecutar el programa

Desde la raíz del proyecto, correr:
```bash
go run main.go
```

---

### 🔄 6. Antes de volver a ejecutar

Antes de volver a ejecutar los tests, se recomienda limpiar el entorno para evitar datos residuales o conexiones pendientes. Ejecutar:

```bash
docker compose down
docker compose up -d
```
Esto reiniciará los contenedores desde cero y asegurará que todo funcione limpio.

### ✅ 7. Verificar y eliminar los alumnos ingresados

Una vez finalizada la ejecución, podés entrar a la base de datos y consultar la cantidad de alumnos cargados.

Entrar al contenedor de PostgreSQL:

```bash
docker exec -it dev_sysacad psql -U user_sysacad -d DEV_SYSACAD
```

Si se quieren contar los alumnos que hay cargados en la base de datos se usa:

```SQL
SELECT COUNT(*) FROM alumnos;
```

En el caso de querer eliminar los alumnos ingresados en la tabla ejecutar dentro del prompt de PostgreSQL:

```SQL
DELETE FROM ALUMNOS;
```

Para salir del contenedor de PostgresSQL escribir:
```SQL
\q 
```

---

### 🧪 8. Ejecutar los tests

Para ejercutar los test usar:

go test ./test

Para ejecutarlos desde la raíz:

go test ./...

Qué testea este proyecto:

- ✅ Lectura y parseo correcto de un CSV

- ✅ Validación de cada campo (legajo, fechas, sexo, tipo documento)

- ✅ Inserción en la base de datos (batch insert)

- ✅ Conexión correcta a la base (ConectarDB())

---

### 🚨 9. Notas y advertencias importantes

- ❌ No se utiliza COPY de PostgreSQL (no está permitido por la consigna)

- ✅ El insert se hace por lote (batch insert), lo que acelera la carga

- 🧪 El proyecto fue desarrollado usando TDD: cada función crítica tiene su test