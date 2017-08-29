package conf2toml_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cage1016/conf2toml"
)

func TestNormalization(t *testing.T) {

	tt := []struct {
		name      string
		givenConf string

		wantErr     error
		wantStrings []string
	}{
		{
			name:      "success match",
			givenConf: "example/uLunix.conf",
			wantErr:   nil,
			wantStrings: []string{
				"[System]",
				"Workgroup=\"NAS\"",
				"LoginTheme410=10",
				"Server_Name=\"\"",
				"UPNP_UUID=\"1e26ba4f-8a21-48e0-84b6-0d0f90c41a17\"",
				"Latest_Check_Live_Update=\"2017/06/20 14:14:14\"",
				"Qphoto_LOGO=false",
				"[Network_Group]",
				"Qdownlod_LOGO=false",
				"QPHOTOSTATION_LOGO=true",
				"[END_FLAG]",
				"[container_station]",
				"Name=\"container-station\"",
				"Class=\"null\"",
				"Status=\"\"",
				"cfg__etc_config_qdk_conf=0",
				"Version=\"2.2.13\"",
				"ePassword=\"V2@W5Q9N91N4fXGEEyL+yXOlw==\"",
			},
		},
		{
			name:      "fail cause does not exist file",
			givenConf: "notexistpath.conf",
			wantErr:   errors.New("open notexistpath.conf: no such file or directory"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			b, err := conf2toml.Normalization(tc.givenConf)
			if err != nil {
				assert.EqualError(t, err, tc.wantErr.Error(), "An error was expected")
			} else {
				for _, s := range tc.wantStrings {
					assert.Contains(t, string(b), s)
				}
			}
		})
	}
}

func TestNormalizationStdin(t *testing.T) {

	tt := []struct {
		name      string
		givenConf io.Reader

		wantErr     error
		wantStrings []string
	}{
		{
			name: "success match",
			givenConf: func() io.Reader {
				f, _ := os.Open("example/uLunix.conf")
				return f
			}(),
			wantErr: nil,
			wantStrings: []string{
				"[System]",
				"LoginTheme410=10",
				"Server_Name=\"\"",
				"UPNP_UUID=\"1e26ba4f-8a21-48e0-84b6-0d0f90c41a17\"",
				"Latest_Check_Live_Update=\"2017/06/20 14:14:14\"",
				"Qphoto_LOGO=false",
				"[Network_Group]",
				"Qdownlod_LOGO=false",
				"QPHOTOSTATION_LOGO=true",
				"[END_FLAG]",
				"[container_station]",
				"Name=\"container-station\"",
				"Class=\"null\"",
				"Status=\"\"",
				"cfg__etc_config_qdk_conf=0",
				"Version=\"2.2.13\"",
				"ePassword=\"V2@W5Q9N91N4fXGEEyL+yXOlw==\"",
			},
		},
		{
			name: "",
			givenConf: func() io.Reader {
				return bytes.NewReader(nil)
			}(),
			wantErr: nil,
			wantStrings: []string{
				"",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f2 := conf2toml.NormalizationReader(tc.givenConf)
			b, err := ioutil.ReadAll(f2)

			if err != nil {
				assert.EqualError(t, err, tc.wantErr.Error(), "An error was expected")
			} else {
				for _, s := range tc.wantStrings {
					assert.Contains(t, string(b), s)
				}
			}
		})
	}
}
