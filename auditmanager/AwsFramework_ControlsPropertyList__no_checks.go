//go:build no_runtime_type_checking

package auditmanager

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsFramework_ControlsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsFramework_ControlsPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsFramework_ControlsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_ControlsPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_ControlsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_ControlsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_ControlsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsFramework_ControlsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

