package data

// GetCertificates returns certifications.
func GetCertificates() []Certificate {
	return []Certificate{
		{
			Name:   "Laravel Certified Developer",
			Issuer: "Laravel",
			Date:   "2024",
			URL:    "",
		},
	}
}
