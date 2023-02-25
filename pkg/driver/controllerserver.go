package driver

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/sirupsen/logrus"
)

var (
	// controllerCaps 代表Controller Plugin支持的功能，可选类型见https://github.com/container-storage-interface/spec/blob/4731db0e0bc53238b93850f43ab05d9355df0fd9/lib/go/csi/csi.pb.go#L181:6
	// 这里只实现Volume的创建/删除，附加/卸载
	controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	}
)

// CreateVolume 创建
func (d *Driver) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (resp *csi.CreateVolumeResponse, finalErr error) {
	logrus.Infof("CreateVolume: called with args %+v", *req)
	vm := &csi.Volume{
		VolumeId:      "w-123",
		CapacityBytes: 20 * (1 << 30),
		VolumeContext: req.GetParameters(),
	}
	return &csi.CreateVolumeResponse{Volume: vm}, nil
}

// DeleteVolume 删除
func (d *Driver) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	logrus.Infof("DeleteVolume: called with args %+v", *req)
	return &csi.DeleteVolumeResponse{}, nil
}

// ControllerPublishVolume 附加
func (d *Driver) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	logrus.Infof("ControllerPublishVolume: called with args %+v", *req)
	pvInfo := map[string]string{"DevicePathKey": "/dev/sdb"}
	return &csi.ControllerPublishVolumeResponse{PublishContext: pvInfo}, nil
}

// ControllerUnpublishVolume 卸载
func (d *Driver) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	logrus.Infof("ControllerUnpublishVolume: called with args %+v", *req)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}
