//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfTable_KeySchemaPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfTable_KeySchemaPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfTable_KeySchemaPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfTable_KeySchemaPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfTable_KeySchemaPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfTable_KeySchemaPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfTable_KeySchemaPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

