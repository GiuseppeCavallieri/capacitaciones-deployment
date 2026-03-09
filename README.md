# Veterinaria Citiaps — Capacitación Deployment

Aplicación de gestión veterinaria construida con Go (backend), Nuxt 3 (frontend) y MongoDB. Todo se levanta con un solo comando de Docker.
Permite administrar las dueños, las mascotas de estos y las vacunas que cada mascota lleva.

---

## ¿Cómo correr la aplicación?

Primero se debe cambiar a la rama gcavallieri-mongodb o gcavallieri-postgresql
Solo se necesita tener docker desktop abierto. A través de una terminal, hay que ir al directorio "Docker" de este repositorio y ejecutar:

```bash
cd Docker
docker compose up --build -d
```

Tarda un poco al hacer el build. Las rutas quedan disponibles para ser accedidas mediante:

| Qué                      | URL                   |
| ------------------------ | --------------------- |
| 🌐 Aplicación (Frontend) | http://localhost:3000 |
| ⚙️ API (Backend)         | http://localhost:8080 |

---

## ¿Cómo bajar los contenedores?

```bash
cd Docker
docker compose down
```

Para borrar los datos de la base de datos:

```bash
docker compose down -v
```

---

## ¿Cómo volver a levantarlos sin el build?

```bash
cd Docker
docker compose up -d
```

---

## Rutas de la API

| Método | Ruta                | Descripción             |
| ------ | ------------------- | ----------------------- |
| GET    | /perros             | Lista todos los perros  |
| GET    | /perros/:id         | Detalle de un perro     |
| POST   | /perros             | Crear un perro          |
| PUT    | /perros/:id         | Editar un perro         |
| DELETE | /perros/:id         | Eliminar un perro       |
| GET    | /perros/:id/dueno   | Dueño de un perro       |
| GET    | /perros/:id/vacunas | Vacunas de un perro     |
| GET    | /duenos             | Lista todos los dueños  |
| GET    | /duenos/:id         | Detalle de un dueño     |
| POST   | /duenos             | Crear un dueño          |
| PUT    | /duenos/:id         | Editar un dueño         |
| DELETE | /duenos/:id         | Eliminar un dueño       |
| GET    | /vacunas            | Lista todas las vacunas |
| GET    | /vacunas/:id        | Detalle de una vacuna   |
| POST   | /vacunas            | Crear una vacuna        |
| PUT    | /vacunas/:id        | Editar una vacuna       |
| DELETE | /vacunas/:id        | Eliminar una vacuna     |

---

## Estructura de carpetas

- `Config/Backend` — Código fuente del backend (Go + Gin + MongoDB)
- `Config/Frontend` — Código fuente del frontend (Nuxt 3 + Bulma)
- `Docker` — docker-compose y script de inicialización de la base de datos
- `Databases` — Dumps de MongoDB y PostgreSQL

---

## Datos Personales

Nombre y Apellido: Giuseppe Cavallieri
Correo: giuseppe.cavallieri@usach.cl

---

## Backup PostgreSQL

Restaurar el backup usando la herramienta 'Restore' de pgAdmin.

## Backup MongoDB

Restaurar el backup inyectando los JSON en la base de datos
