package profile

import (
	"campus-compass/backend/config"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ProfileResponse struct {
	User struct {
		Email        string `json:"email"`
		ProfileImage string `json:"profile_image"`
	} `json:"user"`
	Contribution *struct {
		Review    string  `json:"review"`
		Rating    float64 `json:"rating"`
		ImageURL  string  `json:"image_url"`
	} `json:"contribution,omitempty"`
	Spot *struct {
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
		ImageURL    string  `json:"image_url"`
	} `json:"spot,omitempty"`
}

// func GetProfileHandler(c *gin.Context) {
// 	// Get user ID from JWT token
// 	claims, exists := c.Get("claims")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	userID := claims.(jwt.MapClaims)["sub"].(string)

// 	// Fetch user details
// 	var user struct {
// 		Email        string
// 		ProfileImage string
// 	}
// 	err := config.DB.QueryRow(
// 		"SELECT email, profile_image FROM users WHERE id = $1", 
// 		userID,
// 	).Scan(&user.Email, &user.ProfileImage)

// 	if err != nil {
// 		handleDatabaseError(c, err)
// 		return
// 	}

// 	// Build response
// 	response := buildProfileResponse(userID, user)
// 	c.JSON(http.StatusOK, response)
// }
// Add this to your SetupRoutes function


// Uncomment and keep your original GetProfileHandler
func GetProfileHandler(c *gin.Context) {
	// Get user ID from JWT token
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := claims.(jwt.MapClaims)["sub"].(string)

	// Fetch user details
	var user struct {
		Email        string
		ProfileImage string
	}
	err := config.DB.QueryRow(
		"SELECT email, profile_image FROM users WHERE id = $1", 
		userID,
	).Scan(&user.Email, &user.ProfileImage)

	if err != nil {
		handleDatabaseError(c, err)
		return
	}

	// Build response
	response := buildProfileResponse(userID, user)
	c.JSON(http.StatusOK, response)
}

// Keep your GetProfileByIDHandler implementation
func GetProfileByIDHandler(c *gin.Context) {
    profileID := c.Param("id")
    
    // Verify the requesting user has permission to view this profile
    claims, _ := c.Get("claims")
    requestingUserID := claims.(jwt.MapClaims)["sub"].(string)
    _ = requestingUserID // Use this for auth checks if needed
    
    // Fetch the target user's details
    var user struct {
        Email        string
        ProfileImage string
    }
    err := config.DB.QueryRow(
        "SELECT email, profile_image FROM users WHERE id = $1", 
        profileID,
    ).Scan(&user.Email, &user.ProfileImage)

    if err != nil {
        handleDatabaseError(c, err)
        return
    }

    // Build and return response
    response := buildProfileResponse(profileID, user)
    c.JSON(http.StatusOK, response)
}
func handleDatabaseError(c *gin.Context, err error) {
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	} else {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	}
}

func buildProfileResponse(userID string, user struct {
	Email        string
	ProfileImage string
}) ProfileResponse {
	response := ProfileResponse{}
	response.User.Email = user.Email
	response.User.ProfileImage = user.ProfileImage

	// Fetch and add contribution data if exists
	if contribution := getContribution(userID); contribution != nil {
		response.Contribution = contribution
	}

	// Fetch and add spot data if exists
	if spot := getSpot(userID); spot != nil {
		response.Spot = spot
	}

	return response
}

func getContribution(userID string) *struct {
	Review    string  `json:"review"`
	Rating    float64 `json:"rating"`
	ImageURL  string  `json:"image_url"`
} {
	var contribution struct {
		Review   string
		Rating   float64
		SpotID   string
		ImageURL sql.NullString
	}

	err := config.DB.QueryRow(`
		SELECT c.review, c.rating, c.spot_id, i.url 
		FROM contributions c
		LEFT JOIN image i ON c.spot_id = i.spot_id
		WHERE c.user_id = $1
		LIMIT 1`,
		userID,
	).Scan(&contribution.Review, &contribution.Rating, &contribution.SpotID, &contribution.ImageURL)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Contribution query error: %v", err)
		}
		return nil
	}

	return &struct {
		Review    string  `json:"review"`
		Rating    float64 `json:"rating"`
		ImageURL  string  `json:"image_url"`
	}{
		Review:    contribution.Review,
		Rating:    contribution.Rating,
		ImageURL:  contribution.ImageURL.String,
	}
}

func getSpot(userID string) *struct {
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	ImageURL    string  `json:"image_url"`
} {
	var spot struct {
		Description string
		Rating      float64
		SpotID      string
	}

	err := config.DB.QueryRow(
		"SELECT spot_id, description, rating FROM spots WHERE user_id = $1 LIMIT 1",
		userID,
	).Scan(&spot.SpotID, &spot.Description, &spot.Rating)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Spot query error: %v", err)
		}
		return nil
	}

	var spotImageURL sql.NullString
	_ = config.DB.QueryRow(
		"SELECT url FROM image WHERE spot_id = $1 LIMIT 1",
		spot.SpotID,
	).Scan(&spotImageURL)

	return &struct {
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
		ImageURL    string  `json:"image_url"`
	}{
		Description: spot.Description,
		Rating:      spot.Rating,
		ImageURL:    spotImageURL.String,
	}
}