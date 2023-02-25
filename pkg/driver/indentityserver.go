package driver

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/sirupsen/logrus"
)

// CSI Identity 服务，主要负责对外暴露这个插件本身的信息

// GetPluginInfo 返回插件信息
func (d *Driver) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	logrus.Infof("GetPluginInfo: called with args: %+v", *req)

	return &csi.GetPluginInfoResponse{
		Name:          d.config.DriverName,
		VendorVersion: d.config.VendorVersion,
	}, nil
}

// GetPluginCapabilities 返回插件支持的功能
func (d *Driver) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	logrus.Infof("GetPluginCapabilities: called with args: %+v", *req)
	caps := []*csi.PluginCapability{
		{
			Type: &csi.PluginCapability_Service_{
				Service: &csi.PluginCapability_Service{
					Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
				},
			},
		},
		{
			Type: &csi.PluginCapability_Service_{
				Service: &csi.PluginCapability_Service{
					Type: csi.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS,
				},
			},
		},
	}
	return &csi.GetPluginCapabilitiesResponse{Capabilities: caps}, nil
}

// Probe 插件健康检测
func (d *Driver) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	logrus.Infof("Probe: called with args %+v", req)
	return &csi.ProbeResponse{}, nil
}
