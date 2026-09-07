//go:build no_runtime_type_checking

package mskconnect

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsConnector_PluginPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsConnector_PluginPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsConnector_PluginPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsConnector_PluginPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsConnector_PluginPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsConnector_PluginPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsConnector_PluginPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsConnector_PluginPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

