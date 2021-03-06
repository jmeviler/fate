package main

import (
	"context"
	"fmt"
	"github.com/godcong/chronos"
	"github.com/godcong/fate"
	"github.com/godcong/fate/config"
	"github.com/goextension/log"
	"github.com/spf13/cobra"
	"time"
)

func cmdName() *cobra.Command {
	var path string
	var born string
	var last string
	cmd := &cobra.Command{
		Use:   "name",
		Short: "output the name",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("config will output to %s/config.json\n", path)
			config.DefaultJSONPath = path
			cfg := config.LoadConfig()
			fmt.Printf("config loaded: %+v", cfg)
			bornTime, e := time.Parse(chronos.DateFormat, born)
			if e != nil {
				return
			}
			var ops []fate.Options
			db := fate.InitDatabaseFromConfig(*cfg)

			ops = append(ops, fate.DBOption(db))
			if cfg.SupplyFilter {
				ops = append(ops, fate.SupplyFilter())
			}
			if cfg.BaguaFilter {
				ops = append(ops, fate.BaGuaFilter())
			}
			if cfg.ZodiacFilter {
				ops = append(ops, fate.ZodiacFilter())
			}

			f := fate.NewFate(last, bornTime, ops...)

			e = f.MakeName(context.Background())
			if e != nil {
				log.Errorw("makename", "error", e)
				return
			}
		},
	}
	cmd.Flags().StringVarP(&last, "last", "l", "", "set lastname")
	cmd.Flags().StringVarP(&born, "born", "b", time.Now().Format(chronos.DateFormat), "set birth as 2016/01/02 15:04")
	cmd.Flags().StringVarP(&path, "path", "p", ".", "set the input path")
	return cmd
}
