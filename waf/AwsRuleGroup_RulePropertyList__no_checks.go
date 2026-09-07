//go:build no_runtime_type_checking

package waf

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsRuleGroup_RulePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsRuleGroup_RulePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsRuleGroup_RulePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsRuleGroup_RulePropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsRuleGroup_RulePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsRuleGroup_RulePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsRuleGroup_RulePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsRuleGroup_RulePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

