//go:build no_runtime_type_checking

package awsinspector

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfFilter_ResourceTagsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfFilter_ResourceTagsPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfFilter_ResourceTagsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfFilter_ResourceTagsPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfFilter_ResourceTagsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfFilter_ResourceTagsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfFilter_ResourceTagsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfFilter_ResourceTagsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

