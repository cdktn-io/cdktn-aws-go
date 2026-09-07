//go:build no_runtime_type_checking

package vpc

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsSecurityGroup_EgressPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsSecurityGroup_EgressPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsSecurityGroup_EgressPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsSecurityGroup_EgressPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsSecurityGroup_EgressPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsSecurityGroup_EgressPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsSecurityGroup_EgressPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsSecurityGroup_EgressPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

