# Changelog 1.6

## 1.6.0 - 2022-03-29

- The Egress feature is graduated from Alpha to Beta and is therefore enabled by default.
- The support for proxying all Service traffic by Antrea Proxy (enabled by `antreaProxy.proxyAll`) is now Beta.

### Added

- Add the following capabilities to the [Antrea IPAM] feature:
  * Support pre-allocating continuous IPs for StatefulSet. ([#3281](https://github.com/antrea-io/antrea/pull/3281), [@annakhm])
  * Support specifying VLAN for IPPool. Traffic from Pods whose IPPools are configured with a VLAN ID will be tagged when leaving the Node uplink. ([#3247](https://github.com/antrea-io/antrea/pull/3247), [@gran-vmv])
- Add the following capabilities to the [Antrea Multi-cluster] feature:
  * Support Antrea ClusterNetworkPolicy replication in a ClusterSet. ([#3363](https://github.com/antrea-io/antrea/pull/3363), [@Dyanngg])
  * Add `antctl mc` subcommand for Antrea Multi-cluster resources. ([#3341](https://github.com/antrea-io/antrea/pull/3341) [#3287](https://github.com/antrea-io/antrea/pull/3287), [@hjiajing] [@bangqipropel])
- Add the following capabilities to the [AntreaPolicy] feature:
  * Add Node selector in Antrea-native policies to allow matching traffic originating from specific Nodes or destined to specific Nodes. ([#3038](https://github.com/antrea-io/antrea/pull/3038), [@wenqiq])
  * Add ServiceAccount selector in Antrea-native policies to allow selecting Pods by ServiceAccount. ([#3044](https://github.com/antrea-io/antrea/pull/3044), [@GraysonWu])
  * Support Pagination for ClusterGroupMembership API. ([#3183](https://github.com/antrea-io/antrea/pull/3183), [@qiyueyao])
  * Add Port Number to Audit Logging. ([#3277](https://github.com/antrea-io/antrea/pull/3277), [@qiyueyao])
- [Flow Visibility] Add Grafana Flow Collector as the new visualization tool for flow records.
  * Add Grafana dashboards, Clickhouse data schema, deployment files, and doc. ([#3063](https://github.com/antrea-io/antrea/pull/3063) [#3525](https://github.com/antrea-io/antrea/pull/3525), [@heanlan] [@zyiou] [@dreamtalen])
  * Add support for exporting flow records to ClickHouse from Flow Aggregator. ([#3196](https://github.com/antrea-io/antrea/pull/3196) [#3526](https://github.com/antrea-io/antrea/pull/3526), [@wsquan171] [@dreamtalen])
  * Add ClickHouse monitor to ensure data retention for in-memory ClickHouse deployment. ([#3244](https://github.com/antrea-io/antrea/pull/3244) [#3498](https://github.com/antrea-io/antrea/pull/3498), [@yanjunz97])
- [Multicast] Support IGMPv3 leave action. ([#3389](https://github.com/antrea-io/antrea/pull/3389), [@wenyingd])
- [Windows] Add support for EndpointSlices on Windows Nodes. ([#3321](https://github.com/antrea-io/antrea/pull/3321), [@XinShuYang])
- Add SKIP_CNI_BINARIES environment variable to support skipping the installation of specified CNI plugins. ([#3454](https://github.com/antrea-io/antrea/pull/3454), [@jainpulkit22])
- Support UBI8-based container image to run Antrea. ([#3273](https://github.com/antrea-io/antrea/pull/3273), [@ksamoray])
- Add the following documentations:
  * Add [documentation](https://github.com/antrea-io/antrea/blob/v1.6.0/docs/service-loadbalancer.md) for ServiceExternalIP feature and Service of type LoadBalancer. ([#3322](https://github.com/antrea-io/antrea/pull/3322), [@hty690])
  * Add [documentation](https://github.com/antrea-io/antrea/blob/v1.6.0/docs/minikube.md) for deploying Antrea to Minikube cluster. ([#3391](https://github.com/antrea-io/antrea/pull/3391), [@jainpulkit22])
  * Add [documentation](https://github.com/antrea-io/antrea/blob/v1.6.0/docs/multicluster/antctl.md) for `antctl` Multi-cluster commands. ([#3414](https://github.com/antrea-io/antrea/pull/3414), [@bangqipropel])
  * Add [documentation](https://github.com/antrea-io/antrea/blob/v1.6.0/docs/antrea-ipam.md) for Multiple-VLAN support. ([#3507](https://github.com/antrea-io/antrea/pull/3507), [@gran-vmv])
  * Add [upgrade guide](https://github.com/antrea-io/antrea/blob/v1.6.0/docs/multicluster/upgrade.md) for Multi-cluster. ([#3374](https://github.com/antrea-io/antrea/pull/3374), [@luolanzone])
  * Add [a per-rule example](https://github.com/antrea-io/antrea/blob/v1.6.0/docs/feature-gates.md#networkpolicystats) for NetworkpolicyStats docs. ([#3356](https://github.com/antrea-io/antrea/pull/3356), [@ceclinux])

### Changed

- Remove all legacy (*.antrea.tanzu.vmware.com) APIs. ([#3299](https://github.com/antrea-io/antrea/pull/3299), [@antoninbas])
- Remove Kind-specific manifest and scripts. Antrea now uses OVS kernel datapath for Kind clusters. ([#3413](https://github.com/antrea-io/antrea/pull/3413), [@antoninbas])
- [Windows] Use uplink MAC as source MAC when transmitting packets to underlay network from Windows Nodes. Therefore, MAC address spoofing configuration like "Forged transmits" in VMware vSphere doesn't need to be enabled. ([#3516](https://github.com/antrea-io/antrea/pull/3516), [@wenyingd])
- Add an agent config parameter "enableBridgingMode" for enabling flexible IPAM (bridging mode). ([#3297](https://github.com/antrea-io/antrea/pull/3297) [#3365](https://github.com/antrea-io/antrea/pull/3365), [@jianjuns])
- Use iptables-wrapper in Antrea container to support distros that runs iptables in "nft" mode. ([#3276](https://github.com/antrea-io/antrea/pull/3276), [@antoninbas])
- Install CNI configuration files after installing CNI binaries to support container runtime `cri-o`. ([#3154](https://github.com/antrea-io/antrea/pull/3154), [@tnqn])
- Upgrade packaged Whereabouts version to v0.5.1. ([#3511](https://github.com/antrea-io/antrea/pull/3511), [@antoninbas])
- Upgrade to go-ipfix v0.5.12. ([#3352](https://github.com/antrea-io/antrea/pull/3352), [@yanjunz97])
- Upgrade Kustomize from v3.8.8 to v4.4.1 to fix Cronjob patching bugs. ([#3402](https://github.com/antrea-io/antrea/pull/3402), [@yanjunz97])
- Fail in Agent initialization if GRE tunnel type is used with IPv6. ([#3156](https://github.com/antrea-io/antrea/pull/3156), [@antoninbas])
- Refactor the OpenFlow pipeline for future extensibility. ([#3058](https://github.com/antrea-io/antrea/pull/3058), [@hongliangl])
- Validate IP ranges of IPPool for Antrea IPAM. ([#2995](https://github.com/antrea-io/antrea/pull/2995), [@ksamoray])
- Validate protocol in the CRD schema of Antrea-native policies. ([#3342](https://github.com/antrea-io/antrea/pull/3342), [@KMAnju-2021])
- Validate labels in the CRD schema of Antrea-native policies and ClusterGroup. ([#3331](https://github.com/antrea-io/antrea/pull/3331), [@GraysonWu])
- Reduce permissions of Antrea ServiceAccounts. ([#3393](https://github.com/antrea-io/antrea/pull/3393), [@tnqn])
- Remove --k8s-1.15 flag from hack/generate-manifest.sh. ([#3350](https://github.com/antrea-io/antrea/pull/3350), [@antoninbas])
- Remove unnecessary CRDs and RBAC rules from Multi-cluster manifest. ([#3491](https://github.com/antrea-io/antrea/pull/3491), [@luolanzone])
- Update label and image repo of antrea-mc-controller to be consistent with antrea-controller and antrea-agent. ([#3266](https://github.com/antrea-io/antrea/pull/3266) [#3466](https://github.com/antrea-io/antrea/pull/3466), [@luolanzone])
- Add clusterID annotation to ServiceExport/Import resources. ([#3359](https://github.com/antrea-io/antrea/pull/3359), [@luolanzone])
- Do not log error when Service for Endpoints is not found to avoid log spam. ([#3256](https://github.com/antrea-io/antrea/pull/3256), [@tnqn])
- Ignore Services of type ExternalName for NodePortLocal feature. ([#3114](https://github.com/antrea-io/antrea/pull/3114), [@antoninbas])
- Add powershell command replacement in the Antrea Windows documentation. ([#3264](https://github.com/antrea-io/antrea/pull/3264), [@GraysonWu])

### Fixed

- Add userspace ARP/NDP responders to fix Egress and ServiceExternalIP support for IPv6 clusters. ([#3318](https://github.com/antrea-io/antrea/pull/3318), [@hty690])
- Fix incorrect results by `antctl get networkpolicy` when both Pod and Namespace are specified. ([#3499](https://github.com/antrea-io/antrea/pull/3499), [@Dyanngg])
- Fix IP leak issue when AntreaIPAM is enabled. ([#3314](https://github.com/antrea-io/antrea/pull/3314), [@gran-vmv])
- Fix error when dumping OVS flows for a NetworkPolicy via `antctl get ovsflows`. ([#3335](https://github.com/antrea-io/antrea/pull/3335), [@jainpulkit22])
- Fix IPsec encryption for IPv6 overlays. ([#3155](https://github.com/antrea-io/antrea/pull/3155), [@antoninbas])
- Add ignored interfaces names when getting interface by IP to fix NetworkPolicyOnly mode in AKE. ([#3219](https://github.com/antrea-io/antrea/pull/3219), [@wenyingd])
- Fix duplicate IP case for NetworkPolicy. ([#3467](https://github.com/antrea-io/antrea/pull/3467), [@tnqn])
- Don't delete the routes which are added for the peer IPv6 gateways on Agent startup. ([#3336](https://github.com/antrea-io/antrea/pull/3336) [#3490](https://github.com/antrea-io/antrea/pull/3490), [@Jexf] [@xliuxu])
- Fix pkt mark conflict between HostLocalSourceMark and SNATIPMark. ([#3430](https://github.com/antrea-io/antrea/pull/3430), [@tnqn])
- Unconditionally sync CA cert for Controller webhooks to fix Egress support when AntreaPolicy is disabled. ([#3421](https://github.com/antrea-io/antrea/pull/3421), [@antoninbas])
- Fix inability to access NodePort in particular cases. ([#3371](https://github.com/antrea-io/antrea/pull/3371), [@hongliangl])
- Fix ipBlocks referenced in nested ClusterGroup not processed correctly. ([#3383](https://github.com/antrea-io/antrea/pull/3383), [@Dyanngg])
- Realize Egress for a Pod as soon as its network is created. ([#3360](https://github.com/antrea-io/antrea/pull/3360), [@tnqn])
- Fix NodePort/LoadBalancer issue when proxyAll is enabled. ([#3295](https://github.com/antrea-io/antrea/pull/3295), [@hongliangl])
- Do not panic when processing a PacketIn message for a denied connection. ([#3447](https://github.com/antrea-io/antrea/pull/3447), [@antoninbas])
- Fix CT mark matching without range in flow exporter. ([#3348](https://github.com/antrea-io/antrea/pull/3348), [@hongliangl])
- [Windows] Enable IP forwarding of the Windows bridge local interface to fix support for Service of type LoadBalancer. ([#3137](https://github.com/antrea-io/antrea/pull/3137), [@hongliangl])

[Antrea Multi-cluster]: https://github.com/antrea-io/antrea/blob/v1.6.0/docs/multicluster/user-guide.md
[Antrea IPAM]: https://github.com/antrea-io/antrea/blob/v1.6.0/docs/antrea-ipam.md
[AntreaPolicy]: https://github.com/antrea-io/antrea/blob/v1.6.0/docs/antrea-network-policy.md

[@Atish-iaf]: https://github.com/Atish-iaf
[@Dyanngg]: https://github.com/Dyanngg
[@GraysonWu]: https://github.com/GraysonWu
[@Jexf]: https://github.com/Jexf
[@KMAnju-2021]: https://github.com/KMAnju-2021
[@XinShuYang]: https://github.com/XinShuYang
[@annakhm]: https://github.com/annakhm
[@antoninbas]: https://github.com/antoninbas
[@antrea-bot]: https://github.com/antrea-bot
[@bangqipropel]: https://github.com/bangqipropel
[@ceclinux]: https://github.com/ceclinux
[@dependabot]: https://github.com/dependabot
[@dreamtalen]: https://github.com/dreamtalen
[@gran-vmv]: https://github.com/gran-vmv
[@heanlan]: https://github.com/heanlan
[@hjiajing]: https://github.com/hjiajing
[@hongliangl]: https://github.com/hongliangl
[@hty690]: https://github.com/hty690
[@jainpulkit22]: https://github.com/jainpulkit22
[@jianjuns]: https://github.com/jianjuns
[@ksamoray]: https://github.com/ksamoray
[@luolanzone]: https://github.com/luolanzone
[@qiyueyao]: https://github.com/qiyueyao
[@tnqn]: https://github.com/tnqn
[@wenqiq]: https://github.com/wenqiq
[@wenyingd]: https://github.com/wenyingd
[@wsquan171]: https://github.com/wsquan171
[@xliuxu]: https://github.com/xliuxu
[@yanjunz97]: https://github.com/yanjunz97
[@zyiou]: https://github.com/zyiou
