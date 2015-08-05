// +build linux

package slabinfo

import (
	"strconv"
	"strings"

	common "github.com/srfraser/gopsutil/common"
)

func parseSlabinfo20(filename string) ([]SlabinfoStat, error) {
	lines, _ := common.ReadLinesOffsetN(filename, 2, -1)

	ret := make([]SlabinfoStat, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)

		// Could also calculate:
		// CacheSize = NumberSlabs * PagesPerSlab * PageSize (getpagesize() or sysconf(_SC_PAGESIZE))
		// UsedPercentage = 100 * NumberActiveObjects / NumberObjects
		// TotalSize = NumberObjects * ObjectSize
		// ActiveSize = NumberActiveObjects * ObjectSize
		d := SlabinfoStat{
			Name:                fields[0],
			NumberActiveObjects: common.MustParseUint64(fields[1]),
			NumberObjects:       common.MustParseUint64(fields[2]),
			ObjectSize:          common.MustParseUint64(fields[3]),
			ObjectsPerSlab:      common.MustParseUint64(fields[4]),
			PagesPerSlab:        common.MustParseUint64(fields[5]),
			NumberActiveSlabs:   common.MustParseUint64(fields[13]),
			NumberSlabs:         common.MustParseUint64(fields[14]),
		}
		ret = append(ret, d)
	}
	return ret, nil
}

func parseSlabinfo11(filename string) ([]SlabinfoStat, error) {
	lines, _ := common.ReadLinesOffsetN(filename, 2, -1)

	ret := make([]SlabinfoStat, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)

		// Could also calculate:
		// CacheSize = NumberSlabs * PagesPerSlab * PageSize (getpagesize() or sysconf(_SC_PAGESIZE))
		// UsedPercentage = 100 * NumberActiveObjects / NumberObjects
		// TotalSize = NumberObjects * ObjectSize
		// ActiveSize = NumberActiveObjects * ObjectSize
		d := SlabinfoStat{
			Name:                fields[0],
			NumberActiveObjects: common.MustParseUint64(fields[1]),
			NumberObjects:       common.MustParseUint64(fields[2]),
			ObjectSize:          common.MustParseUint64(fields[3]),
			NumberActiveSlabs:   common.MustParseUint64(fields[4]),
			NumberSlabs:         common.MustParseUint64(fields[5]),
			PagesPerSlab:        common.MustParseUint64(fields[6]),
		}
		ret = append(ret, d)
	}
	return ret, nil
}

func Slabinfo() ([]SlabinfoStat, error) {
	filename := "/proc/slabinfo"

	// Check for a more idiomatic way to do this test.
	// /proc/slabinfo is only readable by root
	header, err := common.ReadLinesOffsetN(filename, 0, 1)
	if err != nil {
		return nil, err
	}

	version, err := strconv.ParseFloat(strings.Fields(header[0])[3], 64)
	if err != nil {
		return nil, err
	}

	if version >= 2.0 && version < 3.0 {
		return parseSlabinfo20(filename)
	} else if version == 1.1 {
		return parseSlabinfo11(filename)
	} // Add error if version 1.0

	// Set err properly.
	return nil, nil
}
