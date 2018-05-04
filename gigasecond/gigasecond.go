// Calculates when someone will the be 1 billion seconds old
package gigasecond

// import path for the time package from the standard library
import "time"

// This return the input Time value plus 1 billion seconds
func AddGigasecond(t time.Time) time.Time {

	return t.Add(time.Second * 1000000000)
}
