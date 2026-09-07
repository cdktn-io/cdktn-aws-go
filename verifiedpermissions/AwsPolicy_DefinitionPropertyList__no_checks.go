//go:build no_runtime_type_checking

package verifiedpermissions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsPolicy_DefinitionPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsPolicy_DefinitionPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsPolicy_DefinitionPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsPolicy_DefinitionPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsPolicy_DefinitionPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsPolicy_DefinitionPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsPolicy_DefinitionPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsPolicy_DefinitionPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

