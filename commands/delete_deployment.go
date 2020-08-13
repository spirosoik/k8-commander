package commands

import (
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // // Import solely to initialize client auth plugins.
)

// DeleteDeployment deletes the deployment
type DeleteDeployment struct {
	Namespace string
	Clientset *kubernetes.Clientset
	Logger    logrus.FieldLogger
}

// Execute run the command
func (c *DeleteDeployment) Execute() error {
	deletePolicy := metav1.DeletePropagationForeground
	if err := c.Clientset.CoreV1().Namespaces().Delete(c.Namespace, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		c.Logger.Error(err)
		return err
	}
	return nil
}
