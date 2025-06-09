package main

func main() {
	configs := []Config{
		{
			NameDir:       "elk",
			LogDir:        "/var/log/elk",
			RetentionDays: 10,
		},
		{
			NameDir:       "elk2",
			LogDir:        "/var/log/elk2",
			RetentionDays: 20,
		},
		{
			NameDir:       "nginx",
			LogDir:        "/var/log/nginx",
			RetentionDays: 5,
		},
	}

	FilesToDelete(configs)
}
