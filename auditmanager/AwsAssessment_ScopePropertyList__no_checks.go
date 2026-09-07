//go:build no_runtime_type_checking

package auditmanager

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsAssessment_ScopePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsAssessment_ScopePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsAssessment_ScopePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsAssessment_ScopePropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsAssessment_ScopePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsAssessment_ScopePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsAssessment_ScopePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsAssessment_ScopePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

