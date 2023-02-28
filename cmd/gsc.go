package cmd

import "github.com/server-gin/utils"

func generateSsl(path string) error {
	if len(path) == 0 {
		path = "ssl"
	}
	// Change this parameter to issue certificates to different addresses
	c := utils.NewCertConifg("localhost")
	err := c.Generate(path, "cert", "key")
	if err != nil {
		return err
	}
	return nil
}
