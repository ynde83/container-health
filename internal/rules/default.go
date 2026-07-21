package rules

func Default() []Rule {
	return []Rule{
		HealthcheckRule{},
		RootRule{},
		PrivilegedRule{},
		RestartCountRule{},
		RestartPolicyRule{},
	}
}
