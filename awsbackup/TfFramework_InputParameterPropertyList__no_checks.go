//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfFramework_InputParameterPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfFramework_InputParameterPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfFramework_InputParameterPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfFramework_InputParameterPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfFramework_InputParameterPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfFramework_InputParameterPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfFramework_InputParameterPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfFramework_InputParameterPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

