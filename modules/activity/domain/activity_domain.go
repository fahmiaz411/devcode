package domain

import "time"

const (
	Model = "Activity"
)

type Activity struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// Create

type ActivityCreateRequest struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type ActivityCreateResponse struct {
	Activity
}

// Update

type ActivityUpdateRequest struct {
	ID int64 `json:"-"`
	Title string `json:"title"`
	UpdatedAt time.Time `json:"-"`
}

type ActivityUpdateResponse struct {
	Activity
}

// Delete

type ActivityDeleteRequest struct {
	ID int64
}

type ActivityDeleteResponse struct {
}

// Get All

type ActivityGetAllRequest struct {
}

type ActivityGetAllResponse []Activity

// Get One

type ActivityGetOneRequest struct {
	ID int64
}

type ActivityGetOneResponse struct {
	Activity
}