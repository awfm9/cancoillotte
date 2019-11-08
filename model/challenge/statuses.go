// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package challenge

// List of challenge statuses.
const (
	Sent        = "SENT"                // [sent] & [received]
	CommitHome  = "MOVES_SIGNED_HOME"   // [active]
	CommitAway  = "MOVES_SIGNED_AWAY"   // [active]
	CommitBoth  = "MOVES_SIGNED_BOTH"   // [active] "PROCESSING MOVES"
	RevealReady = "REVEAL_READY"        // [active]
	RevealHome  = "MOVES_REVEALED_HOME" // [active]
	RevealAway  = "MOVES_REVEALED_AWAY" // [active]
	RevealBoth  = "MOVES_REVEALED_BOTH" // [active] "FINISHING DUEL"
	Done        = "DUEL_READY"          // [active]
	Rejected    = "REJECTED"            // [terminated]
)
