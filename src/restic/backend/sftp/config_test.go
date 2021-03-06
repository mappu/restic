package sftp

import "testing"

var configTests = []struct {
	in  string
	cfg Config
}{
	// first form, user specified sftp://user@host/dir
	{
		"sftp://user@host/dir/subdir",
		Config{User: "user", Host: "host", Dir: "dir/subdir"},
	},
	{
		"sftp://host/dir/subdir",
		Config{Host: "host", Dir: "dir/subdir"},
	},
	{
		"sftp://host//dir/subdir",
		Config{Host: "host", Dir: "/dir/subdir"},
	},
	{
		"sftp://host:10022//dir/subdir",
		Config{Host: "host:10022", Dir: "/dir/subdir"},
	},
	{
		"sftp://user@host:10022//dir/subdir",
		Config{User: "user", Host: "host:10022", Dir: "/dir/subdir"},
	},
	{
		"sftp://user@host/dir/subdir/../other",
		Config{User: "user", Host: "host", Dir: "dir/other"},
	},
	{
		"sftp://user@host/dir///subdir",
		Config{User: "user", Host: "host", Dir: "dir/subdir"},
	},

	// second form, user specified sftp:user@host:/dir
	{
		"sftp:user@host:/dir/subdir",
		Config{User: "user", Host: "host", Dir: "/dir/subdir"},
	},
	{
		"sftp:host:../dir/subdir",
		Config{Host: "host", Dir: "../dir/subdir"},
	},
	{
		"sftp:user@host:dir/subdir:suffix",
		Config{User: "user", Host: "host", Dir: "dir/subdir:suffix"},
	},
	{
		"sftp:user@host:dir/subdir/../other",
		Config{User: "user", Host: "host", Dir: "dir/other"},
	},
	{
		"sftp:user@host:dir///subdir",
		Config{User: "user", Host: "host", Dir: "dir/subdir"},
	},
}

func TestParseConfig(t *testing.T) {
	for i, test := range configTests {
		cfg, err := ParseConfig(test.in)
		if err != nil {
			t.Errorf("test %d:%s failed: %v", i, test.in, err)
			continue
		}

		if cfg != test.cfg {
			t.Errorf("test %d:\ninput:\n  %s\n wrong config, want:\n  %v\ngot:\n  %v",
				i, test.in, test.cfg, cfg)
			continue
		}
	}
}
