# 🛵 Proyecto Rappi Clone con Go – Roadmap Técnico

---

## 🎯 Objetivo General

Construir una plataforma tipo **Rappi/Uber Eats** desde cero, utilizando Go como backend, PostgreSQL como base de datos, Kafka o RabbitMQ como sistema de colas, y un frontend moderno (React/Next.js).

---

## 🧱 Épicas del Proyecto

---

### 🟦 ÉPICA 1: Fundamentos de Go y monolito inicial

**Historias de Usuario**
- Como desarrollador, quiero aprender la sintaxis de Go para comenzar a construir.
- Como usuario, quiero registrarme, loguearme y crear pedidos simples.

**Tareas**
- [ ] Instalar Go y configurar entorno
- [ ] Crear servidor HTTP
- [ ] Endpoints básicos: registro, login, crear pedido
- [ ] Crear base de datos PostgreSQL y tablas iniciales
- [ ] JWT y autenticación

---

### 🟨 ÉPICA 2: Separación en Microservicios

**Historias de Usuario**
- Como arquitecto, quiero que cada funcionalidad esté separada para escalar mejor.

**Tareas**
- [ ] Dividir en servicios: auth, users, orders, products, notifications
- [ ] Crear repositorios independientes para cada microservicio
- [ ] Dockerizar cada servicio
- [ ] Crear Docker Compose para entorno local
- [ ] Configurar API Gateway / Nginx reverse proxy

---

### 🟥 ÉPICA 3: Concurrencia y Workers

**Historias de Usuario**
- Como repartidor, quiero ver y aceptar pedidos de forma rápida y concurrente.

**Tareas**
- [ ] Usar goroutines y channels para procesamiento concurrente de pedidos
- [ ] Crear sistema de matching: clientes ↔ repartidores
- [ ] Workers para tareas en segundo plano
- [ ] Control de concurrencia con mutex / waitgroups

---

### 🟩 ÉPICA 4: Kafka / RabbitMQ y Arquitectura de Eventos

**Historias de Usuario**
- Como sistema, quiero comunicar servicios con eventos para escalar horizontalmente.

**Tareas**
- [ ] Levantar Kafka o RabbitMQ en Docker
- [ ] Productor: “pedido creado”
- [ ] Consumidor: notificaciones o sistema de tracking
- [ ] Eventos: pedido confirmado, entregado, cancelado
- [ ] Manejo de errores, retries, DLQ

---

### 🟧 ÉPICA 5: Frontend moderno

**Historias de Usuario**
- Como usuario, quiero tener una interfaz amigable para hacer pedidos.
- Como repartidor, quiero aceptar y ver pedidos.

**Tareas**
- [ ] Crear frontend con Next.js + TypeScript + Tailwind
- [ ] Vista cliente: ver menú, agregar al carrito, hacer pedido
- [ ] Vista repartidor: ver pedidos disponibles, aceptar
- [ ] Vista comercio: publicar productos, ver pedidos recibidos
- [ ] Conexión con backend por REST

---

### 🟪 ÉPICA 6: Observabilidad, Testing y Deploy Local

**Historias de Usuario**
- Como equipo técnico, quiero monitorear el sistema y asegurar su calidad.

**Tareas**
- [ ] Logs estructurados con zerolog
- [ ] Test unitarios en capa de dominio y servicios
- [ ] Test de integración con base de datos
- [ ] Observabilidad con Prometheus + Grafana
- [ ] Docker Compose para frontend, backend y colas
- [ ] README completo para instalación local

---

## 📦 Microservicios Detallados

| Servicio | Funcionalidad | Base de Datos | Cola |
|----------|----------------|---------------|------|
| Auth     | Registro, login, JWT | PostgreSQL | - |
| Usuarios | Clientes, repartidores, comercios | PostgreSQL | - |
| Pedidos  | Gestión de órdenes, carrito | PostgreSQL | ✅ |
| Productos| Menú y stock | PostgreSQL | - |
| Pagos    | Simulación de pagos | PostgreSQL | - |
| Tracking | Mock de geolocalización | (opcional) | ✅ |
| Notificaciones | Emails y push | - | ✅ |

---

## 🗺️ Roadmap por Semana (Sugerencia)

| Semana | Objetivo |
|--------|----------|
| 1      | Aprender Go + HTTP + estructuras |
| 2      | CRUD de productos + pedidos + base de datos |
| 3      | JWT Auth + perfiles de usuario |
| 4      | Concurrencia y workers |
| 5      | Kafka/RabbitMQ entre pedidos y notificaciones |
| 6      | Dividir en microservicios + Docker Compose |
| 7      | Frontend cliente con Next.js |
| 8      | Frontend repartidor y comercio |
| 9      | Observabilidad + testing |
| 10     | Revisión final + mejoras + deploy local

---

## ✅ Extras (si tenés más tiempo)

- Stripe / MercadoPago (mock de pagos)
- WebSockets para notificaciones en tiempo real
- Chat entre cliente y repartidor (opcional)
- Deploy en Railway, Fly.io o VPS

---

## 🧰 Stack

**Backend**
- Go 1.22+
- PostgreSQL
- Kafka o RabbitMQ
- Docker

**Frontend**
- Next.js (React)
- TypeScript
- Tailwind CSS

**DevOps**
- Docker Compose
- Prometheus + Grafana
- GitHub Actions (CI/CD opcional)

---

## ✨ Tips

- Empezá con un monolito si sos nuevo en Go.
- Escribí pruebas al menos en la capa de dominio y servicios.
- Evitá acoplar servicios: usá eventos.
- Documentá bien cada microservicio y su puerto/API.

