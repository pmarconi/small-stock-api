package routes

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"stock-manager-api/database"
	"stock-manager-api/handlers"
	"stock-manager-api/utils"

	"time"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var privateKey *rsa.PrivateKey

// New create an instance of Book app routes
func New() *fiber.App {
	// var (
	// 	buf    bytes.Buffer
	// 	logger = log.New(&buf, "logger: ", log.Lshortfile)
	// )

	utils.Warning().Println("Loading router")

	// Create private key
	rng := rand.Reader
	var err error
	privateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	// log.Print("privateKey: %v", privateKey)

	// Create fiber application
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/Argentina",
	}))

	// app.Get("/docs/*", swagger.Handler)
	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "🐣 v1",
		})
		return c.Next()
	})

	v1.Post("/register", func(c *fiber.Ctx) error {
		type RegisterReqBody struct {
			Username string `json:"username" xml:"username" form:"username"`
			Password string `json:"password" xml:"password" form:"password"`
		}
		b := new(RegisterReqBody)
		if err := c.BodyParser(b); err != nil {
			return err
		}
		hash, err := HashPassword(b.Password)
		passwordHash := hash
		// log.Print("hash: %v", passwordHash)
		if err != nil {
			// logger.Print("Hello, log file!", hash)
			// fmt.Print(&buf)
		}

		// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		newUser := database.NewUser()
		newUser.Username = b.Username
		newUser.Password = passwordHash

		database.CreateUser(newUser)

		// if result.Error() != "" {
		// 	return c.JSON(fiber.Map{"success": false})
		// }

		return c.JSON(fiber.Map{"success": true})
	})

	v1.Post("/login", func(c *fiber.Ctx) error {

		type LoginReqBody struct {
			Username string `json:"username" xml:"username" form:"username"`
			Password string `json:"password" xml:"password" form:"password"`
		}

		b := new(LoginReqBody)
		if err := c.BodyParser(b); err != nil {
			return err
		}

		// u := c.FormValue("username")
		// p := c.FormValue("password")
		u := b.Username
		p := b.Password

		user, err := database.GetUserByUsername(u)
		if err != nil {
			return c.JSON(fiber.Map{"success": false})
		}

		var hasCorrectUserAndPassword bool = p != "" && CheckPasswordHash(p, user.Password)
		if !hasCorrectUserAndPassword {
			// return c.SendStatus(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{"success": false})
		}

		// Create the Claims
		claims := jwt.MapClaims{
			"name":  "John Doe",
			"admin": true,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString(privateKey)
		if err != nil {
			log.Printf("token.SignedString: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"success": true,
			"token":   t,
		})

		// return c.Next()
	})

	// JWT Middleware
	v1.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    privateKey.Public(),
	}))

	// v1.Get("/books", handlers.GetAllBooks)
	// v1.Get("/books/:id", handlers.GetBookByID)
	// v1.Post("/books", handlers.RegisterBook)
	// v1.Delete("/books/:id", handlers.DeleteBook)

	v1.Get("/products", handlers.GetAllProducts)
	v1.Get("/products/:id", handlers.GetProductByID)

	utils.Warning().Println("Good")

	return app
}
