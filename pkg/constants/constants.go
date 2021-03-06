package constants

const (
	ProxyName                  = "cluster"
	TrustedCABundleKey         = "ca-bundle.crt"
	TrustedCABundleMountDir    = "/etc/pki/ca-trust/extracted/pem/"
	TrustedCABundleMountFile   = "tls-ca-bundle.pem"
	InjectTrustedCABundleLabel = "config.openshift.io/inject-trusted-cabundle"
	TrustedCABundleHashName    = "logging.openshift.io/hash"
	KibanaTrustedCAName        = "kibana-trusted-ca-bundle"
	SecretHashPrefix           = "logging.openshift.io/"
	KibanaInstanceName         = "instance"
)

var ReconcileForGlobalProxyList = []string{KibanaTrustedCAName}
