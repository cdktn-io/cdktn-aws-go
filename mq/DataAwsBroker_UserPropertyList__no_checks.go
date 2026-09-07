//go:build no_runtime_type_checking

package mq

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsBroker_UserPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsBroker_UserPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsBroker_UserPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsBroker_UserPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsBroker_UserPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsBroker_UserPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsBroker_UserPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

