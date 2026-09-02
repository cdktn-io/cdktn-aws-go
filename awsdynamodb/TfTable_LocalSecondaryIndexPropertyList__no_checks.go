//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfTable_LocalSecondaryIndexPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfTable_LocalSecondaryIndexPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfTable_LocalSecondaryIndexPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfTable_LocalSecondaryIndexPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TfTable_LocalSecondaryIndexPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfTable_LocalSecondaryIndexPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfTable_LocalSecondaryIndexPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfTable_LocalSecondaryIndexPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

