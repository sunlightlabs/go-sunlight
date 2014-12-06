package openstates

// Open States API root URL.
var openstatesRoot string = "http://openstates.org/api/v1"

/**
 * Open States data Source object. Present on nearly all objects in the
 * Open States database.
 */
type Source struct {
	Url string
}
