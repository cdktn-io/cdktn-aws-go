//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfTable_AttributePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfTable_AttributePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfTable_AttributePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfTable_AttributePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfTable_AttributePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfTable_AttributePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfTable_AttributePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

