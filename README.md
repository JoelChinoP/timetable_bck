# 📅 Timetable Backend

Backend para la gestión de horarios universitarios desarrollado con **Go**, **GraphQL** y **PostgreSQL**.

## 🚀 Tecnologías

- **Go** - Lenguaje de programación
- **GraphQL** - API con [gqlgen](https://gqlgen.com/)
- **Gin** - Framework HTTP
- **PostgreSQL** - Base de datos
- **SQLC** - Generador de código SQL type-safe

## 📁 Estructura del Proyecto

```
├── cmd/                  # Punto de entrada de la aplicación
├── graph/                # Schema y resolvers de GraphQL
│   ├── schema.graphqls   # Definición del schema GraphQL
│   └── resolver.go       # Implementación de resolvers
├── internal/db/          # Capa de base de datos
│   ├── schema.sql        # Schema de PostgreSQL
│   └── queries/          # Queries SQL para SQLC
└── pkg/                  # Utilidades compartidas
```

## ⚙️ Instalación

```bash
# Clonar el repositorio
git clone https://github.com/JoelChinoP/timetable_bck.git
cd timetable_bck

# Instalar dependencias
go mod download

# Ejecutar la aplicación
go run cmd/main.go
```

## 🔧 Desarrollo

```bash
# Generar código GraphQL
go run github.com/99designs/gqlgen generate

# Generar código SQLC
sqlc generate
```

## 📡 Endpoints

| Método | Ruta     | Descripción        |
| ------ | -------- | ------------------ |
| GET    | `/`      | GraphQL Playground |
| POST   | `/query` | API GraphQL        |
