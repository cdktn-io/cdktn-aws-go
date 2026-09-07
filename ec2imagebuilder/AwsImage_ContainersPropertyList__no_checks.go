//go:build no_runtime_type_checking

package ec2imagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsImage_ContainersPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsImage_ContainersPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsImage_ContainersPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsImage_ContainersPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsImage_ContainersPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsImage_ContainersPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsImage_ContainersPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

