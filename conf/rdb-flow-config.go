package conf

import "github.com/tacogips/rdb-flow/rdb"

type RDBFlowConfig struct {
	LatestDDLPath   DirAndFile           `toml:"latest_ddl"`
	GeneratedEntity GeneratedEntity      `toml:"generated_entity"`
	SchemaMetaDir   string               `toml:"schema_meta_dir"`
	InitialDataDir  string               `toml:"initial_data_dir"`
	RDBConfig       rdb.ConnectionConfig `toml:"rdb"`
}

type DirAndFile struct {
	Dir  string `toml:"dir"`
	File string `toml:"file"`
}

type GeneratedEntity struct {
	GeneratedEntityExportPath []GeneratedEntityExportPath `toml:"export_path"`
}

type GeneratedEntityExportPath struct {
	PackageName string `toml:"package_name"`
	Dir         string `toml:"dir"`
	File        string `toml:"file"`
}
