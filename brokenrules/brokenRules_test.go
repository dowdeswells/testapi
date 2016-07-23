package brokenrules

import (
	. "testing"
)

func TestHasBrokenRules(t *T) {

    br := &brokenRules {
        rules: []string{"",""},
    }
    if (br.HasBrokenRules() == false) {
        t.Errorf("Should have 2 rules")
    }

    br = &brokenRules {
        rules: []string{},
    }
    if (br.HasBrokenRules() == true) {
        t.Errorf("Should have 0 rules")
    }

    br = &brokenRules {
        rules: nil,
    }
    if (br.HasBrokenRules() == true) {
        t.Errorf("Should have nil rules")
    }
    br = nil
    if (br.HasBrokenRules() == true) {
        t.Errorf("Should be nil")
    } else {
        t.Log("OK baby")
    }

    var rules []string
    br2 := &brokenRules {
        rules: rules,
    }
    br3 := br2.AddRule("RuleEndTimesInOrder")
    l := len(br3.GetBrokenRules())
    if l != 1 {
        t.Errorf("Should be 1")
    }
}

func TestAddBrokenRules(t *T) {

    br := &brokenRules {
        rules: []string{},
    }

    br2 := &brokenRules {
        rules: []string{"HeyThere"},
    }

    br.Add(br2)

    if (br.HasBrokenRules() == false) {
        t.Errorf("Should have rules")
    }

    rules := br.GetBrokenRules()
    if (len(rules) != 1) {
        t.Errorf("Should have 1 rule")
    }
}