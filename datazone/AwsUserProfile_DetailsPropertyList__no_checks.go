//go:build no_runtime_type_checking

package datazone

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsUserProfile_DetailsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsUserProfile_DetailsPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsUserProfile_DetailsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsUserProfile_DetailsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsUserProfile_DetailsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsUserProfile_DetailsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsUserProfile_DetailsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

