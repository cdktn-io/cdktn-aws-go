//go:build no_runtime_type_checking

package awsdax

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfCluster_NodesPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfCluster_NodesPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfCluster_NodesPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfCluster_NodesPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfCluster_NodesPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfCluster_NodesPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfCluster_NodesPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

