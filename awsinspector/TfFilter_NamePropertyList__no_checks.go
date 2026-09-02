//go:build no_runtime_type_checking

package awsinspector

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfFilter_NamePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfFilter_NamePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfFilter_NamePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfFilter_NamePropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfFilter_NamePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfFilter_NamePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfFilter_NamePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfFilter_NamePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

