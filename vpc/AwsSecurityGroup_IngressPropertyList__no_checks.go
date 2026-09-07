//go:build no_runtime_type_checking

package vpc

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsSecurityGroup_IngressPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsSecurityGroup_IngressPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsSecurityGroup_IngressPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsSecurityGroup_IngressPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsSecurityGroup_IngressPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsSecurityGroup_IngressPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsSecurityGroup_IngressPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsSecurityGroup_IngressPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

