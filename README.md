# mi-api-go

REST API de productos construida en Go con soporte para dos backends de base de datos: **PostgreSQL** y **SurrealDB**. Incluye un frontend embebido servido directamente desde el binario.

## Stack

- **Go 1.26** — lenguaje principal
- **Gin** — framework HTTP
- **PostgreSQL** (via `pgx/v5`) — backend relacional
- **SurrealDB** — backend alternativo
- **embed** — frontend embebido en el binario

## Estructura

```
mi-api-go/
├── main.go           # Punto de entrada, configuración de rutas
├── domain/           # Modelos de datos (Producto)
├── handler/          # Controladores HTTP
├── service/          # Lógica de negocio
├── repository/
│   ├── postgres/     # Implementación PostgreSQL
│   └── surrealdb/    # Implementación SurrealDB
├── db/               # Conexión a las bases de datos
└── frontend/         # UI embebida (HTML, CSS, JS)
```

## Requisitos

- Go 1.21+
- PostgreSQL **o** SurrealDB (según el backend elegido)

## Configuración

Copia `.env.example` y completa los valores:

```bash
cp .env.example .env
```

### Variables de entorno

| Variable | Descripción | Default |
|---|---|---|
| `DB_BACKEND` | Backend a usar: `surreal` o `postgres` | `postgres` |
| `POSTGRES_URL` | URL de conexión PostgreSQL | — |
| `SURREAL_URL` | URL de SurrealDB (wss://...) | — |
| `SURREAL_USER` | Usuario SurrealDB | — |
| `SURREAL_PASS` | Contraseña SurrealDB | — |
| `SURREAL_NS` | Namespace SurrealDB | — |
| `SURREAL_DB` | Base de datos SurrealDB | — |

## Levantar el proyecto

**Con PostgreSQL:**
```bash
POSTGRES_URL=postgres://user:pass@localhost:5432/dbname go run main.go
```

**Con SurrealDB:**
```bash
DB_BACKEND=surreal \
SURREAL_URL=wss://tu-instancia.surreal.cloud \
SURREAL_USER=root \
SURREAL_PASS=tu_password \
SURREAL_NS=portfolio \
SURREAL_DB=productos \
go run main.go
```

**Usando archivo .env:**
```bash
export $(cat .env | xargs) && go run main.go
```

La API y el frontend quedan disponibles en `http://localhost:8080`.

## Endpoints

| Método | Ruta | Descripción |
|---|---|---|
| `GET` | `/api/productos` | Listar todos los productos |
| `GET` | `/api/productos/:id` | Obtener producto por ID |
| `POST` | `/api/productos` | Crear producto |
| `PUT` | `/api/productos/:id` | Actualizar producto |
| `DELETE` | `/api/productos/:id` | Eliminar producto |

### Modelo Producto

```json
{
  "id": "string",
  "nombre": "string",
  "precio": 0.0
}
```

### Ejemplos

```bash
# Listar productos
curl http://localhost:8080/api/productos

# Crear producto
curl -X POST http://localhost:8080/api/productos \
  -H "Content-Type: application/json" \
  -d '{"nombre": "Laptop", "precio": 999.99}'

# Actualizar producto
curl -X PUT http://localhost:8080/api/productos/1 \
  -H "Content-Type: application/json" \
  -d '{"nombre": "Laptop Pro", "precio": 1299.99}'

# Eliminar producto
curl -X DELETE http://localhost:8080/api/productos/1
```

## Deploy

El proyecto incluye `Dockerfile` y `fly.toml` para despliegue en [Fly.io](https://fly.io).

```bash
fly deploy
```
