//go:build no_runtime_type_checking

package keyspaces

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsTable_ClusteringKeyPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsTable_ClusteringKeyPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsTable_ClusteringKeyPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ClusteringKeyPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ClusteringKeyPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ClusteringKeyPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsTable_ClusteringKeyPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsTable_ClusteringKeyPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

