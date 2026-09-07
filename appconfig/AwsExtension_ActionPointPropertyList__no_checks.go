//go:build no_runtime_type_checking

package appconfig

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsExtension_ActionPointPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsExtension_ActionPointPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsExtension_ActionPointPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsExtension_ActionPointPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsExtension_ActionPointPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsExtension_ActionPointPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsExtension_ActionPointPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsExtension_ActionPointPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

