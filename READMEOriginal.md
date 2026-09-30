# 🚀 Interactive Cloud Resume API & Web Engine

> **High-Performance Server-Driven UI & Cloud-Native REST API in Go**  
> Diseñado para demostrar despliegues hiperligeros, renderizado nativo tipado con **Templ**, dinamismo reactivo con **HTMX**, y contenedores seguros con **Distroless** (<20MB).

---

## 🏛️ Arquitectura y Patrones Técnicos

```
                      ┌─────────────────────────────────────────┐
                      │              Client Browser             │
                      └────────────────────┬────────────────────┘
                                           │
                        HTTP / HTMX (GET / & /components/skills)
                                           │
                                           ▼
┌────────────────────────────────────────────────────────────────────────────────┐
│                           Go Application Core                                 │
│                                                                                │
│   ┌────────────────────────────────────────────────────────────────────────┐  │
│   │                      Middleware Pipeline                               │  │
│   │   1. Recovery (Chi)                                                    │  │
│   │   2. SecurityHeaders (CORS, CSP, HSTS, X-Frame-Options: DENY)          │  │
│   │   3. Structured Logger (Latency with time.Since, X-Request-ID, Real IP)│  │
│   └──────────────────────────────────┬─────────────────────────────────────┘  │
│                                      │                                         │
│                                      ▼                                         │
│   ┌──────────────────────────────────┴─────────────────────────────────────┐  │
│   │                           Chi Router                                   │  │
│   │                                                                        │  │
│   │   [GET /]                   --> views.Home (Full Templ Render)         │  │
│   │   [GET /components/skills]  --> components.SkillsGrid (HTMX Partial)   │  │
│   │   [GET /api/v1/health]      --> handlers.HealthHandler (MemStats/SLA)  │  │
│   │   [GET /api/v1/cv]          --> handlers.CVHandler (JSON Resume)       │  │
│   │   [GET /static/*]           --> Embedded Filesystem (embed.FS)         │  │
│   └────────────────────────────────────────────────────────────────────────┘  │
└────────────────────────────────────────────────────────────────────────────────┘
```

### 1. Server-Driven UI con Go + Templ + HTMX
* **Cero sobrecarga de SPA**: En lugar de empaquetar megabytes de React/Vue/Node, el servidor Go compila plantillas `.templ` a código Go tipado y libre de asignaciones pesadas.
* **Actualizaciones hiperrápidas**: El filtrado por categorías de habilidades (`Backend`, `Bases de Datos`, `Cloud/SRE`) utiliza `hx-get="/components/skills?cat=backend"`, devolviendo únicamente el fragmento HTML necesario (`HTML over the wire`) con latencias inferiores a 1ms.
* **Binario 100% Autónomo**: Los recursos estáticos (CSS) se encuentran embebidos directamente en el binario mediante `embed.FS`.

### 2. Middleware & Observabilidad
* **`LoggerMiddleware`**: Inyecta y propaga `X-Request-ID` (UUID), calcula la latencia con `time.Since()`, extrae la IP real detrás de proxies (`X-Forwarded-For`), monitorea códigos de estado HTTP y genera logs estructurados JSON con `log/slog`.
* **`SecurityHeadersMiddleware`**:
  * `Content-Security-Policy`: Restringido pero habilitado para HTMX y fuentes seguras.
  * `X-Frame-Options: DENY`: Prevención de Clickjacking.
  * `X-Content-Type-Options: nosniff`: Prevención de MIME sniffing.
  * `Strict-Transport-Security`: HSTS habilitado.
  * `Access-Control-Allow-Origin`: Políticas CORS abiertas para consumo de la API REST pública.

### 3. Contenedorización Distroless (<20MB)
* **Multi-Stage Dockerfile**: 
  * **Fase 1 (Builder)**: `golang:1.24-alpine` con `CGO_ENABLED=0` y compilación estática (`-ldflags="-s -w -extldflags '-static'"`).
  * **Fase 2 (Runtime)**: `gcr.io/distroless/static-debian12:nonroot` sin shell, sin paquetes innecesarios, ejecutado bajo un usuario no-root para máxima seguridad.

---

## 📡 Endpoints de la API

| Método | Ruta | Descripción | Tipo de Respuesta |
| :--- | :--- | :--- | :--- |
| `GET` | `/` | Dashboard web interactivo | `text/html; charset=utf-8` |
| `GET` | `/components/skills?cat={cat}` | Fragmento HTMX para filtrado de skills (`all`, `backend`, `database`, `cloud_sre`) | `text/html; charset=utf-8` |
| `GET` | `/api/v1/health` | Diagnóstico de salud, uptime, goroutines y uso de memoria | `application/json` |
| `GET` | `/api/v1/cv` | Payload JSON completo con datos estructurados del CV | `application/json` |

---

## 🛠️ Guía de Ejecución

### Prerrequisitos
* **Go 1.22+**
* **Templ CLI**: `go install github.com/a-h/templ/cmd/templ@latest`
* **Docker** (opcional, para despliegue en contenedor)

### Ejecución Local

```bash
# 1. Generar plantillas Templ y compilar
make build

# 2. Ejecutar la aplicación
make run
```
Abre en tu navegador: [http://localhost:8080](http://localhost:8080)

### Pruebas Automatizadas

```bash
make test
```

### Despliegue con Docker

```bash
# Construir la imagen Distroless
make docker-build

# Ejecutar el contenedor
make docker-run
```

---

## 📊 Ejemplo de Respuesta `/api/v1/health`

```json
{
  "status": "UP",
  "version": "1.0.0",
  "uptime": "2h14m32s",
  "server_time": "2026-09-14T01:15:00Z",
  "go_version": "go1.24.0",
  "diagnostics": {
    "goroutines": 4,
    "allocated_memory": "2.41 MB",
    "system_memory": "14.50 MB",
    "num_gc": 1
  }
}
```
