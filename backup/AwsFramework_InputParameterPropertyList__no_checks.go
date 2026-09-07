//go:build no_runtime_type_checking

package backup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsFramework_InputParameterPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsFramework_InputParameterPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsFramework_InputParameterPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_InputParameterPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_InputParameterPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_InputParameterPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_InputParameterPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsFramework_InputParameterPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

