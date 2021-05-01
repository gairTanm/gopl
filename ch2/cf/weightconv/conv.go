package weightconv

func PtoK(p Pound) Kilogram { return Kilogram(p / 2.205) }

func KtoP(k Kilogram) Pound { return Pound(k * 2.205) }
