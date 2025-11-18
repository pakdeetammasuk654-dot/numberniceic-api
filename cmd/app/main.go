package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/services"
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/handlers"
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/repositories"
)

// initDB เชื่อมต่อกับ PostgreSQL โดยใช้ Environment Variables
func initDB() *sql.DB {
	// 1. Load Environment Variables (จากไฟล์ .env)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 2. Build Connection String (DSN)
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// 3. Open Database Connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// 4. Ping เพื่อยืนยันการเชื่อมต่อ
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL!")

	// ตั้งค่าจำนวน Connection (เป็นทางเลือกแต่แนะนำ)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	return db
}

func main() {
	// Initialize Database Connection
	db := initDB()
	defer db.Close() // ปิด Connection เมื่อโปรแกรมหยุดทำงาน

	// ------------------------------------
	// Hexagonal Wiring (การประกอบร่าง)
	// ------------------------------------

	// 1. Repository Adapter (Output Adapter - เชื่อม DB)
	satNumRepo := repositories.NewSatNumRepoPostgres(db)

	// 2. Core Service (Core Logic - Inject Repository Port)
	satNumService := services.NewSatNumService(satNumRepo)

	// 3. Handler Adapter (Input Adapter - เชื่อม Fiber)
	satNumHandler := handlers.NewSatNumHandler(satNumService)

	// ------------------------------------
	// Fiber Setup and Routes
	// ------------------------------------
	app := fiber.New()

	// 0. Base Route (จากโค้ดเดิม)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, NumberNiceIC API is running!")
	})

	// 4. Grouping API Routes (แนะนำให้ใช้ versioning)
	v1 := app.Group("/api/v1")
	satNumGroup := v1.Group("/sat-nums")

	// Endpoints สำหรับ SatNum
	satNumGroup.Get("/", satNumHandler.GetAllSatNums)          // GET /api/v1/sat-nums
	satNumGroup.Get("/:key", satNumHandler.GetSatNumByCharKey) // GET /api/v1/sat-nums/:key

	// Start Server
	log.Fatal(app.Listen(":3000"))
}
