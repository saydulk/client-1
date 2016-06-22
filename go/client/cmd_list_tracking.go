// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol"
)

type CmdListTracking struct {
	libkb.Contextified
	uid      keybase1.UID
	username string
	filter   string
	json     bool
	verbose  bool
	headers  bool
}

func (s *CmdListTracking) ParseArgv(ctx *cli.Context) error {
	byUID := ctx.Bool("uid")
	if len(ctx.Args()) == 1 {
		if byUID {
			var err error
			s.uid, err = libkb.UIDFromHex(ctx.Args()[0])
			if err != nil {
				return fmt.Errorf("UID must be a 32-character hex string")
			}
		} else {
			s.username = ctx.Args()[0]
		}
	} else if len(ctx.Args()) > 1 {
		return fmt.Errorf("list-tracking takes at most one argument")
	}

	s.json = ctx.Bool("json")
	s.verbose = ctx.Bool("verbose")
	s.headers = ctx.Bool("headers")
	s.filter = ctx.String("filter")

	return nil
}

func displayTable(entries []keybase1.UserSummary, verbose bool, headers bool) (err error) {
	if verbose {
		noun := "users"
		if len(entries) == 1 {
			noun = "user"
		}
		GlobUI.Printf("Tracking %d %s:\n\n", len(entries), noun)
	}

	var cols []string

	if headers {
		if verbose {
			cols = []string{
				"Username",
				"Sig ID",
				"PGP fingerprints",
				"When Tracked",
				"Proofs",
			}
		} else {
			cols = []string{"Username"}
		}
	}

	i := 0
	rowfunc := func() []string {
		if i >= len(entries) {
			return nil
		}
		entry := entries[i]
		i++

		if !verbose {
			return []string{entry.Username}
		}

		fps := make([]string, len(entry.Proofs.PublicKeys))
		for i, k := range entry.Proofs.PublicKeys {
			if k.PGPFingerprint != "" {
				fps[i] = k.PGPFingerprint
			}
		}

		row := []string{
			entry.Username,
			entry.SigIDDisplay,
			strings.Join(fps, ", "),
			keybase1.FormatTime(entry.TrackTime),
		}
		for _, proof := range entry.Proofs.Social {
			row = append(row, proof.IdString)
		}
		return row
	}

	GlobUI.Tablify(cols, rowfunc)
	return
}

func DisplayJSON(jsonStr string) error {
	_, err := GlobUI.Println(jsonStr)
	return err
}

func (s *CmdListTracking) Run() error {
	cli, err := GetUserClient(s.G())
	if err != nil {
		return err
	}

	if s.json {
		var jsonStr string
		if s.uid.Exists() {
			jsonStr, err = cli.ListTrackingForUIDJSON(context.TODO(), keybase1.ListTrackingForUIDJSONArg{
				Uid:     s.uid,
				Filter:  s.filter,
				Verbose: s.verbose,
			})
		} else if len(s.username) > 0 {
			jsonStr, err = cli.ListTrackingForUsernameJSON(context.TODO(), keybase1.ListTrackingForUsernameJSONArg{
				Username: s.username,
				Filter:   s.filter,
				Verbose:  s.verbose,
			})
		} else {
			jsonStr, err = cli.ListTrackingJSON(context.TODO(), keybase1.ListTrackingJSONArg{
				Filter:  s.filter,
				Verbose: s.verbose,
			})
		}
		if err != nil {
			return err
		}
		return DisplayJSON(jsonStr)
	}

	var table []keybase1.UserSummary
	if s.uid.Exists() {
		table, err = cli.ListTrackingForUID(context.TODO(), keybase1.ListTrackingForUIDArg{Filter: s.filter, Uid: s.uid})
	} else if len(s.username) > 0 {
		table, err = cli.ListTrackingForUsername(context.TODO(), keybase1.ListTrackingForUsernameArg{Filter: s.filter, Username: s.username})
	} else {
		table, err = cli.ListTracking(context.TODO(), keybase1.ListTrackingArg{Filter: s.filter})
	}
	if err != nil {
		return err
	}
	return displayTable(table, s.verbose, s.headers)
}

func NewCmdListTracking(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name:         "list-tracking",
		ArgumentHelp: "<username>",
		Usage:        "List who username is tracking",
		Action: func(c *cli.Context) {
			cl.ChooseCommand(&CmdListTracking{Contextified: libkb.NewContextified(g)}, "tracking", c)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "f, filter",
				Usage: "Provide a regex filter.",
			},
			cli.BoolFlag{
				Name:  "H, headers",
				Usage: "Show column headers.",
			},
			cli.BoolFlag{
				Name:  "i, uid",
				Usage: "Load user by UID.",
			},
			cli.BoolFlag{
				Name:  "j, json",
				Usage: "Output as JSON (default is text).",
			},
			cli.BoolFlag{
				Name:  "v, verbose",
				Usage: "A full dump, with more gory details.",
			},
		},
	}
}

func (s *CmdListTracking) GetUsage() libkb.Usage {
	return libkb.Usage{
		Config: true,
		API:    true,
	}
}
