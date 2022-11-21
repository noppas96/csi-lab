type NodePublishVolumeRequest struct {
	VolumeId 			string 					`protobuf:"bytes,1,opt,name=volume_id,json=volumeId,proto3" 					json:"volume_id,omitempty"`
	PublishContext 		map[string]string 		`protobuf:"bytes,2,rep,name=publish_context,json=publishContext,proto3" 		json:"publish_context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StagingTargetPath 	string 					`protobuf:"bytes,3,opt,name=staging_target_path,json=stagingTargetPath,proto3" 	json:"staging_target_path,omitempty"`
	TargetPath 			string 					`protobuf:"bytes,4,opt,name=target_path,json=targetPath,proto3" 				json:"target_path,omitempty"`
	VolumeCapability 	*VolumeCapability 		`protobuf:"bytes,5,opt,name=volume_capability,json=volumeCapability,proto3" 	json:"volume_capability,omitempty"`