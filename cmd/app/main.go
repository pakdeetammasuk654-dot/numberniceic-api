package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2" // Dependency สำหรับ Template Engine
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/services"
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/handlers"
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/repositories"
)

// initDB เชื่อมต่อกับ PostgreSQL โดยใช้ Environment Variables
// (โค้ดส่วนนี้ยังคงเดิม)
func initDB() *sql.DB {
	// ... (โค้ด initDB เดิม) ...

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
	defer db.Close()

	// ------------------------------------
	// Template Engine Setup (แก้ไข)
	// ------------------------------------
	// 1. Initialize HTML template engine: ระบุโฟลเดอร์ views และนามสกุล .html
	engine := html.New("./views", ".html")

	// **NEW: บังคับโหลดและ Parse Template ทั้งหมด**
	if err := engine.Load(); err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	// 2. ตั้งค่า Fiber ให้ใช้ View Engine
	app := fiber.New(fiber.Config{
		Views: engine, // ใช้ engine ที่เราเพิ่งสร้าง
	})

	app.Static("/public", "./public")

	// ------------------------------------
	// Hexagonal Wiring (จากขั้นตอนก่อนหน้า)
	// ...
	satNumRepo := repositories.NewSatNumRepoPostgres(db)
	satNumService := services.NewSatNumService(satNumRepo)
	satNumHandler := handlers.NewSatNumHandler(satNumService)

	// ------------------------------------
	// Fiber Routes
	// ------------------------------------

	// 1. Landing Page Route (หน้าแรก)
	app.Get("/", func(c *fiber.Ctx) error {
		// c.Render(ชื่อไฟล์ Template, data, ชื่อบล็อก Layout)
		return c.Render("pages/index", fiber.Map{
			"Title": "หน้าแรก - NumberNiceIC",
		}, "main") // <--- เปลี่ยน Layout Name ให้ตรงกับ {{define "main"}} ใน main.html
	})

	// 2. API Routes
	v1 := app.Group("/api/v1")
	satNumGroup := v1.Group("/sat-nums")

	// Endpoints สำหรับ SatNum
	satNumGroup.Get("/", satNumHandler.GetAllSatNums)
	satNumGroup.Get("/:key", satNumHandler.GetSatNumByCharKey)

	// Start Server
	log.Fatal(app.Listen(":3000"))
}
