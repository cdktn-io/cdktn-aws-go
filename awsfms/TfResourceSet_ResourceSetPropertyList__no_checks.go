//go:build no_runtime_type_checking

package awsfms

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfResourceSet_ResourceSetPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfResourceSet_ResourceSetPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfResourceSet_ResourceSetPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfResourceSet_ResourceSetPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfResourceSet_ResourceSetPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfResourceSet_ResourceSetPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfResourceSet_ResourceSetPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfResourceSet_ResourceSetPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

