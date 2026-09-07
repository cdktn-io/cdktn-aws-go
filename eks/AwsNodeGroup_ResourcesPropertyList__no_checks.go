//go:build no_runtime_type_checking

package eks

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsNodeGroup_ResourcesPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsNodeGroup_ResourcesPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsNodeGroup_ResourcesPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsNodeGroup_ResourcesPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsNodeGroup_ResourcesPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsNodeGroup_ResourcesPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsNodeGroup_ResourcesPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

