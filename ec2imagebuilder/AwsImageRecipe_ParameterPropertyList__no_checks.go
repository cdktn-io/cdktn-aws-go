//go:build no_runtime_type_checking

package ec2imagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsImageRecipe_ParameterPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsImageRecipe_ParameterPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsImageRecipe_ParameterPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsImageRecipe_ParameterPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsImageRecipe_ParameterPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsImageRecipe_ParameterPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsImageRecipe_ParameterPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsImageRecipe_ParameterPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

