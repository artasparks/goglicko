// Implementation of the Glicko 2 Rating system.
//
// For more information, see:
//
// http://www.glicko.net/glicko/glicko2.pdf
//
// http://en.wikipedia.org/wiki/Glicko_rating_system
//
// The rating is composed of three parts:
// 	Rating (r):         The rating
// 	Deviation (RD):     Measures rating uncertainty
// 	Volatility (sigma): Measures erratic performances
//
// In addition, we define a system constant (tau), which constrains the change
// in volatility over time.  Reasonable values are between 0.3 and 0.12. Smaller
// values prevent the volatility measures from changing by large amounts.
//
// It's usually more informative to measure the rating as a confidence interval,
// which is given as (Rating - [2*Deviation], Rating + [2*Deviation]).  The
// volatility is not involved in this calculation.
//
// The calculation process is broken into 8 steps.
//
// Determine initial values.  For new players, they will be:
//
// 	Rating (r):            1500
// 	Deviation (RD):        350
// 	Volatility (sigma):   0.06
// 	System Constant (tau): 0.3
//
// Otherwise, we use a player's Rating, Deviation, and Volatility.
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
// Step 5
// Determine the new value, sigma', of the volatility, in an iterative process.
//
// Step 6
// Update the rating deviation to the new pre-rating period value, φ_zed
//
// Step 7
// Update the rating and RD to the new values, μ′ and φ′:
package goglicko
