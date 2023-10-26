# Changelog v1.54

## [MALFORMED]


 - #5094 unknown section "dhctl, modules"

## Know before update


 - Check that the alerts come after the update
 - Check that the prometheus metrics come after the update
 - check the pods that are behind the kube-rbac-proxy

## Features


 - **[admission-policy-engine]** Add java vulnerability scan capability to trivy-provider. [#6139](https://github.com/deckhouse/deckhouse/pull/6139)
    trivy-provider will restar.
 - **[chrony]** Chrony image is based on distroless. [#6240](https://github.com/deckhouse/deckhouse/pull/6240)
 - **[deckhouse]** Change deckhouse-controller user to deckhouse [#5841](https://github.com/deckhouse/deckhouse/pull/5841)
 - **[dhctl]** Implemented copying of Deckhouse images to third-party registries for air-gapped installation. [#6257](https://github.com/deckhouse/deckhouse/pull/6257)
 - **[external-module-manager]** Add support for module pull from insecure (http) registry. [#6340](https://github.com/deckhouse/deckhouse/pull/6340)
 - **[ingress-nginx]** Add v1.9.3 Ingress Nginx controller version. [#6312](https://github.com/deckhouse/deckhouse/pull/6312)
    In case of switching to '1.9' controller version, relevant Ingress nginx's pods will be recreated.
 - **[local-path-provisioner]** Move to distroless [#6194](https://github.com/deckhouse/deckhouse/pull/6194)
 - **[monitoring-ping]** Switch monitoring-ping to use distroless image and static build a python [#6204](https://github.com/deckhouse/deckhouse/pull/6204)
 - **[prometheus]** Ability to set custom logo for the Grafana dashboard [#6268](https://github.com/deckhouse/deckhouse/pull/6268)
 - **[user-authn]** Ability to set custom logo for the Dex login page [#6268](https://github.com/deckhouse/deckhouse/pull/6268)

## Fixes


 - **[candi]** fix big time drift on nodes [#6297](https://github.com/deckhouse/deckhouse/pull/6297)
    all chrony pods should be restarted
 - **[common]** fix cve issues in kube-rbac-proxy image [#6316](https://github.com/deckhouse/deckhouse/pull/6316)
    check the pods that are behind the kube-rbac-proxy
 - **[external-module-manager]** add support for hardlinks and symlinks to the module [#6330](https://github.com/deckhouse/deckhouse/pull/6330)
 - **[ingress-nginx]** fix cve issues in protobuf-exporter image [#6327](https://github.com/deckhouse/deckhouse/pull/6327)
 - **[ingress-nginx]** fix cve issues in nginx-exporter image [#6325](https://github.com/deckhouse/deckhouse/pull/6325)
 - **[ingress-nginx]** fix cve issues in kruise image [#6320](https://github.com/deckhouse/deckhouse/pull/6320)
 - **[multitenancy-manager]** Applying a non-valid project or projectType should not stop the main task flow. [#6049](https://github.com/deckhouse/deckhouse/pull/6049)
 - **[pod-reloader]** added the forgotten node selector [#6338](https://github.com/deckhouse/deckhouse/pull/6338)
 - **[prometheus]** Fixed HIGH CVE issues in alertmanager image [#6294](https://github.com/deckhouse/deckhouse/pull/6294)
    Check that the alerts come after the update
 - **[prometheus]** Fixed HIGH CVE issues in trickster image [#6281](https://github.com/deckhouse/deckhouse/pull/6281)
    Check that the prometheus metrics come after the update

