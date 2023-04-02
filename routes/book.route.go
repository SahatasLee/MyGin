package routes

// route.POST("/book", func(c *gin.Context) {
// 	var book Book
// 	if err := c.ShouldBind(&book); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	tx := db.Create(&book)
// 	if tx.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create book"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"results": "Book created"})
// })

// route.GET("/book", func(c *gin.Context) {
// 	var books []Book
// 	tx := db.Find(&books)
// 	if tx.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"results": books})
// 		return
// 	}

// })

// route.PATCH("/book", func(c *gin.Context) {
// 	var book Book
// 	err := c.ShouldBind(&book)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		return
// 	}

// 	tx := db.Updates(&book)
// 	if tx.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot update book"})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"results": "book updated"})
// 		return
// 	}

// })

// // Delete Book
// route.DELETE("/book/:id", func(c *gin.Context) {
// 	id := c.Param("id")

// 	var book Book
// 	tx := db.Delete(&book, id)
// 	if tx.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete book"})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"results": "book deleted"})
// 		return
// 	}

// })
