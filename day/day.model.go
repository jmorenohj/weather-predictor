package day

// Modelo que se usará para la base de datos.
type Day struct {
	Year           int     `bson:"year,omitempty"`
	Day            int     `bson:"day,omitempty"`
	Status         string  `bson:"status,omitempty"`
	RainAmount     float64 `bson:rain_amount`
	FerengiAngle   int     `bson:"ferengi_angle,omitempty"`
	VulcanoAngle   int     `bson:"vulcano_angle,omitempty"`
	BetazoideAngle int     `bson:"betazoide_angle,omitempty"`
}
