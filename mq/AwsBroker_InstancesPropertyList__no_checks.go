//go:build no_runtime_type_checking

package mq

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsBroker_InstancesPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsBroker_InstancesPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsBroker_InstancesPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsBroker_InstancesPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsBroker_InstancesPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsBroker_InstancesPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsBroker_InstancesPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

