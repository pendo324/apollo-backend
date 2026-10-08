package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceURLScheme(t *testing.T) {
	t.Parallel()

	token := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	dev := Device{APNSToken: token, Transport: DeviceTransportAPNS}
	assert.NoError(t, dev.Validate())
	assert.Equal(t, "apollo", dev.DeepLinkScheme())

	dev.URLScheme = "phoebus"
	assert.NoError(t, dev.Validate())
	assert.Equal(t, "phoebus", dev.DeepLinkScheme())

	dev.URLScheme = "javascript:alert(1)//"
	assert.Error(t, dev.Validate())
}
