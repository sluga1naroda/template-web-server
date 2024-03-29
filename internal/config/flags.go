package config

import (
	cli "github.com/urfave/cli/v2"
)

// GlobalFlags CLI flags.
var GlobalFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "debug",
		Usage:       "run in debug mode",
		EnvVars:     []string{"DEBUG"},
		Value:       false,
		Destination: &Get().Application.Debug,
	},
	&cli.StringFlag{
		Name:        "name",
		Usage:       "application name",
		EnvVars:     []string{"NAME"},
		Value:       "some_name",
		Destination: &Get().Application.Name,
	},
	&cli.StringFlag{
		Name:        "log-level, l",
		Usage:       "trace, debug, info, warning, error, fatal or panic",
		Value:       "info",
		EnvVars:     []string{"LOG_LEVEL"},
		Destination: &Get().Application.LogLevel,
	},
	&cli.IntFlag{
		Name:        "http-port",
		Usage:       "HTTP server port",
		Value:       8080,
		EnvVars:     []string{"HTTP_PORT"},
		Destination: &Get().HTTPServer.Port,
	},
	&cli.StringFlag{
		Name:        "http-host",
		Usage:       "HTTP server host",
		Value:       "0.0.0.0",
		EnvVars:     []string{"HTTP_HOST"},
		Destination: &Get().HTTPServer.Host,
	},
	&cli.StringFlag{
		Name:        "User-name",
		Usage:       "User name info",
		Value:       "Admin",
		EnvVars:     []string{"USER_NAME"},
		Destination: &Get().UserInfo.Name,
	},
	&cli.StringFlag{
		Name:        "User-email",
		Usage:       "User email",
		EnvVars:     []string{"USER_EMAIL"},
		Destination: &Get().UserInfo.Email,
	},
	&cli.StringFlag{
		Name:        "User-education",
		Usage:       "User education",
		EnvVars:     []string{"USER_INFO_EDUCATION"},
		Destination: &Get().UserInfo.Education,
	},
	&cli.StringFlag{
		Name:        "User-education-additional",
		Usage:       "User education additional",
		EnvVars:     []string{"USER_INFO_EDUCATION_ADD"},
		Destination: &Get().UserInfo.EducationAdd,
	},
	&cli.StringFlag{
		Name:        "User-education-director-name",
		Usage:       "User education director name",
		EnvVars:     []string{"USER_EDUCATION_DIRECTOR"},
		Destination: &Get().UserInfo.EducationDirector,
	},
	&cli.StringFlag{
		Name:        "User-education-director-phone",
		Usage:       "User education director phone",
		EnvVars:     []string{"USER_EDUCATION_DIRECTOR_PHONE"},
		Destination: &Get().UserInfo.EducationDirectorPhone,
	},
	&cli.StringFlag{
		Name:        "User-friend-1",
		Usage:       "User friend 1",
		EnvVars:     []string{"USER_FRIEND1_NAME"},
		Destination: &Get().UserInfo.Friend1Name,
	},
	&cli.StringFlag{
		Name:        "User-friend-1-email",
		Usage:       "User friend 1 email",
		EnvVars:     []string{"USER_FRIEND1_EMAIL"},
		Destination: &Get().UserInfo.Friend1Email,
	},
	&cli.StringFlag{
		Name:        "User-friend-1-image",
		Usage:       "User friend 1 image",
		EnvVars:     []string{"USER_FRIEND1_IMAGE_URL"},
		Destination: &Get().UserInfo.Friend1ImageUrl,
	},
	&cli.StringFlag{
		Name:        "User-friend-2",
		Usage:       "User friend 2",
		EnvVars:     []string{"USER_FRIEND2_NAME"},
		Destination: &Get().UserInfo.Friend2Name,
	},
	&cli.StringFlag{
		Name:        "User-friend-2-email",
		Usage:       "User friend 2 email",
		EnvVars:     []string{"USER_FRIEND2_EMAIL"},
		Destination: &Get().UserInfo.Friend2Email,
	},
	&cli.StringFlag{
		Name:        "User-friend-2-image",
		Usage:       "User friend 2 image",
		EnvVars:     []string{"USER_FRIEND2_IMAGE_URL"},
		Destination: &Get().UserInfo.Friend2ImageUrl,
	},
	&cli.StringFlag{
		Name:        "User-status",
		Usage:       "User name status",
		EnvVars:     []string{"USER_STATUS"},
		Destination: &Get().UserInfo.Status,
	},
}
