package goglicko

// Constant transformation value, to transform between Glicko 2 and Glicko 1
const scale = 173.7178

// Represents a player's rating and the confidence in a player's rating.
type Rating struct {
	Rating     float64 // Player's rating. Usually starts off at 1500.
	Deviation  float64 // Confidence in a player's rating
	Volatility float64 // Measures erratic performances
}

// Creates a default Rating using:
// 	Rating     = 1500
// 	Deviation  = 350
// 	Volatility = 0.06
func DefaultRating() *Rating {
	return &Rating{
		1500,
		350,
		0.06,
	}
}

// Creates a new custom Rating.
func NewRating(r, rd, s float64) *Rating {
	return &Rating{r, rd, s}
}

// Creates a new rating, converted from Glicko1 scaling to Glicko2 scaling.
// This assumes the starting rating value is 1500.
func (r *Rating) toGlicko2() *Rating {
	return NewRating(
		(r.Rating-1500)/scale,
		(r.Deviation)/scale,
		r.Volatility)
}

// Creates a new rating, converted from Glicko2 scaling to Glicko1 scaling.
// This assumes the starting rating value is 1500.
func (r *Rating) fromGlicko2() *Rating {
	return NewRating(
		r.Rating*scale+1500,
		r.Deviation*scale,
		r.Volatility)
}

type Result float64

const Win Result = 1
const Loss Result = 0
const Draw Result = 0.5
