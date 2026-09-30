package data

import "strings"

// SkillCategory represents the category of technical skills.
type SkillCategory string

const (
	CategoryAll      SkillCategory = "all"
	CategoryBackend  SkillCategory = "backend"
	CategoryDatabase SkillCategory = "database"
	CategoryCloudSRE SkillCategory = "cloud_sre"
)

// Skill represents an individual technical competence.
type Skill struct {
	Name        string        `json:"name"`
	Category    SkillCategory `json:"category"`
	Proficiency int           `json:"proficiency"` // 1-100 percentage
	Level       string        `json:"level"`       // "Expert", "Advanced", "Proficient"
	Description string        `json:"description"`
	Tags        []string      `json:"tags"`
}

// Metric represents a quantifiable engineering achievement.
type Metric struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Experience represents a career milestone.
type Experience struct {
	Role        string   `json:"role"`
	Company     string   `json:"company"`
	Location    string   `json:"location"`
	Period      string   `json:"period"`
	Highlights  []string `json:"highlights"`
	TechStack   []string `json:"tech_stack"`
}

// Project represents a featured system architecture or open-source contribution.
type Project struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Impact      string   `json:"impact"`
	TechStack   []string `json:"tech_stack"`
	GitHubURL   string   `json:"github_url,omitempty"`
}

// Profile represents personal and biographical information.
type Profile struct {
	Name       string            `json:"name"`
	Title      string            `json:"title"`
	Bio        string            `json:"bio"`
	Location   string            `json:"location"`
	Email      string            `json:"email"`
	Available  bool              `json:"available"`
	Socials    map[string]string `json:"socials"`
	SummaryKey string            `json:"summary_key"`
}

// Resume is the aggregate root for the whole CV document.
type Resume struct {
	Profile     Profile      `json:"profile"`
	Metrics     []Metric     `json:"metrics"`
	Skills      []Skill      `json:"skills"`
	Experiences []Experience `json:"experiences"`
	Projects    []Project    `json:"projects"`
}

// DefaultResume returns the canonical resume dataset.
var DefaultResume = Resume{
	Profile: Profile{
		Name:      "Alexander Vance",
		Title:     "Lead Backend Engineer & Cloud Architect",
		Bio:       "Specialized in high-throughput distributed systems, event-driven architectures, and resilient cloud-native platforms in Go. Passionate about zero-downtime deployments, high concurrency, and server-driven ergonomics.",
		Location:  "Remote / San Francisco, CA",
		Email:     "alexander.vance.eng@gmail.com",
		Available: true,
		Socials: map[string]string{
			"github":   "https://github.com/cloud-resume",
			"linkedin": "https://linkedin.com/in/cloud-resume",
			"twitter":  "https://x.com/cloud_engineer",
		},
		SummaryKey: "Go 1.24+ · Distributed Systems · High Concurrency · Kubernetes · Multi-Region Architecture",
	},
	Metrics: []Metric{
		{
			Label:       "Production Uptime",
			Value:       "99.995%",
			Description: "Maintained across 4 multi-region clusters with zero unplanned outages",
			Icon:        "shield-check",
		},
		{
			Label:       "Peak Throughput",
			Value:       "150k+ QPS",
			Description: "Sustained sub-12ms p99 latency in event-streaming pipeline",
			Icon:        "zap",
		},
		{
			Label:       "Microservices Managed",
			Value:       "35+",
			Description: "Built with Go, gRPC, and containerized on Kubernetes",
			Icon:        "cpu",
		},
		{
			Label:       "Cost Reduction",
			Value:       "-42%",
			Description: "Architected memory footprint optimizations and autoscaling",
			Icon:        "trending-down",
		},
	},
	Skills: []Skill{
		// Backend
		{
			Name:        "Golang (Go)",
			Category:    CategoryBackend,
			Proficiency: 98,
			Level:       "Expert",
			Description: "Goroutines, Channels, Low-allocation profiling, CGO-free binaries, Generics",
			Tags:        []string{"Concurrency", "Memory Optimization", "gRPC", "HTTP/3"},
		},
		{
			Name:        "gRPC & Protobuf",
			Category:    CategoryBackend,
			Proficiency: 92,
			Level:       "Expert",
			Description: "High-performance inter-service RPCs, bidirectional streaming, custom interceptors",
			Tags:        []string{"Streaming", "IDL", "Schema Evolution"},
		},
		{
			Name:        "Distributed Systems & Event-Driven",
			Category:    CategoryBackend,
			Proficiency: 90,
			Level:       "Advanced",
			Description: "Saga patterns, CQRS, transactional outbox, idempotent event consumers",
			Tags:        []string{"Kafka", "NATS", "RabbitMQ", "Outbox Pattern"},
		},
		{
			Name:        "Server-Driven UI (Templ + HTMX)",
			Category:    CategoryBackend,
			Proficiency: 95,
			Level:       "Expert",
			Description: "Type-safe Go template compilation, hypermedia-driven dynamic UIs without heavy SPA overhead",
			Tags:        []string{"Templ", "HTMX", "Zero-JS Bundling"},
		},

		// Databases
		{
			Name:        "PostgreSQL",
			Category:    CategoryDatabase,
			Proficiency: 94,
			Level:       "Expert",
			Description: "Connection pooling (pgx), CTEs, partition tables, EXPLAIN ANALYZE tuning, index strategies",
			Tags:        []string{"pgx", "Partitioning", "Query Optimization", "B-Tree/GIN"},
		},
		{
			Name:        "Redis & DragonFly",
			Category:    CategoryDatabase,
			Proficiency: 90,
			Level:       "Advanced",
			Description: "Distributed locks (Redlock), caching layers, rate limiting with sliding windows, streams",
			Tags:        []string{"Caching", "Pub/Sub", "Rate Limiting", "Lua Scripts"},
		},
		{
			Name:        "ClickHouse & TimescaleDB",
			Category:    CategoryDatabase,
			Proficiency: 86,
			Level:       "Proficient",
			Description: "Columnar OLAP query optimization, real-time telemetry processing and metrics ingestion",
			Tags:        []string{"OLAP", "Time-Series", "Analytics"},
		},

		// Cloud & SRE
		{
			Name:        "Kubernetes & Docker",
			Category:    CategoryCloudSRE,
			Proficiency: 92,
			Level:       "Expert",
			Description: "Multi-stage scratch/distroless builds, Helm charts, CRDs, Pod disruption budgets, HPA",
			Tags:        []string{"K8s", "Distroless", "Helm", "GitOps"},
		},
		{
			Name:        "Observability (OpenTelemetry / Prometheus)",
			Category:    CategoryCloudSRE,
			Proficiency: 90,
			Level:       "Advanced",
			Description: "Distributed tracing context propagation, RED metrics, custom Prometheus collectors, Grafana dashboards",
			Tags:        []string{"Otel", "Prometheus", "Tracing", "Grafana"},
		},
		{
			Name:        "Terraform & AWS/GCP",
			Category:    CategoryCloudSRE,
			Proficiency: 88,
			Level:       "Advanced",
			Description: "Infrastructure as Code, multi-region VPC peering, EKS, Cloudflare Edge routing, IAM least-privilege",
			Tags:        []string{"IaC", "AWS", "GCP", "Cloudflare"},
		},
		{
			Name:        "CI/CD & Release Engineering",
			Category:    CategoryCloudSRE,
			Proficiency: 91,
			Level:       "Expert",
			Description: "Automated vulnerability scanning (Trivy, Govulncheck), semantic versioning, canary releases with ArgoCD",
			Tags:        []string{"GitHub Actions", "ArgoCD", "Canary", "Govulncheck"},
		},
	},
	Experiences: []Experience{
		{
			Role:     "Lead Backend Engineer",
			Company:  "Apex Cloud Infrastructure",
			Location: "San Francisco, CA (Remote)",
			Period:   "2023 - Present",
			Highlights: []string{
				"Architected core ingestion engine in Go, processing 150k+ QPS with sub-15ms p99 latency across 3 global regions.",
				"Migrated legacy monolithic service to Go microservices, reducing cloud compute expenditures by 42%.",
				"Designed zero-downtime canary deployment pipeline with automated rollbacks using ArgoCD and Prometheus alerts.",
				"Mentored team of 8 engineers on Go concurrency paradigms, memory profiling (pprof), and clean architecture.",
			},
			TechStack: []string{"Go", "gRPC", "PostgreSQL", "Kafka", "Kubernetes", "OpenTelemetry", "AWS"},
		},
		{
			Role:     "Senior Distributed Systems Engineer",
			Company:  "HyperScale Data Labs",
			Location: "New York, NY",
			Period:   "2020 - 2023",
			Highlights: []string{
				"Engineered high-concurrency event-streaming pipeline handling 2B+ daily events using Go and Apache Kafka.",
				"Created distributed lock and rate-limiting middleware used by 20+ internal microservices.",
				"Implemented custom PostgreSQL connection pooling and read-replica routing layer, boosting database throughput by 65%.",
			},
			TechStack: []string{"Go", "PostgreSQL", "Redis", "Kafka", "Docker", "Terraform", "GCP"},
		},
		{
			Role:     "Backend Software Engineer",
			Company:  "Nova Fintech Platforms",
			Location: "Austin, TX",
			Period:   "2018 - 2020",
			Highlights: []string{
				"Developed double-entry ledger financial transactions service with strict ACID guarantees and cryptographic verification.",
				"Created automated integration testing framework running in ephemeral Docker containers.",
			},
			TechStack: []string{"Go", "PostgreSQL", "Docker", "REST API", "CI/CD"},
		},
	},
	Projects: []Project{
		{
			Title:       "Dynamic Server-Driven Cloud Resume Engine",
			Description: "Ultra-fast, zero-JS-framework cloud dashboard built with Go, Templ, and HTMX. Demonstrates high-performance Server-Driven UI and distroless multi-stage containerization.",
			Impact:      "Sub-millisecond render times, <15MB Docker image, 100/100 Lighthouse score.",
			TechStack:   []string{"Go", "Templ", "HTMX", "Docker", "Chi"},
			GitHubURL:   "https://github.com/cloud-resume/api",
		},
		{
			Title:       "Resilient Event Gateway (Go-StreamGate)",
			Description: "High-throughput reverse proxy and rate-limiting gateway with dynamic token bucket algorithm and circuit breaking.",
			Impact:      "Handles 200k req/sec with <2MB memory per worker node.",
			TechStack:   []string{"Go", "Redis", "eBPF", "Prometheus"},
			GitHubURL:   "https://github.com/cloud-resume/streamgate",
		},
	},
}

// GetSkillsByCategory returns skills filtered by category, or all if empty/all.
func (r *Resume) GetSkillsByCategory(cat string) []Skill {
	cat = strings.ToLower(strings.TrimSpace(cat))
	if cat == "" || cat == "all" {
		return r.Skills
	}

	var filtered []Skill
	for _, skill := range r.Skills {
		if string(skill.Category) == cat {
			filtered = append(filtered, skill)
		}
	}
	return filtered
}
