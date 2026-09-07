//go:build no_runtime_type_checking

package dynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsTable_ReplicaPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsTable_ReplicaPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsTable_ReplicaPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ReplicaPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ReplicaPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ReplicaPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ReplicaPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsTable_ReplicaPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

