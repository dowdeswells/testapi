package brokenrules

//import "log"


//RuleID is the broken rule identified by a well known code
type RuleID string

//IBrokenRules is a collection of rules that have been broken
type IBrokenRules interface {
    HasBrokenRules() bool
    GetBrokenRules() []RuleID
    AddRule(r RuleID) IBrokenRules
    Add(IBrokenRules) IBrokenRules
}


type brokenRules struct {
    rules []RuleID
}

//New is the factory
func New() IBrokenRules {
    var rules []RuleID
    return &brokenRules {
        rules: rules,
    }
}

func (br *brokenRules) HasBrokenRules() bool {
    return br != nil && br.rules != nil && len(br.rules) > 0;
}

func (br *brokenRules) GetBrokenRules() []RuleID {
    return br.rules
}

func (br *brokenRules) AddRule(r RuleID) IBrokenRules{
    br.rules = append(br.rules, r)
    return br
}

func (br *brokenRules) Add(abr IBrokenRules) IBrokenRules {

    if (abr == nil || len(abr.GetBrokenRules()) == 0 ) {
        return br
    }

    br.rules = append(br.rules, abr.GetBrokenRules()...)
    return br
}
