//go:build no_runtime_type_checking

package bedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsMemory_ResourcePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsMemory_ResourcePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsMemory_ResourcePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsMemory_ResourcePropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsMemory_ResourcePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsMemory_ResourcePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsMemory_ResourcePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsMemory_ResourcePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

