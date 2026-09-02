//go:build no_runtime_type_checking

package awss3tables

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfTable_MetadataPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfTable_MetadataPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfTable_MetadataPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfTable_MetadataPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfTable_MetadataPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfTable_MetadataPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfTable_MetadataPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfTable_MetadataPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

