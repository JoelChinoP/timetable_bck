package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	sqlc "github.com/JoelChinoP/timetable_bck/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool
	once sync.Once
)

// Config contiene la configuración de la base de datos
type Config struct {
	Host, Port, User, Password, Database string
	MaxConns, MinConns                   int32
	MaxConnLifetime, MaxConnIdleTime     time.Duration
}

// InitDB inicializa el pool de conexiones (se ejecuta solo una vez)
func InitDB(ctx context.Context, cfg Config) (err error) {

	once.Do(func() {
		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
		)

		poolCfg, parseErr := pgxpool.ParseConfig(dsn)
		if parseErr != nil {
			err = fmt.Errorf("unable to parse config: %w", parseErr)
			return
		}

		poolCfg.MaxConns = cfg.MaxConns
		poolCfg.MinConns = cfg.MinConns
		poolCfg.MaxConnLifetime = cfg.MaxConnLifetime
		poolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

		pool, err = pgxpool.NewWithConfig(ctx, poolCfg)
		if err != nil {
			err = fmt.Errorf("unable to create connection pool: %w", err)
			return
		}

		// Verificar conexión y seed si es necesario
		queries := sqlc.New(pool)
		count, countErr := queries.CountAcademicHours(ctx)
		if countErr != nil {
			err = fmt.Errorf("unable to verify database connection: %w", countErr)
			return
		}
		if count == 0 {
			if seedErr := queries.SeedAcademicHours(ctx); seedErr != nil {
				err = fmt.Errorf("unable to seed academic_hours table: %w", seedErr)
				return
			}
		}

		log.Printf("Database connection pool initialized successfully: %s", cfg.Database)
	})

	return err
}

// GetPool devuelve el pool de conexiones
func GetPool() *pgxpool.Pool {
	if pool == nil {
		log.Fatal("Database pool not initialized. Call InitDB first.")
	}
	return pool
}

// GetQueries devuelve una instancia de Queries para usar con sqlc
func GetQueries() *sqlc.Queries {
	return sqlc.New(GetPool())
}

// Close cierra el pool de conexiones
func Close() {
	if pool != nil {
		pool.Close()
		pool = nil
		log.Println("Database connection pool closed")
	}
}
