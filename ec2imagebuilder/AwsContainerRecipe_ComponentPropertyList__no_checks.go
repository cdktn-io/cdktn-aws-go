//go:build no_runtime_type_checking

package ec2imagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsContainerRecipe_ComponentPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsContainerRecipe_ComponentPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsContainerRecipe_ComponentPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsContainerRecipe_ComponentPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsContainerRecipe_ComponentPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsContainerRecipe_ComponentPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsContainerRecipe_ComponentPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsContainerRecipe_ComponentPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

