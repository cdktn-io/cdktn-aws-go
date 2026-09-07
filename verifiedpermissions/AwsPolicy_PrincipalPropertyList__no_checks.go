//go:build no_runtime_type_checking

package verifiedpermissions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsPolicy_PrincipalPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsPolicy_PrincipalPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsPolicy_PrincipalPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsPolicy_PrincipalPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsPolicy_PrincipalPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsPolicy_PrincipalPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsPolicy_PrincipalPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsPolicy_PrincipalPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

