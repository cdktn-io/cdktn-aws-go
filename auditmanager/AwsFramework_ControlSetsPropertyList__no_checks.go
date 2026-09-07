//go:build no_runtime_type_checking

package auditmanager

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsFramework_ControlSetsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsFramework_ControlSetsPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsFramework_ControlSetsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_ControlSetsPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_ControlSetsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_ControlSetsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsFramework_ControlSetsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsFramework_ControlSetsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

