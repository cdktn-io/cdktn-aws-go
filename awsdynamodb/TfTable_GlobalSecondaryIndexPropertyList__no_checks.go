//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfTable_GlobalSecondaryIndexPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

