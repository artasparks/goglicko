// Implementation of the Glicko 2 Rating system, for rating players.  Glicko is
// an improvoment on ELO, but is much more computationally intensive.
//
// For more information, see:
//
// http://www.glicko.net/glicko/glicko2.pdf
//
// http://en.wikipedia.org/wiki/Glicko_rating_system
//
// The calculation process is broken into 8 steps.
//
// Step 1:
// Determine initial values.
//
// Step 2:
// Convert to Glicko2 Scale from the Glicko1 scale.
//
// Step 3:
// Compute (v), the estimated variance based only on game outcomes.
//
// Step 4:
// Compute the quantity Delta, the estimated improvement.
//
// Step 5:
// Determine the new value, sigma', of the volatility, in an iterative process.
//
// Step 6:
// Update the rating deviation to the new pre-rating period value, φ_z
//
// Step 7:
// Update the rating and RD to the new values, μ′ and φ′:
//
// Step 8:
// Convert back to the Glicko1 scale.
package goglicko

import (
	"math"
)

// Overrideable Defaults
var (
	// Constrains the volatility. Typically set between 0.3 and 1.2.  Often
	// refered to as the 'system' constant.
	DefaultTau = 0.3

	DefaultRat = 1500.0 // Default starting rating
	DefaultDev = 350.0  // Default starting deviation
	DefaultVol = 0.06   // Default starting volatility
)

// Miscellaneous Mathematical constants.
const (
	piSq = math.Pi * math.Pi // π^2
	// Constant transformation value, to transform between Glicko 2 and Glicko 1
	glicko2Scale = 173.7178
)

// Used to indicate who won/lost/tied the game.
type Result float64

const (
	Win  Result = 1
	Loss Result = 0
	Draw Result = 0.5
)

////////////////////////////
// Sundry of Helper Funcs //
////////////////////////////

// Ensure that two floats are equal, given some epsilon.
func FloatsMostlyEqual(v1, v2, epsilon float64) bool {
	return math.Abs(v1 - v2) < epsilon
}

// Square function for convenience
func sq(x float64) float64 {
	return x * x
}

// The E function. Written as E(μ,μ_j,φ_j).
// For readability, instead of greek we use the variables
// 	r: rating of player
// 	ri: rating of opponent
// 	devi: deviation of opponent
func ee(r, ri, devi float64) float64 {
	return 1.0 / (1 + math.Exp(-gee(devi)*(r-ri)))
}

// The g function. Written as g(φ).
// For readability, instead of greek we use the variables
// 	dev: The deviation of a player's rating
func gee(dev float64) float64 {
	return 1 / math.Sqrt(1+3*dev*dev/piSq)
}

// Calculate the convergence f(x, a, delta^2, phi^2, tau^2). f(x) in the pdf.
func conv(x, a, estvar, deltaSq, phiSq, tauSq float64) float64 {
	eX := math.Exp(x)
	return eX*(deltaSq-phiSq-estvar-eX)/
		(2*sq(phiSq+estvar+eX)) - (x-a)/tauSq
}
