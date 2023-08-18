package hwclock_test

import (
	"awesomeProject/hwclock"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestHwclockTimeToTimeEl7(t *testing.T) {
	exampleList := []string{
		"Mon 31 Jul 2023 12:59:53 AM EDT  -0.985555 seconds",
		"Tue 21 Mar 2023 02:40:23 AM CST  -0.650751 seconds",
		"Mon 31 Jul 2023 01:00:54 AM EDT  -0.985621 seconds",
		"Mon 31 Jul 2023 12:59:23 PM EDT  -0.985565 seconds",
		"Mon 31 Jul 2023 01:00:24 PM EDT  -1.001135 seconds",
	}

	for _, item := range exampleList {
		res, err := hwclock.HwclockTimeToTimeEl7(item)
		if err != nil {
			require.NoError(t, err)
		}

		spew.Dump(time.Unix(res.Unix(), 0).Format("2006-01-02 15:04:05"))
	}
}
