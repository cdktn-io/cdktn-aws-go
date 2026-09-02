//go:build no_runtime_type_checking

package awsmq

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfBroker_UserPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfBroker_UserPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfBroker_UserPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfBroker_UserPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfBroker_UserPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfBroker_UserPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfBroker_UserPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfBroker_UserPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

