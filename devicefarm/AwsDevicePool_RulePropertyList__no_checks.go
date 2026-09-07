//go:build no_runtime_type_checking

package devicefarm

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsDevicePool_RulePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsDevicePool_RulePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsDevicePool_RulePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsDevicePool_RulePropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsDevicePool_RulePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsDevicePool_RulePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsDevicePool_RulePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsDevicePool_RulePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

