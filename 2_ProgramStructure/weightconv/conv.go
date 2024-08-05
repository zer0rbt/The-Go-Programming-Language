package weightconv

// KgToLb converts Kilogram weight to Pound.
func KgToLb(kg Kilogram) Pound { return Pound(kg / 0.45359237) }

// LbToKg converts Pound weight to Kilogram.
func LbToKg(lb Pound) Kilogram { return Kilogram(lb * 0.45359237) }
