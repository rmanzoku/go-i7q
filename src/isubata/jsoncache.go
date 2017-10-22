package main

var jsonifyCache map[int64]User

func initJsonifyCache() error {
	res := []User{}
	err := db.Get(&res, "SELECT id,name,display_name,avatar_icon FROM user")
	if err != nil {
		return err
	}

	for _, value := range res {
		jsonifyCache[value.ID] = value
	}

	return nil
}

func addJsonifyCache(name, display_name, avatar_icon string) {
	u := User{Name: name, DisplayName: display_name, AvatarIcon: avatar_icon}
	jsonifyCache[u.ID] = u
}
