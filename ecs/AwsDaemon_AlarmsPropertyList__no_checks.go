//go:build no_runtime_type_checking

package ecs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsDaemon_AlarmsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsDaemon_AlarmsPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsDaemon_AlarmsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsDaemon_AlarmsPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsDaemon_AlarmsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsDaemon_AlarmsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsDaemon_AlarmsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsDaemon_AlarmsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

