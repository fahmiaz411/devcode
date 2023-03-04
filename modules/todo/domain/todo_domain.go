package domain

import "time"

// Priority
const (
	PriorityVeryHigh = "very-high"
	PriorityHigh = "high"
	PriorityMedium = "medium"
	PriorityLow = "low"
	PriorityVeryLow = "very-low"

	PriorityDefault = PriorityVeryHigh
)

var (
	PriorityAllList = []string{
		PriorityVeryHigh,
		PriorityHigh,
		PriorityMedium,
		PriorityLow,
		PriorityVeryLow,
	}
)

// Is Active
const (
	IsActiveDefault = true
)

const (
	Model = "Todo"
)

type Todo struct {
	ID        		int64     `json:"id"`
	ActivityGroupID int64 	  `json:"activity_group_id"`
	Title     		string    `json:"title"`
	IsActive		bool	  `json:"is_active"`
	Priority		string	  `json:"priority"`
	CreatedAt 		time.Time `json:"createdAt"`
	UpdatedAt 		time.Time `json:"updatedAt"`
}

// Create

type TodoCreateRequest struct {
	Title 	 		string `json:"title"`
	ActivityGroupID int64  `json:"activity_group_id"`
	IsActive		*bool	  `json:"is_active"`
}

type TodoCreateResponse struct {
	Todo
}

// Update

type TodoUpdateRequest struct {
	ID 				int64 	`json:"-"`
	Title 			string 	`json:"title"`
	IsActive		*bool	`json:"is_active"`
	Priority		string	`json:"priority"`
	UpdatedAt time.Time 	`json:"-"`
}

type TodoUpdateResponse struct {
	Todo
}

// Delete

type TodoDeleteRequest struct {
	ID int64
}

type TodoDeleteResponse struct {
}

// Get All

type TodoGetAllRequest struct {
	ActivityGroupID int64 	`json:"activity_group_id"`
}

type TodoGetAllResponse []Todo

// Get One

type TodoGetOneRequest struct {
	ID int64
}

type TodoGetOneResponse struct {
	Todo
}