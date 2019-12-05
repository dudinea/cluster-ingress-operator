package aws

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/route53"

	"k8s.io/apimachinery/pkg/util/wait"
)

// StartWatcher starts a file watcher and periodically ensures AWS sessions
// are using the current ca bundle until a message is received on the stop
// or error channels.
func (m *Provider) StartWatcher(operatorReleaseVersion string, stop <-chan struct{}) error {
	errChan := make(chan error)
	go wait.Until(func() {
		err := m.ensureSessionTransport(operatorReleaseVersion)
		if err != nil {
			log.Error(err, "failed to ensure dns client transport")
		}
	}, 1*time.Minute, stop)

	go func() {
		errChan <- m.fileWatcher.Start(stop)
	}()

	// Wait for the watcher to exit or an explicit stop.
	select {
	case <-stop:
		return nil
	case err := <-errChan:
		return err
	}
}

// ensureSessionTransport ensures AWS sessions use the current certificates
// from the file watcher.
func (m *Provider) ensureSessionTransport(operatorReleaseVersion string) error {
	if m.fileWatcher.FileChanged() {
		sess, err := NewProviderSession(m.config, operatorReleaseVersion, m.fileWatcher.GetFileData())
		if err != nil {
			return fmt.Errorf("failed to create dns provider session: %v", err)
		}
		m.lock.Lock()
		m.elb = elb.New(sess, aws.NewConfig().WithRegion(m.config.Region))
		m.route53 = route53.New(sess)
		m.tags = resourcegroupstaggingapi.New(sess, aws.NewConfig().WithRegion("us-east-1"))
		m.lock.Unlock()
		m.fileWatcher.SetFileChanged(false)
		log.Info("updated dns provider session")
	}
	return nil
}
