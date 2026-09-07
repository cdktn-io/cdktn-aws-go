//go:build no_runtime_type_checking

package keyspaces

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsTable_PartitionKeyPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsTable_PartitionKeyPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsTable_PartitionKeyPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsTable_PartitionKeyPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsTable_PartitionKeyPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsTable_PartitionKeyPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsTable_PartitionKeyPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsTable_PartitionKeyPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

