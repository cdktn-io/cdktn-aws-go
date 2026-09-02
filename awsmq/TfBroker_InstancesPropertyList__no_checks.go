//go:build no_runtime_type_checking

package awsmq

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfBroker_InstancesPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfBroker_InstancesPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfBroker_InstancesPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfBroker_InstancesPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfBroker_InstancesPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfBroker_InstancesPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfBroker_InstancesPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

