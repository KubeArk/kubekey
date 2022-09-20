package kubeark

import (
	"path/filepath"

	"github.com/kubesphere/kubekey/pkg/common"
	"github.com/kubesphere/kubekey/pkg/core/action"
	"github.com/kubesphere/kubekey/pkg/core/connector"
	"github.com/kubesphere/kubekey/pkg/core/util"
	"github.com/kubesphere/kubekey/pkg/kubeark/templates"
	"github.com/pkg/errors"
)

type GenerateKubearkConfigManifest struct {
	common.KubeAction
}

func (g *GenerateKubearkConfigManifest) Execute(runtime connector.Runtime) error {
	templateAction := action.Template{
		Template: templates.KubearkConfigs,
		Dst:      filepath.Join(common.KubeManifestDir, templates.KubearkConfigs.Name()),
		Data: util.Data{
			"AcmeEmail": g.KubeConf.Cluster.Kubeark.AcmeEmail,
			"Storage":   g.KubeConf.Cluster.Kubeark.Storage,
		},
	}

	templateAction.Init(nil, nil)
	if err := templateAction.Execute(runtime); err != nil {
		return err
	}
	return nil
}

type DeployKubearkConfigs struct {
	common.KubeAction
}

func (d *DeployKubearkConfigs) Execute(runtime connector.Runtime) error {
	if _, err := runtime.GetRunner().SudoCmd("/usr/local/bin/kubectl apply -f /etc/kubernetes/manifests/kubeark-configs.yaml", true); err != nil {
		return errors.Wrap(errors.WithStack(err), "deploy Kubeark configs failed")
	}
	return nil
}

type GenerateKubearkAppManifest struct {
	common.KubeAction
}

func (g *GenerateKubearkAppManifest) Execute(runtime connector.Runtime) error {
	templateAction := action.Template{
		Template: templates.KubearkAppManifest,
		Dst:      filepath.Join(common.KubeManifestDir, templates.KubearkAppManifest.Name()),
		Data: util.Data{
			"IngressHost": g.KubeConf.Cluster.Kubeark.IngressHost,
		},
	}

	templateAction.Init(nil, nil)
	if err := templateAction.Execute(runtime); err != nil {
		return err
	}
	return nil
}

type DeployKubearkApp struct {
	common.KubeAction
}

func (d *DeployKubearkApp) Execute(runtime connector.Runtime) error {
	if _, err := runtime.GetRunner().SudoCmd("/usr/local/bin/kubectl apply -f /etc/kubernetes/manifests/kubeark.yaml", true); err != nil {
		return errors.Wrap(errors.WithStack(err), "deploy Kubeark App failed")
	}
	return nil
}
