package brokenrules

import "bytes"



//IBrokenRules is a collection of rules that have been broken
type IBrokenRules interface {
    HasBrokenRules() bool
    GetBrokenRules() []string
    AddRule(r string) IBrokenRules
    Add(IBrokenRules) IBrokenRules
    Error() string
}


type brokenRules struct {
    rules []string
}

//New is the factory
func New() IBrokenRules {
    var rules []string
    return &brokenRules {
        rules: rules,
    }
}

func (br *brokenRules) HasBrokenRules() bool {
    return br != nil && br.rules != nil && len(br.rules) > 0;
}

func (br *brokenRules) GetBrokenRules() []string {
    return br.rules
}

func (br *brokenRules) AddRule(r string) IBrokenRules{
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

func (br *brokenRules) Error() string {

    rules := br.GetBrokenRules()
    buf := bytes.NewBufferString("")
    for _, rule := range rules {
        buf.WriteString(rule)
    }
    s := buf.String()
    return s
}
