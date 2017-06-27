package conf2toml_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cage1016/conf2toml"

	"github.com/stretchr/testify/assert"
)

func TestNormalization(t *testing.T) {
	b, err := conf2toml.Normalization("example/uLunix.conf")

	assert.Nil(t, err)
	assert.Contains(t, string(b), "[System]")
	assert.Contains(t, string(b), "LoginTheme410=10")
	assert.Contains(t, string(b), "ServerName=\"\"")
	assert.Contains(t, string(b), "UPNP_UUID=\"1e26ba4f-8a21-48e0-84b6-0d0f90c41a17\"")
	assert.Contains(t, string(b), "LatestCheckLiveUpdate=\"2017/06/20 14:14:14\"")
	assert.Contains(t, string(b), "Qphoto_LOGO=false")
	assert.Contains(t, string(b), "[NetworkGroup]")
	assert.Contains(t, string(b), "Qdownlod_LOGO=false")
	assert.Contains(t, string(b), "QPHOTOSTATION_LOGO=true")
	assert.Contains(t, string(b), "[END_FLAG]")
	assert.Contains(t, string(b), "[containerstation]")
	assert.Contains(t, string(b), "Name=\"container-station\"")
	assert.Contains(t, string(b), "Class=\"null\"")
	assert.Contains(t, string(b), "Status=\"\"")
}

func TestNormalization2(t *testing.T) {
	b, err := conf2toml.Normalization("notexistpath.conf")
	assert.NotNil(t, err)
	assert.Empty(t, b)
}

func TestNormalizationStdin(t *testing.T) {
	f, err := os.Open("example/uLunix.conf")
	assert.Nil(t, err)

	f2 := conf2toml.NormalizationReader(f)
	assert.NotEmpty(t, f2)

	b, err := ioutil.ReadAll(f2)
	assert.Nil(t, err)
	assert.Contains(t, string(b), "[System]")
	assert.Contains(t, string(b), "LoginTheme410=10")
	assert.Contains(t, string(b), "ServerName=\"\"")
	assert.Contains(t, string(b), "UPNP_UUID=\"1e26ba4f-8a21-48e0-84b6-0d0f90c41a17\"")
	assert.Contains(t, string(b), "LatestCheckLiveUpdate=\"2017/06/20 14:14:14\"")
	assert.Contains(t, string(b), "Qphoto_LOGO=false")
	assert.Contains(t, string(b), "[NetworkGroup]")
	assert.Contains(t, string(b), "Qdownlod_LOGO=false")
	assert.Contains(t, string(b), "QPHOTOSTATION_LOGO=true")
	assert.Contains(t, string(b), "[END_FLAG]")
	assert.Contains(t, string(b), "[containerstation]")
	assert.Contains(t, string(b), "Name=\"container-station\"")
	assert.Contains(t, string(b), "Class=\"null\"")
	assert.Contains(t, string(b), "Status=\"\"")
}

func TestNormalizationStdin2(t *testing.T) {
	f := conf2toml.NormalizationReader(bytes.NewReader(nil))
	assert.NotNil(t, f)
}
