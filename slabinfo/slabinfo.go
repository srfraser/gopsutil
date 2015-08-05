package slabinfo

// Wouldn't it make more sense if this could be accessed as data["name"].ActiveObjects ?
type SlabinfoStat struct {
	Name                string `json:"name"`
	NumberActiveObjects uint64 `json:"activeobjects"`
	NumberObjects       uint64 `json:"numberobjects"`
	ObjectSize          uint64 `json:"objectsize"`
	ObjectsPerSlab      uint64 `json:"objectsperslab"`
	PagesPerSlab        uint64 `json:"pagesperslab"`
	NumberActiveSlabs   uint64 `json:"numberactiveslabs"`
	NumberSlabs         uint64 `json:"numberslabs"`
}
