package controllers

import (
	"fmt"
	"myGin/models"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func isValidatePassword(password string) bool {
	// 8-20 char
	// at least one lowercase letter, one digit, one special character, one Uppercase
	regex := regexp.MustCompile(`^[[:graph:]]{8,20}$`)
	hasUppercase := regexp.MustCompile(`[[:upper:]]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[[:lower:]]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[^[:alnum:]]`).MatchString(password)
	return hasUppercase && hasLowercase && hasDigit && hasSpecial && regex.MatchString(password)
}

func HashPassword(password string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	hashedPassword := string(salt)
	return hashedPassword, nil
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Create User || Register
func (db *DBController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate Password
	if !isValidatePassword(user.Password) {
		fmt.Println(isValidatePassword(user.Password))
		c.JSON(http.StatusBadRequest, gin.H{"error": "password not match requirement"})
		return
	}

	// Hash Password
	var err error
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash password fail"})
		return
	}

	tx := db.Database.Create(&user)
	if tx.Error != nil {
		fmt.Println("Error Create User:", tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"results": "User created"})
}

func (db *DBController) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		fmt.Println("Login Bind Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bind fail"})
		return
	}

	// Look up user

	var user models.User
	tx := db.Database.First(&user, "email = ?", body.Email)

	if tx.Error != nil || user.ID == 0 {
		fmt.Println("Look up user Error", tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	if !VerifyPassword(user.Password, body.Password) {
		fmt.Println("VerifyPassword user Error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("Generate Token Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Response
	// c.JSON(http.StatusOK, gin.H{"token": tokenString})
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true) // 1 days
	c.JSON(http.StatusOK, gin.H{})
}

// Get User By Id
func (db *DBController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := db.Database.First(&user, id)
	// tx := db.Database.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		fmt.Println("Error Get User:", tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	// fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"results": &user})

}

// Get User Lists
func (db *DBController) GetUserLists(c *gin.Context) {
	var user []models.User
	tx := db.Database.Find(&user)
	if tx.Error != nil {
		fmt.Println("Error Get UserLists:", tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"results": &user})
}

// Delete User
func (db *DBController) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	var user []models.User
	tx := db.Database.Delete(&user, id)
	if tx.Error != nil {
		fmt.Println("Error Delete User:", tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user cannot delete"})
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"results": "user deleted"})
}

// Update User
func (db *DBController) UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tx := db.Database.Updates(&user)
	if tx.Error != nil {
		fmt.Println("Error Update User:", tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user cannot update"})
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"results": "user updated"})
}
