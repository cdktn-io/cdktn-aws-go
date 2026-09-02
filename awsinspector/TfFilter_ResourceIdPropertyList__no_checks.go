//go:build no_runtime_type_checking

package awsinspector

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfFilter_ResourceIdPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfFilter_ResourceIdPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfFilter_ResourceIdPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfFilter_ResourceIdPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfFilter_ResourceIdPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfFilter_ResourceIdPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfFilter_ResourceIdPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfFilter_ResourceIdPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

