package manager

type CustomerPlan struct {
	// Duration is the duration of the plan in seconds, after which it needs to be paid
	// or will be deleted
	Duration int64 `json:"duration"`
	// StorageSize is size for the storage in bytes of the docker for i2i
	StorageSize int64 `json:"storage_size"`
	// unique plan ID
	ID string `json:"id"`
	// The plan’s name, meant to be displayable to the customer.
	Name string `json:"name"`
	// The plan’s description, meant to be displayable to the customer.
	Description string `json:"description"`
	// ISO currency description
	Currency string `json:"currency"`
	// price, note: for 10 USD, need to use 1000
	Price int64 `json:"price"`
}
