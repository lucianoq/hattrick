package chpp

// SelfConfidence ...
type SelfConfidence uint

// List of SelfConfidence constants.
const (
	SelfConfidenceNonExistent           SelfConfidence = 0
	SelfConfidenceDisastrous            SelfConfidence = 1
	SelfConfidenceWretched              SelfConfidence = 2
	SelfConfidencePoor                  SelfConfidence = 3
	SelfConfidenceDecent                SelfConfidence = 4
	SelfConfidenceStrong                SelfConfidence = 5
	SelfConfidenceWonderful             SelfConfidence = 6
	SelfConfidenceSlightlyExaggerated   SelfConfidence = 7
	SelfConfidenceExaggerated           SelfConfidence = 8
	SelfConfidenceCompletelyExaggerated SelfConfidence = 9
)
