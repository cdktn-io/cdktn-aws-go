//go:build no_runtime_type_checking

package appconfig

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsEnvironment_MonitorPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsEnvironment_MonitorPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsEnvironment_MonitorPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsEnvironment_MonitorPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsEnvironment_MonitorPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsEnvironment_MonitorPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsEnvironment_MonitorPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsEnvironment_MonitorPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

