//go:build no_runtime_type_checking

package datazone

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsUserProfile_SsoPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsUserProfile_SsoPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsUserProfile_SsoPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsUserProfile_SsoPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsUserProfile_SsoPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsUserProfile_SsoPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsUserProfile_SsoPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

