// Code generated by "stringer -type=Resource"; DO NOT EDIT.

package images

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CalicoCNI-1]
	_ = x[CalicoController-2]
	_ = x[CalicoNode-3]
	_ = x[Flannel-4]
	_ = x[Cilium-5]
	_ = x[CiliumOperator-6]
	_ = x[HubbleRelay-7]
	_ = x[HubbleUI-8]
	_ = x[HubbleUIBackend-9]
	_ = x[HubbleProxy-10]
	_ = x[CiliumCertGen-11]
	_ = x[WeaveNetCNIKube-12]
	_ = x[WeaveNetCNINPC-13]
	_ = x[DNSNodeCache-14]
	_ = x[MachineController-15]
	_ = x[MetricsServer-16]
	_ = x[OperatingSystemManager-17]
	_ = x[ClusterAutoscaler-18]
	_ = x[CSIAttacher-19]
	_ = x[CSINodeDriverRegistar-20]
	_ = x[CSIProvisioner-21]
	_ = x[CSISnapshotter-22]
	_ = x[CSIResizer-23]
	_ = x[CSILivenessProbe-24]
	_ = x[AwsCCM-25]
	_ = x[AzureCCM-26]
	_ = x[AzureCNM-27]
	_ = x[AwsEbsCSI-28]
	_ = x[AwsEbsCSIAttacher-29]
	_ = x[AwsEbsCSILivenessProbe-30]
	_ = x[AwsEbsCSINodeDriverRegistrar-31]
	_ = x[AwsEbsCSIProvisioner-32]
	_ = x[AwsEbsCSIResizer-33]
	_ = x[AwsEbsCSISnapshotter-34]
	_ = x[AwsEbsCSISnapshotController-35]
	_ = x[AzureFileCSI-36]
	_ = x[AzureFileCSIAttacher-37]
	_ = x[AzureFileCSILivenessProbe-38]
	_ = x[AzureFileCSINodeDriverRegistar-39]
	_ = x[AzureFileCSIProvisioner-40]
	_ = x[AzureFileCSIResizer-41]
	_ = x[AzureFileCSISnapshotter-42]
	_ = x[AzureFileCSISnapshotterController-43]
	_ = x[AzureDiskCSI-44]
	_ = x[AzureDiskCSIAttacher-45]
	_ = x[AzureDiskCSILivenessProbe-46]
	_ = x[AzureDiskCSINodeDriverRegistar-47]
	_ = x[AzureDiskCSIProvisioner-48]
	_ = x[AzureDiskCSIResizer-49]
	_ = x[AzureDiskCSISnapshotter-50]
	_ = x[AzureDiskCSISnapshotterController-51]
	_ = x[NutanixCSILivenessProbe-52]
	_ = x[NutanixCSI-53]
	_ = x[NutanixCSIProvisioner-54]
	_ = x[NutanixCSIRegistrar-55]
	_ = x[NutanixCSIResizer-56]
	_ = x[NutanixCSISnapshotter-57]
	_ = x[NutanixCSISnapshotController-58]
	_ = x[NutanixCSISnapshotValidationWebhook-59]
	_ = x[DigitalOceanCSI-60]
	_ = x[DigitalOceanCSIAlpine-61]
	_ = x[DigitalOceanCSIAttacher-62]
	_ = x[DigitalOceanCSINodeDriverRegistar-63]
	_ = x[DigitalOceanCSIProvisioner-64]
	_ = x[DigitalOceanCSIResizer-65]
	_ = x[DigitalOceanCSISnapshotController-66]
	_ = x[DigitalOceanCSISnapshotValidationWebhook-67]
	_ = x[DigitalOceanCSISnapshotter-68]
	_ = x[OpenstackCSI-69]
	_ = x[OpenstackCSINodeDriverRegistar-70]
	_ = x[OpenstackCSILivenessProbe-71]
	_ = x[OpenstackCSIAttacher-72]
	_ = x[OpenstackCSIProvisioner-73]
	_ = x[OpenstackCSIResizer-74]
	_ = x[OpenstackCSISnapshotter-75]
	_ = x[DigitaloceanCCM-76]
	_ = x[HetznerCCM-77]
	_ = x[HetznerCSI-78]
	_ = x[OpenstackCCM-79]
	_ = x[EquinixMetalCCM-80]
	_ = x[VsphereCCM-81]
	_ = x[VsphereCSIDriver-82]
	_ = x[VsphereCSISyncer-83]
	_ = x[VsphereCSIProvisioner-84]
	_ = x[CalicoVXLANCNI-85]
	_ = x[CalicoVXLANController-86]
	_ = x[CalicoVXLANNode-87]
}

const _Resource_name = "CalicoCNICalicoControllerCalicoNodeFlannelCiliumCiliumOperatorHubbleRelayHubbleUIHubbleUIBackendHubbleProxyCiliumCertGenWeaveNetCNIKubeWeaveNetCNINPCDNSNodeCacheMachineControllerMetricsServerOperatingSystemManagerClusterAutoscalerCSIAttacherCSINodeDriverRegistarCSIProvisionerCSISnapshotterCSIResizerCSILivenessProbeAwsCCMAzureCCMAzureCNMAwsEbsCSIAwsEbsCSIAttacherAwsEbsCSILivenessProbeAwsEbsCSINodeDriverRegistrarAwsEbsCSIProvisionerAwsEbsCSIResizerAwsEbsCSISnapshotterAwsEbsCSISnapshotControllerAzureFileCSIAzureFileCSIAttacherAzureFileCSILivenessProbeAzureFileCSINodeDriverRegistarAzureFileCSIProvisionerAzureFileCSIResizerAzureFileCSISnapshotterAzureFileCSISnapshotterControllerAzureDiskCSIAzureDiskCSIAttacherAzureDiskCSILivenessProbeAzureDiskCSINodeDriverRegistarAzureDiskCSIProvisionerAzureDiskCSIResizerAzureDiskCSISnapshotterAzureDiskCSISnapshotterControllerNutanixCSILivenessProbeNutanixCSINutanixCSIProvisionerNutanixCSIRegistrarNutanixCSIResizerNutanixCSISnapshotterNutanixCSISnapshotControllerNutanixCSISnapshotValidationWebhookDigitalOceanCSIDigitalOceanCSIAlpineDigitalOceanCSIAttacherDigitalOceanCSINodeDriverRegistarDigitalOceanCSIProvisionerDigitalOceanCSIResizerDigitalOceanCSISnapshotControllerDigitalOceanCSISnapshotValidationWebhookDigitalOceanCSISnapshotterOpenstackCSIOpenstackCSINodeDriverRegistarOpenstackCSILivenessProbeOpenstackCSIAttacherOpenstackCSIProvisionerOpenstackCSIResizerOpenstackCSISnapshotterDigitaloceanCCMHetznerCCMHetznerCSIOpenstackCCMEquinixMetalCCMVsphereCCMVsphereCSIDriverVsphereCSISyncerVsphereCSIProvisionerCalicoVXLANCNICalicoVXLANControllerCalicoVXLANNode"

var _Resource_index = [...]uint16{0, 9, 25, 35, 42, 48, 62, 73, 81, 96, 107, 120, 135, 149, 161, 178, 191, 213, 230, 241, 262, 276, 290, 300, 316, 322, 330, 338, 347, 364, 386, 414, 434, 450, 470, 497, 509, 529, 554, 584, 607, 626, 649, 682, 694, 714, 739, 769, 792, 811, 834, 867, 890, 900, 921, 940, 957, 978, 1006, 1041, 1056, 1077, 1100, 1133, 1159, 1181, 1214, 1254, 1280, 1292, 1322, 1347, 1367, 1390, 1409, 1432, 1447, 1457, 1467, 1479, 1494, 1504, 1520, 1536, 1557, 1571, 1592, 1607}

func (i Resource) String() string {
	i -= 1
	if i < 0 || i >= Resource(len(_Resource_index)-1) {
		return "Resource(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Resource_name[_Resource_index[i]:_Resource_index[i+1]]
}
