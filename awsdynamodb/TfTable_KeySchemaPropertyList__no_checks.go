//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfTable_KeySchemaPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfTable_KeySchemaPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfTable_KeySchemaPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfTable_KeySchemaPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfTable_KeySchemaPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfTable_KeySchemaPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfTable_KeySchemaPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfTable_KeySchemaPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

