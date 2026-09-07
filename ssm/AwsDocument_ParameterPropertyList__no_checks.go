//go:build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsDocument_ParameterPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsDocument_ParameterPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsDocument_ParameterPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsDocument_ParameterPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsDocument_ParameterPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsDocument_ParameterPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsDocument_ParameterPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

