//go:build no_runtime_type_checking

package ec2imagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsImageRecipe_ParameterPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsImageRecipe_ParameterPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsImageRecipe_ParameterPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsImageRecipe_ParameterPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsImageRecipe_ParameterPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsImageRecipe_ParameterPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsImageRecipe_ParameterPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

