//go:build no_runtime_type_checking

package awss3tables

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfTable_FieldPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfTable_FieldPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfTable_FieldPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfTable_FieldPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfTable_FieldPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfTable_FieldPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfTable_FieldPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfTable_FieldPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

