//go:build no_runtime_type_checking

package keyspaces

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsTable_ColumnPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsTable_ColumnPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsTable_ColumnPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ColumnPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ColumnPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ColumnPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ColumnPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsTable_ColumnPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

