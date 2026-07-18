package rules

func Default() []Rule {
	return []Rule{
		HealthcheckRule{},
	}
}
