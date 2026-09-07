//go:build no_runtime_type_checking

package eks

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCapability_ArgoCdPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsCapability_ArgoCdPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsCapability_ArgoCdPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsCapability_ArgoCdPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsCapability_ArgoCdPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsCapability_ArgoCdPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsCapability_ArgoCdPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsCapability_ArgoCdPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

