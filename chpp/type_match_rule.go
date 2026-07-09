package chpp

// MatchRule is the rule set in place for a ladder or tournament match
// (e.g. an age restriction), defaulting to 0 = no rules for other match
// types.
type MatchRule uint

// List of MatchRule constants.
const (
	// MatchRuleNoRules ...
	MatchRuleNoRules MatchRule = 0
	// MatchRuleHomegrown ...
	MatchRuleHomegrown MatchRule = 1
	// MatchRuleUnderEqual20 ...
	MatchRuleUnderEqual20 MatchRule = 2
	// MatchRuleOver33 ...
	MatchRuleOver33 MatchRule = 3
)
