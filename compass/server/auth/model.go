// // // Define structs like User, LoginRequest, SignupRequest
// // package auth
// package auth

// import "time"

// type User struct {
// 	ID                string    `json:"id"`
// 	Email             string    `json:"email"`
// 	Password          string    `json:"-"`
// 	ProfileImage      string    `json:"profile_image"`
// 	RegistrationDate  time.Time `json:"registration"`
// 	LastUpdate        time.Time `json:"last_update"`
// 	IsVerified        bool      `json:"is_verified"`
// 	Contributions     int       `json:"contributions_count"`
// }
// type LoginRequest struct {
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required,min=6"`
// }

// type LoginResponse struct {
// 	Success bool   `json:"success"`
// 	Token   string `json:"token"`
// 	UserID  string `json:"user_id"`
// }

// // type SignupRequest struct {
// // 	Email string `json:"email" binding:"required,email"`
// // 	ID    string `json:"id" binding:"required"`
// // 	Photo string `json:"photo,omitempty"`
// // }
package auth

import "time"

type User struct {
    ID                string    `json:"id"`
    Email             string    `json:"email" binding:"required,email"`
    Password          string    `json:"-"` // Hidden from JSON responses
    ProfileImage      string    `json:"profile_image"`
    RegistrationDate  time.Time `json:"registration"`
    LastUpdate        time.Time `json:"last_update"`
    IsVerified        bool      `json:"is_verified"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}