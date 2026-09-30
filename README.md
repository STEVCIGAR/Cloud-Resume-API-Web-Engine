# Cloud-Resume-API-Web-Engine
**High-Performance Server-Driven UI & Cloud-Native REST API in Go**  
Diseñado para demostrar despliegues hiperligeros, renderizado nativo tipado con **Templ**, dinamismo reactivo con **HTMX**, y contenedores seguros con **Distroless** (<20MB).

### Prerrequisitos
* **Go 1.22+
* **Templ CLI: `go install github.com/a-h/templ/cmd/templ@latest`
* **Docker** (opcional, para despliegue en contenedor)

## Ejecución Local

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

## Ejemplo de Respuesta `/api/v1/health`

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
