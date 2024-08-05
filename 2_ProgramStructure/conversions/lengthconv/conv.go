package lengthconv

// MToFt converts Metric length to Feet.
func MToFt(m Meter) Foot { return Foot(m / 0.3048) }

// FtToM converts Feet length to Metric.
func FtToM(ft Foot) Meter { return Meter(ft * 0.3048) }
